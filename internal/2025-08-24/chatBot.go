package main

import (
   "encoding/json"
   "fmt"
   "sort"
   "strings"
   "os"
)

func main() {
   data, err := os.ReadFile("../../ignore/chatBot.json")
   if err != nil {
      panic(err)
   }
   var p struct {
      Data Payload
   }
   if err := json.Unmarshal(data, &p); err != nil {
      panic(err)
   }
   
   
   canon := selectCanonicals(p.Data.Models, "openai")
   for _, c := range canon {
      fmt.Printf("canonical=%-24s  rep.slug=%-28s  name=%s\n",
         c.CanonicalKey, c.Rep.Slug, c.Rep.Name)
   }
   

}

type Payload struct{ Models []Model `json:"models"` }

type Model struct {
   Slug        string  `json:"slug"`
   Permaslug   string  `json:"permaslug"`
   Name        string  `json:"name"`
   ShortName   string  `json:"short_name"`
   Author      string  `json:"author"`
   Variant     string  `json:"variant"`
   IsFree      bool    `json:"is_free"`
   IsHidden    bool    `json:"is_hidden"`
   IsDeranked  bool    `json:"is_deranked"`
   IsDisabled  bool    `json:"is_disabled"`

   ModelVersionGroupID *string   `json:"model_version_group_id"`
   ContextLength       *int      `json:"context_length"`
   MaxCompletionTokens *int      `json:"max_completion_tokens"`
   SupportsReasoning   bool      `json:"supports_reasoning"`
   SupportsToolParams  bool      `json:"supports_tool_parameters"`
   HasChatCompletions  bool      `json:"has_chat_completions"`
   WarningMessage      string    `json:"warning_message"`

   ProviderDisplayName string    `json:"provider_display_name"`
   ProviderSlug        string    `json:"provider_slug"`
   Endpoint            *Endpoint `json:"endpoint"`
}

type Endpoint struct {
   ContextLength    *int   `json:"context_length"`
   ProviderModelID  string `json:"provider_model_id"`
   ProviderDisplay  string `json:"provider_display_name"`
   ProviderSlug     string `json:"provider_slug"`
}

type Canonical struct {
   CanonicalKey string
   Rep          Model
}

// ---------------- Canonical key (no dates, no regex) ----------------

func canonicalKey(m Model) string {
   // Highest priority if present
   if m.ModelVersionGroupID != nil && *m.ModelVersionGroupID != "" {
      return *m.ModelVersionGroupID
   }

   // Prefer a stable-looking identifier, then normalize its stem
   if m.Permaslug != "" {
      return normalizeStem(m.Permaslug)
   }
   if m.Slug != "" {
      return normalizeStem(m.Slug)
   }

   // Fallbacks (rare)
   if pid := providerModelID(m); pid != "" {
      // strip transport suffix like ":free"
      if i := strings.IndexByte(pid, ':'); i >= 0 {
         pid = pid[:i]
      }
      return normalizeStem(pid)
   }
   return ""
}

// normalizeStem keeps the vendor prefix (before '/'),
// then truncates the model part at the first "version-like" token:
//   - purely numeric tokens (e.g., "2024", "07", "18", "2")
//   - tokens equal to "preview", "beta", or "latest" (case-insensitive)
func normalizeStem(id string) string {
   id = strings.TrimSpace(id)
   if id == "" {
      return id
   }
   parts := strings.SplitN(id, "/", 2)
   if len(parts) != 2 {
      // no vendor prefix; just normalize the single part
      return truncateAtVersionish(id)
   }
   vendor := parts[0]
   model := truncateAtVersionish(parts[1])
   if model == "" {
      // if everything after truncation vanished, fall back to the original second part
      model = parts[1]
   }
   return vendor + "/" + model
}

func truncateAtVersionish(s string) string {
   toks := strings.Split(s, "-")
   out := make([]string, 0, len(toks))
   for _, t := range toks {
      lt := strings.ToLower(strings.TrimSpace(t))
      if lt == "" {
         continue
      }
      // stop at the first version-like token
      if isNumericOnly(lt) || lt == "preview" || lt == "beta" || lt == "latest" {
         break
      }
      out = append(out, lt)
   }
   // if we never appended anything (e.g., started numeric), return original
   if len(out) == 0 {
      return s
   }
   return strings.Join(out, "-")
}

func isNumericOnly(s string) bool {
   if s == "" {
      return false
   }
   for i := 0; i < len(s); i++ {
      if s[i] < '0' || s[i] > '9' {
         return false
      }
   }
   return true
}

func providerModelID(m Model) string {
   if m.Endpoint != nil && m.Endpoint.ProviderModelID != "" {
      return m.Endpoint.ProviderModelID
   }
   return ""
}

// ---------------- Selection/scoring (unchanged, no dates) ----------------

type CanonicalScore struct {
   Key   string
   Model Model
   Score int64
}

func selectCanonicals(models []Model, authorFilter string) []Canonical {
   group := map[string][]Model{}
   for _, m := range models {
      if authorFilter != "" && !strings.EqualFold(m.Author, authorFilter) {
         continue
      }
      key := canonicalKey(m) // normalized stem ensures no dated canonical
      if key == "" {
         continue
      }
      group[key] = append(group[key], m)
   }

   out := make([]Canonical, 0, len(group))
   for key, list := range group {
      // pick the best representative for this family
      best := list[0]
      bestScore := scoreModel(key, best)
      for i := 1; i < len(list); i++ {
         if sc := scoreModel(key, list[i]); sc > bestScore {
            best, bestScore = list[i], sc
         }
      }
      out = append(out, Canonical{CanonicalKey: key, Rep: best})
   }

   sort.Slice(out, func(i, j int) bool { return out[i].CanonicalKey < out[j].CanonicalKey })
   return out
}

func scoreModel(key string, m Model) int64 {
   var s int64
   // Prefer entries whose slug/permaslug equals the canonical stem
   if m.Slug == key {
      s += 1_000_000
   }
   if m.Permaslug == key {
      s += 800_000
   }

   // Penalize unusable/less desirable
   if m.IsDisabled {
      s -= 1_000_000
   }
   if m.IsHidden {
      s -= 200_000
   }
   if m.IsDeranked {
      s -= 100_000
   }
   if looksPreview(m) {
      s -= 50_000
   }

   // Prefer production traits
   if strings.EqualFold(m.Variant, "standard") {
      s += 10_000
   }
   if !m.IsFree {
      s += 5_000
   }

   // Capabilities & limits
   if m.SupportsReasoning {
      s += 2_000
   }
   if m.SupportsToolParams {
      s += 1_500
   }
   if m.HasChatCompletions {
      s += 1_000
   }
   s += int64(maxInt(ptr(m.ContextLength), ptr(endpointContext(m))))
   s += int64(ptr(m.MaxCompletionTokens))

   // Gentle tie-breaker for shorter slugs
   s -= int64(len(m.Slug)) / 10
   return s
}

func looksPreview(m Model) bool {
   fields := []string{m.Slug, m.Permaslug, m.Name, m.ShortName, m.WarningMessage}
   if m.Endpoint != nil {
      fields = append(fields, m.Endpoint.ProviderModelID)
   }
   for _, f := range fields {
      f = strings.ToLower(f)
      if strings.Contains(f, "preview") || strings.Contains(f, "beta") {
         return true
      }
   }
   return false
}

func endpointContext(m Model) *int {
   if m.Endpoint != nil && m.Endpoint.ContextLength != nil {
      return m.Endpoint.ContextLength
   }
   return nil
}

func ptr(p *int) int {
   if p == nil {
      return 0
   }
   return *p
}

func maxInt(a, b int) int {
   if a > b {
      return a
   }
   return b
}

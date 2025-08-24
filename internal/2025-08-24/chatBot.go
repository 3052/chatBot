package main

import (
   "encoding/json"
   "fmt"
   "sort"
   "strings"
)

// ---- Minimal JSON shapes. Extend as needed to match your payload. ----

type Payload struct {
   Models []Model `json:"models"`
}

type Model struct {
   Slug        string `json:"slug"`
   Permaslug   string `json:"permaslug"`
   Name        string `json:"name"`
   ShortName   string `json:"short_name"`
   Author      string `json:"author"`

   ModelVersionGroupID *string `json:"model_version_group_id"`

   ContextLength       *int    `json:"context_length"`
   MaxCompletionTokens *int    `json:"max_completion_tokens"`
   Variant             string  `json:"variant"` // e.g., "standard", "free", etc.
   IsFree              bool    `json:"is_free"`
   IsHidden            bool    `json:"is_hidden"`
   IsDeranked          bool    `json:"is_deranked"`
   IsDisabled          bool    `json:"is_disabled"`

   // Capability flags (flattened from your JSON; add more if you want)
   SupportsReasoning       bool `json:"supports_reasoning"`
   SupportsToolParameters  bool `json:"supports_tool_parameters"`
   HasChatCompletions      bool `json:"has_chat_completions"`

   Endpoint *Endpoint `json:"endpoint"`

   // Some deployments also expose provider fields top-level. If yours does, add them.
   ProviderDisplayName string `json:"provider_display_name"`
   ProviderSlug        string `json:"provider_slug"`

   // Optional text to help exclude previews without regex/date checks
   WarningMessage string `json:"warning_message"`
}

type Endpoint struct {
   ContextLength     *int   `json:"context_length"`
   ProviderModelID   string `json:"provider_model_id"`
   ProviderDisplay   string `json:"provider_display_name"`
   ProviderSlug      string `json:"provider_slug"`
}

// ---- Canonical family key (no date, no regex) ----

// Priority: model_version_group_id > permaslug > provider_model_id (sans transport suffix) > slug.
func canonicalKey(m Model) string {
   if m.ModelVersionGroupID != nil && *m.ModelVersionGroupID != "" {
      return *m.ModelVersionGroupID
   }
   if m.Permaslug != "" {
      return m.Permaslug
   }
   if pid := providerModelID(m); pid != "" {
      // Drop transport suffix after ':' (e.g., "openai/gpt-oss-20b:free" -> "openai/gpt-oss-20b")
      if idx := strings.IndexByte(pid, ':'); idx >= 0 {
         return pid[:idx]
      }
      return pid
   }
   // Final fallback
   return m.Slug
}

func providerModelID(m Model) string {
   if m.Endpoint != nil && m.Endpoint.ProviderModelID != "" {
      return m.Endpoint.ProviderModelID
   }
   return ""
}

// ---- Preference scoring (no date logic) ----

func scoreModel(m Model) int64 {
   var s int64 = 0

   // Strong negatives for unusable entries
   if m.IsDisabled {
      s -= 1_000_000
   }
   if m.IsHidden {
      s -= 200_000
   }
   if m.IsDeranked {
      s -= 100_000
   }

   // Prefer production over preview/beta by simple substring checks (no regex)
   if looksPreview(m) {
      s -= 50_000
   }

   // Prefer "standard" variant
   if strings.EqualFold(m.Variant, "standard") {
      s += 10_000
   }

   // Non-free often has fewer limits
   if !m.IsFree {
      s += 5_000
   }

   // Feature richness
   if m.SupportsReasoning {
      s += 2_000
   }
   if m.SupportsToolParameters {
      s += 1_500
   }
   if m.HasChatCompletions {
      s += 1_000
   }

   // Bigger context and completions
   s += int64(maxInt(ptrOrZero(m.ContextLength), ptrOrZero(endpointContext(m))))
   s += int64(ptrOrZero(m.MaxCompletionTokens))

   // Provider preference (when the same family exists via multiple providers)
   provider := strings.ToLower(preferredProviderName(m))
   if provider == "openai" {
      s += 2_500
   }

   return s
}

func looksPreview(m Model) bool {
   // Quick, non-regex, date-free checks
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

func preferredProviderName(m Model) string {
   if m.ProviderDisplayName != "" {
      return m.ProviderDisplayName
   }
   if m.Endpoint != nil && m.Endpoint.ProviderDisplay != "" {
      return m.Endpoint.ProviderDisplay
   }
   // Fallback: derive from provider slug if display name is empty
   if m.ProviderSlug != "" {
      return m.ProviderSlug
   }
   if m.Endpoint != nil && m.Endpoint.ProviderSlug != "" {
      return m.Endpoint.ProviderSlug
   }
   return ""
}

func endpointContext(m Model) *int {
   if m.Endpoint != nil && m.Endpoint.ContextLength != nil {
      return m.Endpoint.ContextLength
   }
   return nil
}

func ptrOrZero(p *int) int {
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

// ---- Select canonical per family ----

func selectCanonical(models []Model, authorFilter string) map[string]Model {
   // Group by canonical key, optionally filter by author
   group := map[string][]Model{}
   for _, m := range models {
      if authorFilter != "" && !strings.EqualFold(m.Author, authorFilter) {
         continue
      }
      key := canonicalKey(m)
      group[key] = append(group[key], m)
   }

   // Pick best per key
   best := make(map[string]Model, len(group))
   for key, list := range group {
      sort.SliceStable(list, func(i, j int) bool {
         // Higher score first; tie-breakers keep stable ordering
         return scoreModel(list[i]) > scoreModel(list[j])
      })
      best[key] = list[0]
   }
   return best
}

// ---- Example usage ----

func main() {
   raw := `{"models":[
     {"slug":"openai/gpt-4o-2024-11-20","permaslug":"openai/gpt-4o","name":"OpenAI: GPT-4o","author":"openai","variant":"standard","supports_reasoning":true,"has_chat_completions":true,"context_length":128000},
     {"slug":"openai/gpt-4o-2024-08-06","permaslug":"openai/gpt-4o","name":"OpenAI: GPT-4o","author":"openai","variant":"free","is_free":true,"has_chat_completions":true,"context_length":128000},
     {"slug":"openai/o1-preview-2024-09-12","permaslug":"openai/o1-preview-2024-09-12","name":"OpenAI: o1-preview","author":"openai","variant":"standard","supports_reasoning":true,"has_chat_completions":true},
     {"slug":"openai/gpt-5","permaslug":"openai/gpt-5","name":"OpenAI: GPT-5","author":"openai","variant":"standard","supports_reasoning":true,"has_chat_completions":true,"context_length":400000}
   ]}`
   var p Payload
   if err := json.Unmarshal([]byte(raw), &p); err != nil {
      panic(err)
   }

   canon := selectCanonical(p.Models, "openai")
   fmt.Println("Canonical OpenAI models (no date logic):")
   // Stable display: sort keys for consistent output
   keys := make([]string, 0, len(canon))
   for k := range canon {
      keys = append(keys, k)
   }
   sort.Strings(keys)
   for _, k := range keys {
      m := canon[k]
      fmt.Printf("key=%-28s  slug=%-28s  variant=%-10s  score=%d\n",
         k, m.Slug, m.Variant, scoreModel(m))
   }
}

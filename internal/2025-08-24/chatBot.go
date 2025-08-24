package main

import (
   "fmt"
   "io"
   "os"
   "strings"
)

// ----- exact input layout: {"data":{"models":[ ... ]}} -----

type Root struct {
   Data struct {
      Models []Model `json:"models"`
   } `json:"data"`
}

type Model struct {
   Author string `json:"author"`

   // NOT used to form canonical; only for guards/display
   Slug      string `json:"slug"`
   Permaslug string `json:"permaslug"`

   // canonical fields
   ModelVersionGroupID *string `json:"model_version_group_id"`
   ShortName           string  `json:"short_name"`
   Name                string  `json:"name"`

   // optional (irrelevant to family key; kept for possible rep scoring)
   Variant             string `json:"variant"`
   IsFree              bool   `json:"is_free"`
   IsHidden            bool   `json:"is_hidden"`
   IsDeranked          bool   `json:"is_deranked"`
   IsDisabled          bool   `json:"is_disabled"`
   ContextLength       *int   `json:"context_length"`
   MaxCompletionTokens *int   `json:"max_completion_tokens"`
   SupportsReasoning   bool   `json:"supports_reasoning"`
   SupportsToolParams  bool   `json:"supports_tool_parameters"`
   HasChatCompletions  bool   `json:"has_chat_completions"`
   ProviderDisplayName string `json:"provider_display_name"`
   ProviderSlug        string `json:"provider_slug"`
   Endpoint            *struct {
      ContextLength   *int   `json:"context_length"`
      ProviderDisplay string `json:"provider_display_name"`
      ProviderSlug    string `json:"provider_slug"`
   } `json:"endpoint"`
}

type Family struct {
   Key    string
   Members int
}


// ---------- canonical families (fields-only; no slug/date heuristics) ----------

func canonicalFamilies(models []Model, authorFilter string) []Family {
   group := map[string]int{}
   for _, m := range models {
      if authorFilter != "" && !strings.EqualFold(m.Author, authorFilter) {
         continue
      }
      key, ok := familyKey(m)
      if !ok {
         continue // skip: cannot form a stable, human canonical
      }
      group[key]++
   }
   out := make([]Family, 0, len(group))
   for k, n := range group {
      out = append(out, Family{Key: k, Members: n})
   }
   return out
}

// familyKey chooses a canonical ONLY from MVGID or a human label.
// It NEVER returns slug/permaslug or any date-coded identifier.
// Additional guard: reject labels that contain parentheses.
func familyKey(m Model) (string, bool) {
   // 1) explicit group id
   if m.ModelVersionGroupID != nil && *m.ModelVersionGroupID != "" {
      return *m.ModelVersionGroupID, true
   }

   author := strings.TrimSpace(m.Author)
   if author == "" {
      return "", false
   }
   al := strings.ToLower(author)

   // 2) short_name (must look human)
   if s := strings.TrimSpace(m.ShortName); isAcceptableHumanLabel(s, m) {
      return al + "/" + collapseSpaces(s), true
   }

   // 3) name (drop "Author: " prefix; must look human)
   if n := strings.TrimSpace(m.Name); n != "" {
      n = dropVendorPrefix(n, m.Author)
      if isAcceptableHumanLabel(n, m) {
         return al + "/" + collapseSpaces(n), true
      }
   }

   // no canonical fields -> skip
   return "", false
}

// isAcceptableHumanLabel rejects anything that is clearly an identifier,
// without doing date parsing or regex:
//  - empty
//  - contains '/'
//  - contains '(' or ')'   (e.g., "GPT-4o-mini (2024-07-18)", "(Preview)", "(extended)")
//  - exactly equals the slug or permaslug tail (case-insensitive)
func isAcceptableHumanLabel(label string, m Model) bool {
   label = strings.TrimSpace(label)
   if label == "" {
      return false
   }
   if strings.Contains(label, "/") {
      return false
   }
   if strings.Contains(label, "(") || strings.Contains(label, ")") {
      return false
   }
   ll := strings.ToLower(label)
   if tailEquals(ll, m.Slug) || tailEquals(ll, m.Permaslug) {
      return false
   }
   return true
}

// tailEquals compares label to the part of id after the first '/' (case-insensitive).
func tailEquals(labelLower, id string) bool {
   if id == "" {
      return false
   }
   parts := strings.SplitN(strings.ToLower(id), "/", 2)
   tail := parts[len(parts)-1]
   return labelLower == tail
}

func dropVendorPrefix(name, author string) string {
   i := strings.Index(name, ":")
   if i < 0 {
      return strings.TrimSpace(name)
   }
   left := strings.TrimSpace(name[:i])
   right := strings.TrimSpace(name[i+1:])
   if left != "" && right != "" && strings.EqualFold(left, author) {
      return right
   }
   return strings.TrimSpace(name)
}

func collapseSpaces(s string) string {
   return strings.Join(strings.Fields(s), " ")
}

// ---------- io / util ----------

func readAll() ([]byte, error) {
   if len(os.Args) > 1 {
      return os.ReadFile(os.Args[1])
   }
   return io.ReadAll(os.Stdin)
}

func must(err error) {
   if err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
   }
}


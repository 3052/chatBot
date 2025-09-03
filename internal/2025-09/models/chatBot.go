package models

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "slices"
   "time"
)

func (a *model) tokens(front *frontend) int64 {
   i := slices.IndexFunc(front.Models, func(b *metadata) bool {
      return b.Slug == a.slug
   })
   return front.Analytics[front.Models[i].Permaslug].TotalPromptTokens
}

type frontend struct {
   Analytics map[string]struct {
      TotalPromptTokens int64 `json:"total_prompt_tokens"`
   }
   Models []*metadata
}

func delete_metadata(m *metadata) bool {
   if m.ContextLength < 128000 {
      return true
   }
   if m.Endpoint == nil {
      return true
   }
   if m.Endpoint.ModelVariantSlug != m.Slug {
      return true
   }
   const day = 24 * time.Hour
   const month = 30 * day
   // 6 month is 150
   const updated_at = 5 * month
   if time.Since(m.UpdatedAt) >= updated_at {
      return true
   }
   return m.WarningMessage != ""
}

func (b *model) contains(a []*metadata) bool {
   for _, a1 := range a {
      if a1.Slug == b.slug {
         return true
      }
   }
   return false
}

func (b *metadata) contains(a []*model) bool {
   for _, a1 := range a {
      if a1.slug == b.Slug {
         return true
      }
   }
   return false
}

type metadata struct {
   Author        string
   ContextLength int       `json:"context_length"`
   Endpoint      *struct { // DELETE
      ModelVariantSlug string `json:"model_variant_slug"`
   }
   Permaslug      string
   ShortName      string `json:"short_name"`
   Slug           string
   UpdatedAt      time.Time `json:"updated_at"`
   WarningMessage string    `json:"warning_message"`
}

func (m *model) String() string {
   b := fmt.Appendln(nil, "slug =", m.slug)
   b = fmt.Append(b, "url = ", m.url)
   if m.info != "" {
      b = fmt.Append(b, "\ninfo = ", m.info)
   }
   if m.ok {
      b = append(b, "\nok = true"...)
   }
   return string(b)
}

type byte_slice[T any] []byte

func get_frontend() (byte_slice[frontend], error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (f *frontend) unmarshal(data byte_slice[frontend]) error {
   var value struct {
      Data frontend
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *f = value.Data
   return nil
}

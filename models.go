package models

import (
   "encoding/json"
   "io"
   "net/http"
   "time"
)

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

func (m *metadata) tokens(front *frontend) int64 {
   return front.Analytics[m.Permaslug].TotalPromptTokens
}

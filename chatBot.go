package chatBot

import (
   "encoding/json"
   "io"
   "net/http"
   "time"
)

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
   if time.Since(m.UpdatedAt) >= 5*month {
      return true
   }
   return m.WarningMessage != ""
}

type metadata struct {
   Author        string
   ContextLength int       `json:"context_length"`
   Endpoint      *struct { // DELETE
      ModelVariantSlug string `json:"model_variant_slug"`
   }
   ShortName string `json:"short_name"`
   Slug      string
   UpdatedAt time.Time `json:"updated_at"`
   WarningMessage string `json:"warning_message"`
}

type byte_slice[T any] []byte

func (m *metadatas) unmarshal(data byte_slice[metadatas]) error {
   var value struct {
      Data struct {
         Models []*metadata
      }
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *m = value.Data.Models
   return nil
}

type metadatas []*metadata

func (a models) contains(b *metadata) bool {
   for _, a1 := range a {
      if a1.slug == b.Slug {
         return true
      }
   }
   return false
}

func (a metadatas) contains(b *model) bool {
   for _, a1 := range a {
      if a1.Slug == b.slug {
         return true
      }
   }
   return false
}

func get_metadatas() (byte_slice[metadatas], error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

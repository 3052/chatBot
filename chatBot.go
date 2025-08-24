package chatBot

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
)

func get_models() (byte_slice[models], error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type model struct {
   Slug string
   Author string
   ShortName string `json:"short_name"`
   ContextLength int `json:"context_length"` // DELETE
   Endpoint *struct { // DELETE
      Pricing struct {
         Completion price `json:",string"`
         Prompt price `json:",string"`
      }
   }
}

func (m *model) String() string {
   var b []byte
   b = fmt.Appendln(b, "slug =", m.Slug)
   b = fmt.Appendln(b, "author =", m.Author)
   b = fmt.Appendln(b, "short name =", m.ShortName)
   b = fmt.Append(b, "context length = ", m.ContextLength)
   if m.Endpoint != nil {
      b = fmt.Append(b, "\ncompletion = ", m.Endpoint.Pricing.Completion)
      b = fmt.Append(b, "\nprompt = ", m.Endpoint.Pricing.Prompt)
   }
   return string(b)
}

type byte_slice[T any] []byte

type models []*model

func (m *models) unmarshal(data byte_slice[models]) error {
   var value struct {
      Data struct {
         Models []*model
      }
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *m = value.Data.Models
   return nil
}

func (p price) String() string {
   return fmt.Sprintf("$%.2f/M", float64(p) * 1_000_000)
}

type price float64

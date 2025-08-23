package chatBot

import (
   "encoding/json"
   "log"
   "net/http"
)

type model struct {
   Author string
   ShortName string `json:"short_name"`
   Slug string
}

type bytes[T any] []byte

func get_models() (bytes[models], error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   req.URL.RawQuery = "context=128000"
   log.Print("BEGIN")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type models struct {
   Data struct {
      Models []model
   }
}

func (m *models) unmarshal(data bytes[models]) error {
   return json.Unmarshal(data, m)
}

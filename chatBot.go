package chatBot

import (
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

type model struct {
   Author string
   ShortName string `json:"short_name"`
   Slug string
}

type bytes[T any] []byte

// 128000
func get_models(content int) (bytes[models], error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   if content >= 1 {
      req.URL.RawQuery = "context=" + strconv.Itoa(content)
   }
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

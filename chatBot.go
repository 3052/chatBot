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

func get_models() ([]model, error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   req.URL.RawQuery = "context=128000"
   log.Print("BEGIN")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         Models []model
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return value.Data.Models, nil
}

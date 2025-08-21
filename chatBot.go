package chatBot

import "net/http"

func find() (*http.Response, error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   req.URL.RawQuery = "context=128000"
   return http.DefaultClient.Do(req)
}

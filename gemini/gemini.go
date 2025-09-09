package gemini

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "os/exec"
)

type completion struct {
   Candidates []struct {
      Content content
   }
}

type content struct {
   Parts []part `json:"parts"`
}

type part struct {
   Text string `json:"text"`
}

func (p prompt) generate(key string) ([]completion, error) {
   data, err := json.Marshal(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://generativelanguage.googleapis.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "key=" + key
   req.URL.Path = "/v1beta/models/gemini-2.5-pro:streamGenerateContent"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var completions []completion
   err = json.NewDecoder(resp.Body).Decode(&completions)
   if err != nil {
      return nil, err
   }
   return completions, nil
}

type prompt struct {
   Contents []content `json:"contents"`
}

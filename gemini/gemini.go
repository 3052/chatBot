package gemini

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strings"
)

func (c Completion) String() string {
   var (
      data strings.Builder
      line bool
   )
   for _, candidate := range c.Candidates {
      for _, part := range candidate.Content.Parts {
         if line {
            data.WriteByte('\n')
         } else {
            line = true
         }
         data.WriteString(part.Text)
      }
   }
   return data.String()
}

type Completion struct {
   Candidates []struct {
      Content Content
   }
}

type Part struct {
   Text string `json:"text"`
}

type Content struct {
   Parts []Part `json:"parts"`
}

func (p Prompt) Generate(key string) ([]Completion, error) {
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
   var completions []Completion
   err = json.NewDecoder(resp.Body).Decode(&completions)
   if err != nil {
      return nil, err
   }
   return completions, nil
}

type Prompt struct {
   Contents []Content `json:"contents"`
}

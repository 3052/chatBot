package gemini

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type content struct {
   Role  string `json:"role"`
   Parts []part `json:"parts"`
}

type part struct {
   Text string `json:"text"`
}

func generate(text string) (*http.Response, error) {
   var value struct {
      Contents []content `json:"contents"`
      GenerationConfig struct {
         ThinkingConfig struct {
            ThinkingBudget int `json:"thinkingBudget"`
         } `json:"thinkingConfig"`
      } `json:"generationConfig"`
      Tools []struct {
         GoogleSearch struct{} `json:"googleSearch"`
      } `json:"tools"`
   }
   value.Contents = []content{
      {
         Role: "user",
         Parts: []part{
            {Text: text},
         },
      },
   } 
   value.GenerationConfig.ThinkingConfig.ThinkingBudget = -1
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://generativelanguage.googleapis.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v1beta/models/gemini-2.5-pro:streamGenerateContent"
   req.URL.RawQuery = "key=${GEMINI_API_KEY}"
   req.Header.Set("content-type", "application/json")
   return http.DefaultClient.Do(req)
}

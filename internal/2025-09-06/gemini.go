package main

import (
   "bytes"
   "encoding/json"
   "net/http"
   "os"
)

type part struct {
   Text string `json:"text"`
}

type content struct {
   Role  string `json:"role"`
   Parts []part `json:"parts"`
}

func main() {
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
            {Text: "2+3"},
         },
      },
   } 
   value.GenerationConfig.ThinkingConfig.ThinkingBudget = -1
   data, err := json.Marshal(value)
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest(
      "POST", "https://generativelanguage.googleapis.com", bytes.NewReader(data),
   )
   req.URL.Path = "/v1beta/models/gemini-2.5-pro:streamGenerateContent"
   req.URL.RawQuery = "key=${GEMINI_API_KEY}"
   req.Header.Set("content-type", "application/json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      panic(err)
   }
   err = resp.Write(os.Stdout)
   if err != nil {
      panic(err)
   }
}

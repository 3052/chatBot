package main

import (
   "fmt"
   "os"
   "sort"
   "encoding/json"
   "errors"
)

func main() {
   data, err := os.ReadFile("../../ignore/chatBot.json")
   if err != nil {
      panic(err)
   }
   
   
   

   var r Root
   must(json.Unmarshal(data, &r))
   if r.Data.Models == nil {
      must(errors.New(`missing "data.models"`))
   }

   // filter to OpenAI; set to "" to include all authors
   const authorFilter = "openai"

   fams := canonicalFamilies(r.Data.Models, authorFilter)

   // stable print
   sort.Slice(fams, func(i, j int) bool { return fams[i].Key < fams[j].Key })
   for _, f := range fams {
      fmt.Printf("canonical=%s  members=%d\n", f.Key, f.Members)
   }
}

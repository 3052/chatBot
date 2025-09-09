package main

import (
   "chatBot/gemini"
   "flag"
   "fmt"
   "os/exec"
)

func completion(text string) error {
   key, err := exec.Command("credential", "-k", "GEMINI_API_KEY").Output()
   if err != nil {
      return err
   }
   prompt := gemini.Prompt{
      Contents: []gemini.Content{
         {
            Parts: []gemini.Part{
               { Text: text },
            },
         },
      },
   }
   completions, err := prompt.Generate(string(key))
   if err != nil {
      return err
   }
   for _, completion := range completions {
      fmt.Println(completion)
   }
   return nil
}

func main() {
   prompt := flag.String("p", "", "prompt")
   flag.Parse()
   if *prompt != "" {
      err := completion(*prompt)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

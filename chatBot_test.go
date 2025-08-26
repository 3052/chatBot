package chatBot

import (
   "bytes"
   "encoding/json"
   "fmt"
   "os"
   "slices"
   "testing"
)

func TestModel(t *testing.T) {
   data, err := os.ReadFile("ignore/chatBot.json")
   if err != nil {
      t.Fatal(err)
   }
   var modelsVar models
   err = modelsVar.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   modelsVar = slices.DeleteFunc(modelsVar, delete_model)
   for _, slug := range good_slugs {
      i := slices.IndexFunc(modelsVar, func(m *model) bool {
         return m.Slug == slug
      })
      if i == -1 {
         t.Fatal(slug)
      }
   }
   for _, modelVar := range modelsVar {
      fmt.Print(modelVar, "\n\n")
   }
   for _, modelVar := range modelsVar {
      fmt.Println(modelVar.Slug)
   }
}

func TestMarshal(t *testing.T) {
   data, err := get_models()
   if err != nil {
      t.Fatal(err)
   }
   var data1 bytes.Buffer
   err = json.Indent(&data1, data, "", " ")
   if err != nil {
      t.Fatal(err)
   }
   err = os.WriteFile("ignore/chatBot.json", data1.Bytes(), os.ModePerm)
   if err != nil {
      t.Fatal(err)
   }
}

var good_slugs = []string{
   "anthropic/claude-opus-4.1",
   "anthropic/claude-sonnet-4",
   "google/gemini-2.5-flash",
   "google/gemini-2.5-pro",
   "openai/gpt-4o",
   "openai/gpt-5",
   "qwen/qwen3-coder",
}

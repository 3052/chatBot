package chatBot

import (
   "bytes"
   "cmp"
   "encoding/json"
   "fmt"
   "os"
   "slices"
   "testing"
)

func TestUnmarshal(t *testing.T) {
   data, err := os.ReadFile("ignore/chatBot.json")
   if err != nil {
      t.Fatal(err)
   }
   var modelsVar models
   err = modelsVar.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   modelsVar = slices.DeleteFunc(modelsVar, func(m *model) bool {
      if m.ContextLength < 128000 {
         return true
      }
      if m.Endpoint == nil {
         return true
      }
      return false
   })
   slices.SortFunc(modelsVar, func(a, b *model) int {
      return cmp.Compare(a.Slug, b.Slug)
   })
   modelsVar = slices.CompactFunc(modelsVar, func(a, b *model) bool {
      return a.Slug == b.Slug
   })
   for _, slug := range good_slugs {
      i := slices.IndexFunc(modelsVar, func(m *model) bool {
         return m.Slug == slug
      })
      if i == -1 {
         t.Fatal(slug)
      }
   }
   file, err := os.Create("chatBot.txt")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   for _, modelVar := range modelsVar {
      _, err = fmt.Fprintf(file, "https://openrouter.ai/%v\n", modelVar.Slug)
      if err != nil {
         t.Fatal(err)
      }
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

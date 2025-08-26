package chatBot

import (
   "bytes"
   "encoding/json"
   "os"
   "slices"
   "testing"
)

func TestOne(t *testing.T) {
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
   // too many?
   for _, modelVar := range modelsVar {
      _, ok := canonical[modelVar.Slug]
      if !ok {
         t.Fatal(modelVar.Slug)
      }
   }
   // too few?
   for key, value := range canonical {
      if value {
         i := slices.IndexFunc(modelsVar, func(m *model) bool {
            return m.Slug == key
         })
         if i == -1 {
            t.Fatal(key)
         }
      }
   }
}

func TestZero(t *testing.T) {
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

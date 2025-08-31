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
   var all_metadatas metadatas
   err = all_metadatas.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   all_metadatas = slices.DeleteFunc(all_metadatas, delete_metadata)
   for _, one_metadata := range all_metadatas {
      if !all_models.contains(one_metadata) {
         t.Fatal(one_metadata, " missing from all_models")
      }
   }
   for _, one_model := range all_models {
      if !all_metadatas.contains(one_model) {
         if one_model.err != nil {
            t.Fatal(one_model.slug, " extra in all_models")
         } else {
            t.Fatal(one_model.slug, " missing from JSON")
         }
      }
   }
}

func TestZero(t *testing.T) {
   data, err := get_metadatas()
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

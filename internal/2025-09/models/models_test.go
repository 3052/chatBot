package models

import (
   "bytes"
   "encoding/json"
   "fmt"
   "log"
   "os"
   "slices"
   "testing"
)

const name = "ignore/chatBot.json"

func TestOne(t *testing.T) {
   log.SetFlags(log.Ltime)
   data, err := os.ReadFile(name)
   if err != nil {
      t.Fatal(err)
   }
   var front frontend
   err = front.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   log.Println("len(front.Models)", len(front.Models))
   front.Models = slices.DeleteFunc(front.Models, delete_metadata)
   log.Println("len(front.Models)", len(front.Models))
   for _, one_metadata := range front.Models {
      // log.Println("tokens", one_metadata.tokens(&front))
      if !one_metadata.contains(all_models) {
         t.Fatal(one_metadata.Slug, " missing from all_models")
      }
   }
   for _, one_model := range all_models {
      if !one_model.contains(front.Models) {
         t.Fatal(one_model.slug, " extra in all_models")
      }
   }
}

func TestTwo(t *testing.T) {
   var count int
   for _, one_model := range all_models {
      if one_model.ok {
         fmt.Print(one_model, "\n\n")
         count++
      }
   }
   log.Println("count", count)
}

func TestZero(t *testing.T) {
   data, err := get_frontend()
   if err != nil {
      t.Fatal(err)
   }
   var data1 bytes.Buffer
   err = json.Indent(&data1, data, "", " ")
   if err != nil {
      t.Fatal(err)
   }
   err = os.WriteFile(name, data1.Bytes(), os.ModePerm)
   if err != nil {
      t.Fatal(err)
   }
}

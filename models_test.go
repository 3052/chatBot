package models

import (
   "bytes"
   "encoding/json"
   "fmt"
   "log"
   "os"
   "slices"
   "strings"
   "testing"
)

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
   slices.SortFunc(front.Models, func(a, b *metadata) int {
      return strings.Compare(a.Slug, b.Slug)
   })
   for _, one_metadata := range front.Models {
      fmt.Println(one_metadata.Slug)
   }
}

const name = "ignore/chatBot.json"

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

package models

import (
   "bytes"
   "encoding/json"
   "log"
   "os"
   "slices"
   "testing"
)

const name = "ignore/chatBot.json"

func TestRequestRead(t *testing.T) {
   log.SetFlags(log.Ltime)
   // A
   data, err := os.ReadFile(name)
   if err != nil {
      t.Fatal(err)
   }
   var modelsA models
   err = modelsA.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   modelsA = slices.DeleteFunc(modelsA, delete_model)
   // B
   data, err = find()
   if err != nil {
      t.Fatal(err)
   }
   var modelsB models
   err = modelsB.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   modelsB = slices.DeleteFunc(modelsB, delete_model)
   for _, modelA := range modelsA {
      if !modelsB.contains(modelA) {
         log.Println("removed", modelA)
      }
   }
   for _, modelB := range modelsB {
      if !modelsA.contains(modelB) {
         log.Println("added", modelB)
      }
   }
}

func TestRequestWrite(t *testing.T) {
   data, err := find()
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

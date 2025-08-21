package chatBot

import (
   "os"
   "testing"
)

func Test(t *testing.T) {
   resp, err := find()
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Create("chatBot.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   _, err = file.ReadFrom(resp.Body)
   if err != nil {
      t.Fatal(err)
   }
}

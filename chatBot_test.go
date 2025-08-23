package chatBot

import (
   "os"
   "testing"
)

func TestMarshal(t *testing.T) {
   data, err := get_models(0)
   if err != nil {
      t.Fatal(err)
   }
   err = os.WriteFile("ignore/chatBot.json", data, os.ModePerm)
   if err != nil {
      t.Fatal(err)
   }
}

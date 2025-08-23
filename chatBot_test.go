package chatBot

import (
   "fmt"
   "log"
   "testing"
)

func Test(t *testing.T) {
   models, err := get_models()
   if err != nil {
      t.Fatal(err)
   }
   for _, modelVar := range models {
      fmt.Printf("%+v\n", modelVar)
   }
   log.Print(len(models))
}

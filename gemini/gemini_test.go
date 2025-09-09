package gemini

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "os/exec"
)

func Test(t *testing.T) {
   key, err := exec.Command("credential", "-k", "GEMINI_API_KEY").Output()
   if err != nil {
      t.Fatal(err)
   }
   var promptVar prompt
   promptVar.Contents = []content{
      {
         Parts: []part{
            { Text: "2 + 3" },
         },
      },
   }
   completions, err := promptVar.generate(string(key))
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", completions)
}

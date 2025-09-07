package gemini

import (
   "os"
   "testing"
)

func Test(t *testing.T) {
   os.Getenv("GEMINI_API_KEY")
   exec.Command("credential", "google.com", "srpen6@gmail.com").Output()
   resp, err := generate("2 + 3")
   if err != nil {
      t.Fatal(err)
   }
   err = resp.Write(os.Stdout)
   if err != nil {
      t.Fatal(err)
   }
}

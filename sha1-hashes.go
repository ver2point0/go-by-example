// compute identities with SHA1 hashes

package main

import "crypto/sha1"
import "fmt"

func main() {
  
  s := "sha1 this string"
  
  // pattern for new hash: sha1.New(), sha1.Write(bytes), sha1.Sum([]byte{})
  // start with new hash
  h := sha1.New()
  
  // coerce to bytes
  h.Write([]byte(s))
  
  // get finalized hash result as a btye slice
  bs := h.Sum(nil)
  
  // print sha1 value in hex
  fmt.Println(s)
  fmt.Printf("%x\n", bs)
}
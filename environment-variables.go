// environment variables are a universal mechanism for conveying configuration
// information to UNIX programs

package main

import "os"
import "strings"
import "fmt"

func main() {
  
  // to set a key/value pair, use os.Setenv
  // to get a value for a key, os.Getenv
  os.Setenv("FOO", "1")
  fmt.Println("FOO:", os.Getenv("FOO"))
  fmt.Println("BAR:", os.Getenv("BAR"))
  
  fmt.Println()
  // use os.Environ to list all key/value pairs in the environment
  // this returns a slice in the form KEY=value
  // You can strings.Split them to get the key and value
  // Here we print all they keys
  for _, e := range os.Environ() {
    pair := strings.Split(e, "=")
    fmt.Println(pair[0])
  }  
}
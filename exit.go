// use os.Exit to immediately exit with a given status

package main

import "fmt"
import "os"

func main() {
  
  defer fmt.Println("!")
  os.Exit(3)
}
// Go offers excellent support for string formatting in the printf tradition
// common example string formatting tasks

package main

import "fmt"
import "os"

type point struct {
  x, y int
}

func main() {
  
  p := point{1, 2}
  // print an instance of our point struct
  fmt.Printf("%v\n", p)
  
  // if struct value, %+v variant will inlcude struct's field names
  fmt.Printf("%+v\n", p)
  
  // %#v variant prints a Go syntax representation of value
  fmt.Printf("%#v\n", p)
  
  // can format+print to io.Writers other than os.Stdout using Fprintf
  fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
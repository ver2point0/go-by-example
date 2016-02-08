// used to ensure that a function call is performed
// later in a program's execution, usually for
// cleanup
package main

import "fmt"
import "os"

func main() {
  
  // after creating a file, defer closing with closeFiile
  // it will be executed at end of enclosing function (main)
  // after writeFile has finished
  f := createFile("/workspace/code/defer.txt")
  defer closeFile(f)
  writeFile(f)
}

// creating a file
func createFile(p string) *os.File {
  fmt.Println("creating")
  f, err := os.Create(p)
  if err != nil {
    panic(err)
  }
  return f
}

// writing to it
func writeFile(f *os.File) {
  fmt.Println("writing")
  fmt.Fprintln(f, "data")
}

// closing it
func closeFile(f *os.File) {
  fmt.Println("closing")
  f.Close()
}
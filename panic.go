// use panics to fail fast on errors that shouldn't
// occur during normal operation, or that we aren't 
// prepared to handle gracefully
package main

import "os"

func main() {
  // use panic to check for unexpected errors
  panic("There is a problem!")
  
  // common use of panic is to abort if a function
  // returns an error value that we don't know
  // how to (or want to ) handle
  // panicking example if unexpected error when
  // creating a new file
  _, err := os.Create("/tmp/file")
  if err != nil {
    panic(err)
  }
}

// running panic.go will cause it to panic,
// print an error message and goroutine traces,
// and exit with a non-zero status
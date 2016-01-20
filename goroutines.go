// lightweight execution thread

package main

import "fmt"

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  // calling a function, running it synchronously
  f("direct")
  
  // execute function concurrently with calling one
  go f("goroutine")
  
  // start a goroutine for an anonymous function call
  go func(msg string) {
    fmt.Println(msg)
  }("going")
  
  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
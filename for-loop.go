// Go's only looping construct

package main

import "fmt"

func main() {
  
  // single condition
  i := 1
  for i <= 3{
    fmt.Println(i)
    i = i + 1
  }
  
  // initial/condition/after 
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }
  
  // without condition, loop repeatedly unless break loop or return from function
  for {
    fmt.Println("loop")
    break
  }
}
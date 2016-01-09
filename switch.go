package main

import "fmt"
import "time"

func main() {
  
  // basic switch
  i := 2
  fmt.Print("write ", i, " as ")
  switch i {
    case 1:
    fmt.Println("one")
    case 2:
    fmt.Println("two")
    case 3:
    fmt.Println("three")
  }
  
  // use commas to separate multiple expressions in same case statement
  // default case is optional
  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("it's the weekend")
    default:
      fmt.Println("it's a weekday")
  }
  
  // switch without an expression is an alternate way to use if/else logic
  // case expressions can be non-constants
  t := time.Now()
  switch {
    case t.Hour() < 12:
      fmt.Println("it's before noon")
    default:
      fmt.Println("it's after noon")
  }
  
}
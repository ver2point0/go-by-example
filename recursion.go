package main

import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n - 1)
}

func main() {
  
  fmt.Print("Factorial 0: ")
  fmt.Println(fact(0))
  fmt.Print("Factorial 1: ")
  fmt.Println(fact(1))
  fmt.Print("Factorial 3: ")
  fmt.Println(fact(3))
  fmt.Print("Factorial 5: ")
  fmt.Println(fact(5))
  fmt.Print("Factorial 7: ")
  fmt.Println(fact(7))
}
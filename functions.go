package main

import "fmt"

// explicit returns are needed
func plus(a int, b int) int {
  return a + b
}

// OPTIONAL: 
// omit type name for like-typed parameters up to final parameter that
// declares the type
func plusPlus(a, b, c int) int {
  return a + b + c
}

func main() {
  
  sum := plus(1, 2)
  fmt.Println("1 + 2 =", sum)
  
  sum = plusPlus(1, 2, 3)
  fmt.Println("1 + 2 + 3 =", sum)
}
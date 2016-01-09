package main

import "fmt"

func main() {
  
  // if with else
  if 7 % 2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }
  
  // if without else
  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }
  
  // if, else if, else with a statement
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
  
  // parentheses are not needed around conditions, but braces are required
  
  // ternary if does not exist in Go, full if statements are needed
  
  
  
  
}
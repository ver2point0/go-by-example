package main

import "fmt"

func main() {
  
  s := make([]string, 3)
  fmt.Println("emp: ", s)
  
  s[0] = "1"
  s[1] = "2"
  s[2] = "3"
  fmt.Println("set: ", s)
  fmt.Println("get: ", s[2])
  
  fmt.Println("len: ", len(s))
  
  s = append(s, "4")
  s = append(s, "5", "6")
  fmt.Println("apd: ", s)
  
  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy: ", c)
  
  l := s[2:5]
  fmt.Println("sl1: ", l)
  
  l = s[:5]
  fmt.Println("sl2: ", l)
  
  l = s[2:]
  fmt.Println("sl3: ", l)
  
  t := []string{"7", "8", "9"}
  fmt.Println("dcl: ", t)
  
  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  
  fmt.Println("2d: ", twoD)
}
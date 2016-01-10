package main

import "fmt"

func main() {
  
  var a [5]int
  fmt.Println("emp: ", a)
  
  a[0] = 100
  a[4] = 2
  // set and get array values
  fmt.Println("set: ", a)
  fmt.Println("get: ", a[0])
  fmt.Println("get: ", a[4])
  
  // get length of array with len
  fmt.Println("len: ", len(a))
  
  // declare and initialize in one step
  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl: ", b)
  
  // 2D array
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
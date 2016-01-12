package main

import "fmt"

func main() {
  
  // use range to sum numbers in a slice. arrays work like this too.
  // ignored index with _(blank) identifier
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)
  
  // need index in this case because we are looking for a specific one
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }
  
  // range on map iterates over k/v pairs
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }
  
  // iterate over Unicode code points on strings using range
  // i: starting byte index of the rune
  // c: the rune itself
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
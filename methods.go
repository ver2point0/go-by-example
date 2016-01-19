package main

import "fmt"

type rect struct {
  width, height int
}

// pointer receiver, notice the asterisk (*)
func (r *rect) area() int {
  return r.width * r.height
}

// value receiver
func (r rect) perim() int {
  return 2 * r.width + 2 * r.height 
}

func main() {
  // Go automatically handles conversion between values and pointers 
  // for method calls
  
  r := rect{width: 100, height: 20}
  
  fmt.Println("area: ", r.area())
  fmt.Println("perim: ", r.perim())
  
  rp := &r
  fmt.Println("area: ", rp.area())
  fmt.Println("perim: ", rp.perim())
}
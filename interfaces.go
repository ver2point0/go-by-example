package main

import "fmt"
import "math"

type geometry interface {
  area() float64
  perim() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

// to implement an interface in Go: implement all methods in the interface

// rectangle methods
func (r rect) area() float64 {
  return r.width * r.height
}
func (r perim) perim() float64 {
  return 2 * r.width + 2 * r.height
}

//circle methods
func (c circle) area() float64 {
  return math.PI * c.radius * c.radius
}
func (c circle) perim() float64 {
  return 2 * math.PI * c.radius
}

// generic measure function works on any geometry type
func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}

func main() {
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}
  
  measure(r)
  measure(c)
}


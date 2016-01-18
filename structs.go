package main

import "fmt"

type person struct {
  name string
  age int
}

func main() {
  
  // create a new struct
  fmt.Println(person{"Bob", 20})
  
  // can name fields when initializing a struct
  fmt.Println(person{name: "Alice", age: 30})
  
  // omitted fields are zero valued
  fmt.Println(person{name: "Fred"})
  
  // An & prefix yields a pointer to the struct
  fmt.Println(&person{name: "Ann", age: 40})
  
  // access struct fields with a dot
  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)
  
  // can use dots with struct pointers - the pointers are automatically dereferenced
  sp := &s
  fmt.Println(sp.age)
  
  // structs are mutable
  sp.age = 51
  fmt.Println(sp.age)
}
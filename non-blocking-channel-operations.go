// basic sends and receives on channels are blocking
// use select with a default clause to implement
// non-blocking sends, receives, and non-blocking 
// multi-way selects
package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan bool)
  
  // non-blocking receive
  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  default:
    fmt.Println("no message received")
  }
  
  // non-blocking send
  msg := "hi"
  select {
  case messages <- msg:
    fmt.Println("sent message", msg)
  default:
    fmt.Println("no message sent")
  }
  
  // use multiple cases above default clause to implement
  // a multi-way non-blocking select
  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  case sig := <-signals:
    fmt.Println("received signal", sig)
  default:
    fmt.Println("no activity")
  }
}
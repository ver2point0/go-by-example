// select lets you wait on multiple channel operations
// combining goroutines and channels with select is powerful

package main

import "time"
import "fmt"

func main() {
  
  // select across channels c1 and c2
  c1 := make(chan string)
  c2 := make(chan string)
  
  // each channel will receive a value after some amount of time,
  // to simulate blocking remote procedure call (RPC)
  // operations executing in concurrent goroutines
  go func() {
    time.Sleep(time.Second * 1)
    c1 <- "one"
  }()
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "two"
  }()
  
  for i := 0; i < 2; i++ {
    select {
    case msg1 := <-c1:
      fmt.Println("received", msg1)
    case msg2 := <-c2:
      fmt.Println("received", msg2)
    }
  } 
  
}


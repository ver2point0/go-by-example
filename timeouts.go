// timeouts are important for programs that connect to external resources
// or that otherwise need to bound execution time

package main

import "time"
import "fmt"

func main() {
  
  // execute an external call that returns 
  // its result on channel c1 after 2 seconds
  c1 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 2)
    c1 <- "result 1"
  }()
  
  // select is implementing a timeout
  // res := <-c1 awaits the result
  // <-time.After awaits a value to be sent after
  // the timeout of 1 second
  select {
  case res := <-c1:
    fmt.Println(res)
  case <-time.After(time.Second * 1):
    fmt.Println("timeout 1")
  }
  
  // execute an external call that returns 
  // its result on channel c2 after 2 seconds
  c2 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "result 2"
  }()
  // allow a longer timeout of 3 seconds
  // receive from channel c2 will succeed
  select {
  case res := <-c2:
    fmt.Println(res)
  case <-time.After(time.Second * 3):
    fmt.Println("timeout 2")
  }

}
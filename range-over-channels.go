// use for and range to iterate over 
// values received from a channel
package main

import "fmt"

func main() {
  
  // 2 values added to the queue channel
  queue := make(chan string, 2)
  queue <- "one"
  queue <- "two"
  close(queue)
  
  // iterate over each element as it's received from queue
  // closing the channel means the iteration stops
  // after receiving the 2 elements
  for elem := range queue {
    fmt.Println(elem)
  }
}
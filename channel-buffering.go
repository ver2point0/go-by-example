// buffered channels accept a limited number of values without 
// a corresponding receiver for those values

package main

import "fmt"

func main() {
  
  // make a channel of strings buffering up to 2 values
  messages := make(chan string, 2)
  
  // can send these values into channel without corresponding concurrent receivers
  messages <- "buffered"
  messages <- "channel"
  
  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
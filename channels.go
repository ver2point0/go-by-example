package main

import "fmt"

func main() {
  
  // create new channel with make(chan val-type)
  messages := make(chan string)
  
  // send a value into a channel using the channel <- syntax
  // send "ping" to the messages channel made above from a new goroutine
  go func() {
    messages <- "ping"
  }()
  
  // receive a value from a channel using the <-channel syntax
  // receive "ping" message sent above and print it out
  msg := <-messages
  fmt.Println(msg)
}
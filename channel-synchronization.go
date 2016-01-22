// use channels to synchronize execution across goroutines

package main

import "fmt"
import "time"

func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")
  
  // send value to notify we're done
  done <- true
}

func main() {
  
  // start worker goroutine, giving it the channel to notify on
  done := make(chan bool, 1) 
  go worker(done)
  
  // block until we receive a notification from the worker on the channel
  <-done
}
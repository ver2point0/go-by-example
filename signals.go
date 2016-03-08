// how to handle signals in Go with channels

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"


func main() {
  
  // create channel to receive notifications
  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)
  
  // signal.Notify registers the given channel to receive notifications
  // of specified signals
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  
  // goroutine that executes a blocking receive for signals
  go func() {
    sig := <-sigs
    fmt.Println()
    fmt.Println(sig)
    done <- true
  }()
  
  // program will wait until it gets expected signal and then exit
  fmt.Println("awaiting signal")
  <-done
  fmt.Println("exiting")
}
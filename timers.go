// use timers to run execute Go code in the future
// or repeatedly at some interval
package main

import "time"
import "fmt"

func main() {
  
  // timer1 waits 2 seconds
  timer1 := time.NewTimer(time.Second * 2)
  
  // <-timer1.C blocks on the timer's channel C
  // until it sends a value indicating the time expired
  <-timer1.C
  fmt.Println("timer 1 expired")
  
  // cancel a timer before it expires
  timer2 := time.NewTimer(time.Second)
  go func() {
    <-timer2.C
    fmt.Println("timer 2 expired")
  }()
  stop2 := timer2.Stop()
  if stop2 {
    fmt.Println("timer 2 stopped")
  }
}
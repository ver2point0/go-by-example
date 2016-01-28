// use tickers to do something 
// repeatedly at regular intervals
package main

import "time"
import "fmt"

func main() {
  
  // iterate over the values every 500ms
  ticker := time.NewTicker(time.Millisecond * 500)
  go func() {
    for t := range ticker.C {
      fmt.Println("tick at", t)
    }
  }()
  
  // stop ticker at 1600ms
  time.Sleep(time.Millisecond * 1600)
  ticker.Stop()
  fmt.Println("ticker stopped")
}
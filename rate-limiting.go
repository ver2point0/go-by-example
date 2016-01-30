// control resource utilization and maintain qaulity of service
// support rate limiting with goroutines, channels, tickers

package main 

import "time"
import "fmt"

func main() {
  
  // limit handling of incoming requests
  requests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests <- i
  }
  close(requests)
  
  // receive a value every 200 ms
  limiter := time.Tick(time.Millisecond * 200)
  
  // limit ourselves to 1 request every 200 ms
  for req := range requests {
    <-limiter
    fmt.Println("request", req, time.Now())
  }
  
  // allows burst of up to 3 events in our
  // rate limiting scheme
  burstyLimiter := make(chan time.Time, 3)
  
  // fill channel to represent allowed bursting
  for i := 0; i < 3; i++ {
    burstyLimiter <- time.Now()
  }
  
  // add new value, up to 3, to burstyLimiter every 200 ms
  go func() {
    for t := range time.Tick(time.Millisecond * 200) {
      burstyLimiter <- t
    }
  }()
  
  // simulate 5 more incoming requests
  // first 3 will benefit from burstyLimiter
  burstyRequests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    burstyRequests <- i
  }
  close(burstyRequests)
  for req := range burstyRequests {
    <-burstyLimiter
    fmt.Println("request", req, time.Now())
  }
  
}

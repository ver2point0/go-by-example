// closing a channel indicates that 
// no more values will be sent on it
// useful to communicate completion
// to channel's receivers
package main

import "fmt"

func main() {
  // use jobs channel to communicate work to be done
  // from main() goroutine to worker goroutine
  // when no more jobs for worker close jobs channel
  jobs := make(chan int, 5)
  done := make(chan bool)
  
  // worker goroutine:
  // repeatedly receives from jobs with j, more := <-jobs
  // in this 2-value form of receive, more will be false
  // if jobs has been closed and all values in channel
  // have already been received
  go func() {
    for {
      j, more := <-jobs
      if more {
        fmt.Println("received job", j)
      } else {
        fmt.Println("received all jobs")
        done <- true
        return
      }
    }
  }()
  
  // sends 5 jobs to the worker over jobs channel
  // then closes it
  for j := 1; j <= 5; j++ {
    jobs <- j
    fmt.Println("sent job", j)
  }
  close(jobs)
  fmt.Println("sent all jobs")
  
  // await worker using synchronization approach
  <-done
}
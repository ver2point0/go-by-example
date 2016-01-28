// implement a worker pool using
// goroutines and channels
package main

import "fmt"
import "time"

// run several concurrent instances
// receive work on jobs channel
// send results on results channel
// sleep to simulate an expensive task
func worker(id int, jobs <-chan int, results chan<- int) {
  for j := range jobs {
    fmt.Println("worker", id, "processing job", j)
    time.Sleep(time.Second)
    results <- j * 2
  }
}

func main() {
  
  // make 2 channels
  // to send work and collect results
  jobs := make(chan int, 100)
  results := make(chan int, 100)
  
  // starts up 3 workers initially blocked
  // because there are no jobs yet
  for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
  }
  // send 9 jobs, then close channel
  for j := 1; j <= 9; j++ {
    jobs <- j
  }
  close(jobs)
  
  // collect all results of work
  for a := 1; a <= 9; a++ {
    <-results
  }
}
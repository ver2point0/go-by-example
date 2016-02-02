// use built-in synchronization features of goroutines and channels
// to access share state across multiple goroutines
package main

import (
  "fmt"
  "math/rand"
  "sync/atomic"
  "time"
)

// state is owned by a single goroutine
// guarantee that data is never corrupted with 
// concurrent access
// readOp and writeOp structs encapsulate
// sends and receives
type readOp struct {
  key int
  resp chan int
}
type writeOp struct {
  key int
  val int
  resp chan bool
}

func main() {
  
  // count how many operations performed
  var ops int64 = 0
  
  // reads and writes channels will be used by other 
  // goroutines to issue read/write requests, respectively
  reads := make(chan *readOp)
  writes := make(chan *writeOp)
  
  // owns state, private to stateful goroutine
  // repeatedly selects on reads and writes channels
  go func() {
    var state = make(map[int]int)
    for {
      select {
      case read := <-reads:
        read.resp <- state[read.key]
      case write := <-writes:
        state[write.key] = write.val
        write.resp <- true
      }
    }
  }()
  
  // starts 100 goroutines to issue reads to state-owning
  // goroutine via reads channel
  for r := 0; r < 100; r++ {
    go func() {
      for {
        read := &readOp{
          key: rand.Intn(5),
          resp: make(chan int)}
        reads <- read
        <-read.resp
        atomic.AddInt64(&ops, 1)
      }
    }()
  }
  
  // start 10 writes
  for w := 0; w < 10; w++ {
    go func() {
      for {
        write := &writeOp {
          key: rand.Intn(5),
          val: rand.Intn(100),
          resp: make(chan bool)}
        writes <- write
        <-write.resp
        atomic.AddInt64(&ops, 1)
      }
    }()
  }
  
  // let goroutines work for a second
  time.Sleep(time.Second)
  
  // finally, capture and report ops count
  opsFinal := atomic.LoadInt64(&ops)
  fmt.Println("ops:", opsFinal)
}
// use mutexes to safely access data across multiple goroutines

package main

import (
  "fmt"
  "math/rand"
  "runtime"
  "sync"
  "sync/atomic"
  "time"
)

func main() {
  
  // state will be a map
  var state = make(map[int]int)
  
  // mutex will synchronize across to state
  var mutex = &sync.Mutex{}
  
  // ops will count how many operations we perform against the state
  var ops int64 = 0
  
  // start 100 goroutines to execute repeated reads against the state
  for r := 0; r < 100; r++ {
    go func() {
      total := 0
      // for each read pick a key to access
      // Lock() the mutex to ensure exclusive access to state
      // read value at chosen key
      // Unlock() mutex
      // increment ops count
      for {
        key := rand.Intn(5)
        mutex.Lock()
        total += state[key]
        mutex.Unlock()
        atomic.AddInt64(&ops, 1)
        // explicitly yield after each operation
        // to ensure goroutine doesn't starve scheduler
        runtime.Gosched()
      }
    }()
  }
  
  // start 10 goroutines to simulate writes 
  for w := 0; w < 10; w++ {
    go func() {
      for {
        key := rand.Intn(5)
        val := rand.Intn(100)
        mutex.Lock()
        state[key] = val
        mutex.Unlock()
        atomic.AddInt64(&ops, 1)
        runtime.Gosched()
      }
    }()
  }
  
  // allow 10 goroutines to work on state and mutex for 1 second
  time.Sleep(time.Second)
  
  // take and report final ops count
  opsFinal := atomic.LoadInt64(&ops)
  fmt.Println("ops:", opsFinal)
  
  // final lock state, show ending point
  mutex.Lock()
  fmt.Println("state:", state)
  mutex.Unlock()
}
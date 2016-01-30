// manage state using sync/atomic package for
// atomic counters accessed by mulitple goroutines

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {
  
  // use unsigned integer to represent counter (always positive)
  var ops uint64 = 0
  
  // simulate concurrent updates
  // start 50 goroutines that each increment counter
  // once a millisecond
  for i := 0; i < 50; i++ {
    go func() {
      for {
        // to atomically increment the counter
        // use AddUint64, giving memory address 
        // of ops with & syntax
        atomic.AddUint64(&ops, 1)
        // allow other goroutines to proceed
        runtime.Gosched()
      }
    }()
  }
  
  // wait 1 second to allow some ops to accumulate
  time.Sleep(time.Second)
  
  // extract a copy of current value into opsFinal via LoadUint64
  // to safely use counter while it is being updated
  // by other goroutines
  opsFinal := atomic.LoadUint64(&ops)
  fmt.Println("ops:", opsFinal)
}
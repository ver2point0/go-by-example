// Go offers extensive support for times and durations
// Here are some examples

package main

import "fmt"
import "time"

func main() {
  
  p := fmt.Println
  
  // get current time
  now := time.Now()
  p(now)
  
  // build time struct
  // times are always associated with a location(i.e. time zone)
  then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
  p(then)
  
  // extract time components of "then"
  p(then.Year())
  p(then.Month())
  p(then.Day())
  p(then.Hour())
  p(then.Minute())
  p(then.Second())
  p(then.Nanosecond())
  p(then.Location())
  
  // get weekday
  p(then.Weekday())
  
  // compare times "then" and "now"
  p(then.Before(now))
  p(then.After(now))
  p(then.Equal(now))
  
  // find duration between "then" and "now"
  diff := now.Sub(then)
  p(diff)
  
  // compute duration length in various units
  p(diff.Hours())
  p(diff.Minutes())
  p(diff.Seconds())
  p(diff.Nanoseconds())
  
  // advance "then" and move backwards
  p(then.Add(diff))
  p(then.Add(-diff))
}
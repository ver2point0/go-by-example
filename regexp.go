// Go offers built-in support for regular expressions
// Here are common regexp tasks

package main

import "bytes"
import "fmt"
import "regexp"

func main() {
  
  // tests whether a pattern matches a string
  match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
  fmt.Println(match)
  
  // compile an optimized regexp struct
  r, _ := regexp.Compile("p([a-z]+)ch")
  
  // match test
  fmt.Println(r.MatchString("peach"))
  
  // finds match for regexp
  fmt.Println(r.FindString("peach punch"))
    
  // finds match but returns start and end indexes for match
  fmt.Println(r.FindStringIndex("peach punch"))
  
  // submatch variants include info about both
  // whole-pattern matches and submatches within
  // those matches
  fmt.Println(r.FindStringSubmatch("peach punch"))
  
  fmt.Println(r.FindStringSubmatchIndex("peach punch"))
  
  fmt.Println(r.FindAllString("peach punch pinch", -1))
  
  fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
  
  fmt.Println(r.FindAllString("peach punch pinch", 2))
  
  fmt.Println(r.Match([]byte("peach")))
  
  r = regexp.MustCompile("p([a-z]+)ch")
  
  fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
  
  in := []byte("a peach")
  out := r.ReplaceAllFunc(in, bytes.ToUpper)
  fmt.Println(string(out))
}
 
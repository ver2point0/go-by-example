// standard library's strings package provides many
// useful string-related functions
package main

import s "strings"
import "fmt"

// alias fmt.Println to use a short name
var p = fmt.Println

// sample of functions available in strings
// these are functions from package not 
// methods on string object itself
// need to pass string in questions as 1st argument
func main() {
  p("Contains:  ", s.Contains("test", "es"))
  p("Count: ", s.Count("test", "t"))
  p("HasPrefix: ", s.HasPrefix("test", "te"))
  p("HasSuffix: ", s.HasSuffix("test", "st"))
  p("Index: ", s.Index("test", "e"))
  p("Join:      ", s.Join([]string{"a", "b"}, "-"))
  p("Repeat: ", s.Repeat("a", 5))
  p("Replace: ", s.Replace("foo", "o", "0", -1))  
  p("Replace: ", s.Replace("foo", "o", "0", 1))
  p("Split: ", s.Split("a-b-c-d-e", "-"))
  p("ToLower: ", s.ToLower("TEST"))
  p("ToUpper: ", s.ToUpper("test"))
  p()
  p("Len: ", len("hello"))
  p("Char:", "hello"[1])
}
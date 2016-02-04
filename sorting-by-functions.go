// customer sorting in go
// sort collections by something
// other than natural order
package main

import "sort"
import "fmt"

// create ByLength type
type ByLength []string

// implement sort.Interface
// Len, Less, Swap to user sort package's generic sort function
// sort in order of increasing string length
func (s ByLength) Len() int {
  return len(s)
}
func (s ByLength) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
  return len(s[i]) < len(s[j])
}

// cast original fruits slice to ByLength
// use sort.Sort on typed slice
func main() {
  fruits := []string{"peach", "banana", "kiwi"}
  sort.Sort(ByLength(fruits))
  fmt.Println(fruits)
}
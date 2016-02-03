// sorting for builtins and user-defined types
package main

import "fmt"
import "sort"

func main() {
  // sort methods are specific to the builtin type
  // sorting is in-place, changes give slice
  // does not return a new one
  // sorting strings
  strs := []string{"c", "a", "b"}
  sort.Strings(strs)
  fmt.Println("Strings:", strs)
  
  // sorting ints
  ints := []int{7, 2, 4}
  sort.Ints(ints)
  fmt.Println("Ints: ", ints)
  
  // can use sort to check if a slice has been sorted
  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted: ", s)
}
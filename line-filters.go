// a line filter is a common type of program that reads
// input on stdin, processes it, and then prints some derived 
// result to stdout. grep and sed are common line filters

package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  
  scanner := bufio.NewScanner(os.Stdin)
  
  for scanner.Scan() {
    ucl := strings.ToUpper(scanner.Text())
    fmt.Println(ucl)
  }
  
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "error:", err)
    os.Exit(1)
  }
}
// use exec function to replace current Go process with another one

package main

import "syscall"
import "os"
import "os/exec"

func main() {
  
  binary, lookErr := exec.LookPath("ls")
  if lookErr != nil {
    panic(lookErr)
  }
  
  args := []string{"ls", "-a", "-l", "-h"}
  
  env := os.Environ()
  
  execErr := syscall.Exec(binary, args, env)
  if execErr != nil {
    panic(execErr)
  }
}
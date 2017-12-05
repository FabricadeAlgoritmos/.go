package main

import "fmt"

func fibonacci(x uint) uint {
  if x == 0 || x == 1 {
    return 1
  }
  return fibonacci(x - 1) + fibonacci(x - 2)
}

func main() {
  var i uint
  for i = 0; i <= 10; i++ {
    fmt.Printf("%d\t", fibonacci(i))
  }
}

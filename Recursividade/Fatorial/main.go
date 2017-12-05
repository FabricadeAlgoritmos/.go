package main

import "fmt"

func fatorial(x uint) uint{
  if x == 0 || x == 1 {
    return 1
  }
    return x * fatorial(x - 1)
}

func main(){

  var i uint
  for i = 0; i <= 5; i++ {
    fmt.Printf("%d! = %d\n", i, fatorial(i))
  }

}

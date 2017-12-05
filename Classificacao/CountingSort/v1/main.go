package main

import "fmt"

func main(){
  vet := [5]int{5, 4, 3, 2, 1}
  var count [5]int
  var saida [5]int
  fim := 4
  for i := 0; i < fim; i++ {
    for j := i + 1; j <= fim; j++ {
      if vet[i] > vet[j] {
        count[i]++
      }else{
        count[j]++
      }
    } // fim for j
  } // fim for i

  for i := 0; i <= fim; i++ {
    saida[count[i]] = vet[i]
  }

  fmt.Println(saida)
}

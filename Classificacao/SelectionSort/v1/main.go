package main

import "fmt"

func main(){
  vetor := [5]int{5, 4, 3, 2, 1}
  for i := 0; i < len(vetor) - 1; i++ {
    menor := i
    for j := i + 1; j < len(vetor); j++ {
      if vetor[menor] > vetor[j] {
        menor = j
      }
    }
    aux := vetor[i]
    vetor[i] = vetor[menor]
    vetor[menor] = aux
  }

  fmt.Println(vetor)
}

package main

import "fmt"

const MAX = 16

func main(){
  vet := [MAX]int{32, 18, 45, 29, 16, 41, 30, 27 , 14, 37, 29, 38, 20, 32, 17, 26}

  inc := [4]int{8, 4, 2, 1}

  x := 0
  for inc[x] > 1 {
    i := 0
    j := inc[x]
    for j < MAX {
      if vet[i] > vet[j] {
        aux := vet[i]
        vet[i] = vet[j]
        vet[j] = aux
      }
      j++
      i++
    }//fim for
    x++
  }

  Insertion(&vet)

  fmt.Println(vet)
}

func Insertion(vet* [MAX]int){
  for i := 1; i < len(vet); i++ {
      j := i
      for vet[j] < vet[j - 1] && j > 0 {
        aux := vet[j - 1]
        vet[j - 1] = vet[j]
        vet[j] = aux
        j--
      }
  }
}

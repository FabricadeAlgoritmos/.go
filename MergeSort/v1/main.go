package main

import "fmt"

func main(){
  vet1 := [5]int{1, 3, 5, 7, 9}
  vet2 := [5]int{2, 4, 6, 8, 10}
  //fmt.Println(vet1)
  //fmt.Println(vet2)

  var i, j, k int = 0, 0, 0
  var vet3 [10]int
  for i < len(vet1) && j < len(vet2) {
    if vet1[i] < vet2[j] {
      vet3[k] = vet1[i]
      i++
    }else{
      vet3[k] = vet2[j]
      j++
    }
    k++
  }// fim for

  for i < len(vet1) {
    vet3[k] = vet1[i]
    i++
    k++
  }

  for j < len(vet2) {
    vet3[k] = vet2[j]
    j++
    k++
  }

  fmt.Println(vet3)
}

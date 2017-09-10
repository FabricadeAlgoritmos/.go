package main

import "fmt"

const T int = 12

func main(){

  vet := [T]int{26, 19, 34, 19, 23, 16, 31, 26, 31, 22, 21, 29}
  //fmt.Println(vet)
  var saida [T]int
  var maior int = getMaior(vet)
  menor := getMenor(vet)
  //fmt.Println("Maior: ", maior, "\nMenor: ", menor)

  cont := make([]int, maior + 1)
  for i := menor; i <= maior; i++ {
    cont[i] = 0
  }

  for i := 0; i < T; i++{
    cont[vet[i]]++
  }

  for i := menor + 1; i <= maior; i++{
    cont[i] += cont[i - 1];
  }

  for j := 0; j < T; j++ {
    i := cont[vet[j]]
    saida[i - 1] = vet[j]
    cont[vet[j]]--
  }

  fmt.Println(saida);

}

func getMaior(vet [T]int)int{
  var m int = vet[0]
  for i := 1; i < len(vet); i++ {
    if(vet[i] > m){
      m = vet[i]
    }
  }
  return m
}

func getMenor(vet [T]int)int{
  var m int = vet[0]
  for i := 1; i < len(vet); i++ {
    if(vet[i] < m){
      m = vet[i]
    }
  }
  return m
}

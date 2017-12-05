package main

import "fmt"

func main(){
  var vetor [6]int
  var fim, aux, i int
  fim = 0
  fmt.Print("Informe o valor ~:> ")
  fmt.Scan(&vetor[fim])
  fim++
  for fim < 6 {
    fmt.Print("Informe o valor ~:> ")
    fmt.Scan(&aux)
    i = fim
    for i > 0 && aux < vetor[i - 1] {
      vetor[i] = vetor[i - 1]
      i--
    }
    vetor[i] = aux
    fim++
  }

  fmt.Println(vetor);
}

package main

import (
  "fmt"
)

func main(){

  listaVelha := [7]int{1, 5, 25, 65, 98, 101, 220}
  lote := [4]int{2, 60, 79, 80}

  fmt.Print("Lista velha: ")
  fmt.Println(listaVelha, len(listaVelha), "elementos")

  fmt.Print("Lote: ")
  fmt.Println(lote, len(lote), "elementos")

  const SIZE int = (len(listaVelha) + len(lote))

  var novaLista [SIZE]int

  var i, j, k int = 0, 0, 0

  for i < len(listaVelha) && j < len(lote) {
    if listaVelha[i] < lote[j] {
      novaLista[k] = listaVelha[i]
      i++
    }else{
      novaLista[k] = lote[j]
      j++
    }
    k++
  }

  for i < len(listaVelha) {
    novaLista[k] = listaVelha[i]
    i++
    k++
  }

  for j < len(lote) {
    novaLista[k] = lote[j]
    j++
    k++
  }

  fmt.Println("\nInserindo dados...")
  fmt.Println("Nova lista: ", novaLista, len(novaLista), "elementos")

  fmt.Println("\nRemovendo dados...")
  novoLote := [5]int{1, 5, 65, 98, 220}
  fmt.Println("Lista Atual: ", novaLista)
  fmt.Println("Novo Lote: ", novoLote)
  const SIZE2 int = (len(novaLista) - len(novoLote))
  var listaAtualizada [SIZE2]int
  i = 0
  j = 0
  k = 0
  fmt.Println("Atualizando...")
  for j < len(novoLote) {
    if novaLista[i] != novoLote[j] {
      listaAtualizada[k] = novaLista[i]
      fmt.Printf("Lista [%d] recebe %d\n", k, novaLista[i])
      k++
    }else{
      fmt.Println(novoLote[j], "removido!")
      j++
    }
    i++
  }

  for i < len(novaLista){
    listaAtualizada[k] = novaLista[i]
    fmt.Printf("Lista [%d] recebe %d\n", k, novaLista[i])
    k++
    i++
  }
  fmt.Println("Lista Atualizada: ", listaAtualizada)

  fmt.Println("\nBuscando dados...")
  fmt.Println("Lista: ", novaLista)
  var cod int
  fmt.Print("Informe o valor: ")
  fmt.Scan(&cod)

  var inicio, fim int = 0, (len(novaLista) - 1)
  meio := (inicio + fim) / 2
  for cod != novaLista[meio] && inicio < fim {
    if(cod > novaLista[meio]){
      inicio = meio + 1
    }else{
      fim = meio - 1
    }
    meio = (inicio + fim) / 2
  }

  if cod == novaLista[meio] {
    fmt.Println(novaLista[meio], "encontrado na posição", meio)
  }else{
    fmt.Println("Valor não existe na lista!")
  }
}

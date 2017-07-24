package main

import (
  "fmt"
)

type Usuario struct {
  id int
  nome string
  status bool
}

type Indice struct {
  cod int
  pos int
}

func main(){
  var lista [10]Usuario
  var index [10]Indice

  lista[0].id = 50
  lista[0].nome = "Jane"
  lista[0].status = true

  lista[1].id = 10
  lista[1].nome = "Dan"
  lista[1].status = true

  lista[2].id = 36
  lista[2].nome = "Marry"
  lista[2].status = true

  var endList int = 2

  index[0].cod = 10
  index[0].pos = 1

  index[1].cod = 36
  index[1].pos = 2

  index[2].cod = 50
  index[2].pos = 0

  showUser(lista, endList)
  showIndex(index, endList)
  fmt.Println("Inserindo dados na lista...")
  Inserir(&lista, &index, &endList)
  fmt.Println("Um novo usuário foi adicionado!")
  showUser(lista, endList)
  showIndex(index, endList)

  fmt.Println("Buscando um Usuario na lista...")
  Busca(lista, index, endList)

  fmt.Println("\nRemovendo um Usuario...")
  Remover(&lista, &index, endList)
  showUser(lista, endList)

  fmt.Println("\nAtualizando Lista...")
  var novaLista [10]Usuario
  endList = Atualiza(lista, &novaLista, &index, &endList)
  Imprimir(novaLista, endList)
}

func Inserir(u* [10]Usuario, index* [10]Indice, fim * int){
  if *fim < 10 {
    *fim++
    fmt.Print("Informe o id: ")
    fmt.Scan(&u[*fim].id)
    fmt.Print("Informe o nome: ")
    fmt.Scan(&u[*fim].nome)
    u[*fim].status = true

    var aux int = *fim - 1
    for aux >= 0 && u[*fim].id < index[aux].cod {
      index[aux + 1].cod = index[aux].cod
      index[aux + 1].pos = index[aux].pos
      aux--
    }
    index[aux + 1].cod = u[*fim].id
    index[aux + 1].pos = *fim
  }
}

func Remover(u* [10]Usuario, index* [10]Indice, F int){
  var codigo int
  fmt.Print("Informe o codigo de Busca: ")
  fmt.Scan(&codigo)
  var inicio, final int = 0, F
  var meio int = ((inicio + final) / 2)
  for codigo != index[meio].cod && inicio < final {
    if codigo > index[meio].cod {
      inicio = meio + 1
    }else{
      final = meio - 1
    }
    meio = (inicio + final) / 2
  } // fim do for

  var i int = index[meio].pos
  if codigo == u[i].id {
    u[i].status = false
    fmt.Println(u[i].nome, "removido!")
  }else{
    fmt.Println("Codigo não existe na lista!")
  }
}

func Busca(u [10]Usuario, index [10]Indice, F int){
  var codigo int
  fmt.Print("Informe o codigo de Busca: ")
  fmt.Scan(&codigo)
  var inicio, final int = 0, F
  var meio int = ((inicio + final) / 2)
  for codigo != index[meio].cod && inicio < final {
    if codigo > index[meio].cod {
      inicio = meio + 1
    }else{
      final = meio - 1
    }
    meio = (inicio + final) / 2
  } // fim do for

  var i int = index[meio].pos
  if codigo == u[i].id && u[i].status {
    fmt.Println(u[i].nome, "encontrado na posição", i)
  }else{
    fmt.Println("Codigo não existe na lista!")
  }
}

func Atualiza(oldList [10]Usuario, newList* [10]Usuario,index* [10]Indice, fim * int)(int){
  var f int = 0
  for i := 0; i <= *fim; i++{
    if oldList[i].status {
      newList[f] = oldList[i]
      index[f].cod = newList[f].id
      index[f].pos = f
      f++
    }
  }
  return f
}


// Imprime apenas os Usuarios com status igual a true
func Imprimir(u [10]Usuario, F int){
  for i := 0; i <= F; i++{
    if u[i].status {
      fmt.Printf("[id=>%d,\tNome=>%s,\tStatus=>%b]\n", u[i].id, u[i].nome, u[i].status)
    }
  }
}


func showUser(u [10]Usuario, F int){
  fmt.Println("\nLista de Usuarios")
  for i := 0; i <= F; i++{
    fmt.Printf("[id=>%d,\tNome=>%s,\tStatus=>%b]\n", u[i].id, u[i].nome, u[i].status)
  }
}

func showIndex(index [10]Indice, F int){
  fmt.Println("\nArquivo de Indice")
  for i := 0; i <= F; i++{
    fmt.Printf("[cod=>%d,\tposicao=>%d]\n", index[i].cod, index[i].pos)
  }
}

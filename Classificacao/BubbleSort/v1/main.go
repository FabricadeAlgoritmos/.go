package main

import (
	"fmt"
)

func main(){
	vet := [5]int{5, 4, 3, 2, 1}
	var max, troca int = 4, 1
	for troca != 0 {
		troca = 0
		for i := 0; i <= max - 1; i++ {
			if vet[i] > vet[i + 1] {
				aux := vet[i]
				vet[i] = vet[i + 1]
				vet[i + 1] = aux
				troca = 1
			}
		}
		max--
	}

	fmt.Println(vet)
	
}
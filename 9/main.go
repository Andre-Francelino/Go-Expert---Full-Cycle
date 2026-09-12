package main

import "fmt"

// Aula Funções Variádicas

func main() {
	fmt.Println(soma(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543))
	fmt.Println(subtracao(3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543))
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func subtracao(numeros ...int) int {
	total := 11205
	for _, numero := range numeros {
		total -= numero
	}
	return total
}

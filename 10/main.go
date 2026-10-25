package main

import "fmt"

// Aula Funções Anônimas (Closures)

func main() {

	total := func() int {
		return soma(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543) * 2
	}()

	fmt.Println(soma(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543))
	fmt.Println(total)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}


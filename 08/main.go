package main

import (
	"errors"
	"fmt"
)

// Aula Funções

func main() {
	valor, err := sum(45, 10)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", valor)
}

func sum(a, b int) (int, error) {
	if a + b >= 50 {
		return 0, errors.New("A soma é maior ou igual a 50")
	}
	return a + b, nil
}

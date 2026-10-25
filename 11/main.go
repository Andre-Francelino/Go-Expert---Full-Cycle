package main

import "fmt"

// Aula Iniciando com Structs
// Obs: Go não é uma linguagem orientada a objetos. A linguagem possui o seu "Go Way" de programar.
// Struct não é uma classe mas é o que mais se assemelha com orientação a objetos

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {

	andre := Cliente {
		Nome: "André",
		Idade: 34,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", andre.Nome, andre.Idade, andre.Ativo)
	fmt.Println()

	andre.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", andre.Nome, andre.Idade, andre.Ativo)

}

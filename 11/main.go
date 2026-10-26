package main

import "fmt"

// Aula Iniciando com Structs
// Obs: Go não é uma linguagem orientada a objetos. A linguagem possui o seu "Go Way" de programar.
// Struct não é uma classe mas é o que mais se assemelha com orientação a objetos

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

type Produto struct {
	Nome       string
	Preco      float64
	Disponivel bool
}

func main() {

	andre := Cliente{
		Nome:  "André",
		Idade: 34,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", andre.Nome, andre.Idade, andre.Ativo)
	fmt.Println()

	andre.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", andre.Nome, andre.Idade, andre.Ativo)

	baixo_MB100 := Produto{
		Nome: "Baixo Elétrico Waldman MB-100",
		Preco: 969.99,
		Disponivel: true,
	}

	fmt.Println("Nome do produto: ", baixo_MB100.Nome)
	fmt.Println("Preço do produto: ", baixo_MB100.Preco)
	fmt.Println("Disponibilidade do produto: ", baixo_MB100.Disponivel)

	// alterando valores
	baixo_MB100.Preco = 1099.99
	baixo_MB100.Disponivel = false

	//fmt.Printf("Nome: $s, Preco: R$ %.2f, Disponível: %t\n", baixo_MB100.Nome, baixo_MB100.Preco, baixo_MB100.Disponivel)

	fmt.Println("\nNome do produto: ", baixo_MB100.Nome)
	fmt.Println("Preço do produto: ", baixo_MB100.Preco)
	fmt.Println("Disponibilidade do produto: ", baixo_MB100.Disponivel)

}

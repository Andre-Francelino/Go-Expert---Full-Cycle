package main

// Aula importação de pacote fmt

import (
	"fmt"
)

type ID int
type OBRA string
type OCORRENCIA string
type SERVICO string

var(
	id ID = 1
	e float64 = 7.5
	obr OBRA = "0202602731"
	descricao_servico SERVICO = "Substituição de poste"
	numero_ocorrencia OCORRENCIA = "293416/2026"
)

func main() {
	fmt.Printf("o tipo de id é %T", id)
	fmt.Println()
	fmt.Printf("o valor de id é %v", id)
	fmt.Println()

	println()
	fmt.Printf("O tipo de E é %T", e)
	fmt.Println()
	fmt.Printf("O valor de E é %v", e)
	fmt.Println()

	println()
	fmt.Printf("o tipo da variável obr é %T", obr)
	fmt.Println()
	fmt.Printf("O número da obra cadastrada é %v", obr)
	fmt.Println()

	println()
	fmt.Printf("o tipo da ocorrência é %T", numero_ocorrencia)
	fmt.Println()
	fmt.Printf("nº ocorrência cadastrada: %v", numero_ocorrencia)
	fmt.Println()
	fmt.Printf("serviço a ser realizado: %v", descricao_servico)

}
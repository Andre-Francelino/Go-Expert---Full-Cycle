package main

// Aula Maps

import "fmt"


func main() {
	salarios := map[string]int{"André": 3500, "Elder": 2300, "Luísa": 2300, "Richard": 2300, "Julia": 4200, "John": 1850}
	fmt.Println(salarios)
	fmt.Println(salarios["André"])
	fmt.Println(salarios["Julia"])

	// deletando
	delete(salarios, "John")

	// acrescentando
	salarios["Raquel"] = 2000

	fmt.Println(salarios)

	//sal := make(map[string]int)
	//sall := map[string]int{}

	// percorrendo map salarios
	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

}

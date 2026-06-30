package main

// Aula percorrando arrays

import (
	"fmt"
)

func main() {
	var meuArray [5]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	meuArray[3] = 40
	meuArray[4] = 50

	fmt.Println(len(meuArray) - 1)
	fmt.Println(len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

}

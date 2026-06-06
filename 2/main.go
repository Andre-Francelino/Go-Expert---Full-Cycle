package main

// Aula Declaração de Variáveis

// constante
const a = "Olá mundo!"

// var + nome da variavel + tipo da variavel
// escopo global
var b bool
var c bool

// escopo global
var (
	x int
	y int
	z float64
	nome string = "Goku"

	// inferência de valor
	f int
	g string
	h float32

)

func main() {
	b = true
	println(a)
	println(b)
	println(c)

	x = 11
	y = 85
	z = 7.7
	println(x)
	println(y)
	println(z)
	println(nome)
	
	// inferência de valor
	println("inferência de um valor inicial mesmo a variável estando supostamente vazia:")
	println(f)
	println(g)
	println(h)

	// escopo local (declaração dentro da função main)
	var loc string
	print(loc)

	arquivo := "X" //string
	println(arquivo)
	// arquivo := "XXX" Só pode usar os 2 pontos (:) na primeira vez que está declarando a variável
	arquivo = "XXX" // se for atribuir um novo valor de uma variável já criada já se atribui utlizando somente o igual (=)
	println(arquivo)
	
}

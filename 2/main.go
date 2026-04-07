package main

// Aula Declaração de Variáveis

// constante
const a = "Olá mundo!"

// var + nome da variavel + tipo da variavel
var b bool
var c bool

var (
	x int
	y int
	z float64
	nome string = "Goku"
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
}

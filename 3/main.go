package main

// Aula - Criação de tipos
type ID int
type OBRA string
type PRIORIDADE string
type EQ string
type STATUS string

var (
	y ID = 1
	obr OBRA = "0032700001"
	pri PRIORIDADE = "REN1000"
	equipe EQ = "TEC_LDCT 15"
	stts STATUS = "PROGRAMADA"
)

func main() {
	println(y)
	println(obr)
	println(pri)
	println(equipe)
	println(stts)
}

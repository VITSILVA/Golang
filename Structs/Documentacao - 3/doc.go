package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
	
type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "joao",
			idade: 31,
		},
		titulo:  "pizzaiolo",
		salario: 10000,
	}
	pessoa3 := pessoa{"mauricio", 40}
	pessoa4 := profissional{pessoa{"vanderlei", 50}, "politico", 1000000}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.nome)
	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa4)
}

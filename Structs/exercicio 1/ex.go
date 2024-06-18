package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{
		nome:      "Vitor",
		sobrenome: "Augusto",
		sabores:   []string{"banana", "maçã", "jaca"},
	}
	pessoa2 := pessoa{
		nome:      "Joao",
		sobrenome: "Gourlat",
		sabores:   []string{"banana", "maçã", "jaca"},
	}
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}
}

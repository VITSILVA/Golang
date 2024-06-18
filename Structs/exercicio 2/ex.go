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

	meumapa := make(map[string]pessoa)

	meumapa["Augusto"] = pessoa{
		nome:      "Vitor",
		sobrenome: "Augusto",
		sabores:   []string{"banana", "maçã", "jaca"},
	}
	meumapa["Gorulart"] = pessoa{
		nome:      "Joao",
		sobrenome: "Gourlat",
		sabores:   []string{"morango", "uva", "limao"},
	}
	for _, valor := range meumapa {
		fmt.Println("meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores {
			fmt.Println("-", valor)
		}
		fmt.Println("\n")
	}
}

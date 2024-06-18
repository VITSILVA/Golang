package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) metodo() {
	fmt.Println(p.nome, "é o nome e", p.sobrenome, "é o sobrenome")
}

func main() {
	vitor := pessoa{
		nome:      "vitao",
		sobrenome: "augusto",
		idade:     21,
	}
	vitor.metodo()
}

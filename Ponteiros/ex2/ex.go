package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func main() {
	vitao := pessoa{
		nome:  "vitor",
		idade: 21,
	}
	fmt.Println(vitao)
	mudeme(&vitao)
	fmt.Println(vitao)

}

func mudeme(p *pessoa) {
	(*p).nome = "vitinho"
	p.idade = 13
}

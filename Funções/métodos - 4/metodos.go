package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia", "e possui", p.idade, "anos de idade")
}
func main() {
	mauricio := pessoa{"mauricio", 30}
	mauricio.oibomdia()

}

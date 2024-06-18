package main

import (
	"fmt"
)

type pessoa struct {
	Nome      string
	Sobrenome string
}

type humanos interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Println(p.Nome, p.Sobrenome, "Diz oi")
}

func DizerAlgumaCoisa(h humanos) {
	h.falar()

}

func main() {
	Vitor := pessoa{
		Nome:      "Vitor",
		Sobrenome: "Augusto",
	}

	Yasmin := pessoa{
		Nome:      "Yasmin",
		Sobrenome: "Santos",
	}
	Vitor.falar() // é um shortcurt de (&Vitor).falar()
	Yasmin.falar()

	(&Vitor).falar() // forma "mais correta" de chamar
	(&Yasmin).falar()

	DizerAlgumaCoisa(&Vitor) // funciona
	// DizerAlgumaCoisa(Vitor) // não funciona
	DizerAlgumaCoisa(&Yasmin)
}

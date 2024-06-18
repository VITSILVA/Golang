package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salário          float64
}

type arquiteto struct {
	pessoa
	tipodecontrução  string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("meu nome é", x.nome, "eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("meu nome é", x.nome, "e ouve só: bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("eu ganho:", g.(dentista).salário)

	case arquiteto:
		fmt.Println("eu construo:", g.(arquiteto).tipodecontrução)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "mister dente",
			sobrenome: "silva",
			idade:     50,
		},
		dentesarrancados: 200,
		salário:          10000,
	}
	mrprédio := arquiteto{
		pessoa: pessoa{
			nome:      "mister dente",
			sobrenome: "silva",
			idade:     50,
		},
		tipodecontrução:  "prédios",
		tamanhodaloucura: "muita",
	}
	serhumano(mrdente)
	fmt.Println("")
	serhumano(mrprédio)
}

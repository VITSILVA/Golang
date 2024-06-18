package main

import (
	"fmt"
)

type veiculo struct {
	porta int
	cor   string
}
type caminhonete struct {
	veiculo
	traçãonasquatrorodas bool
}
type sedan struct {
	veiculo
	modeloluxo bool
}

func main() {
	carretagrande := caminhonete{
		veiculo: veiculo{
			porta: 4,
			cor:   "azul",
		},
		traçãonasquatrorodas: true,
	}

	carretasedan := sedan{veiculo{2, "vermelha"}, true}

	fmt.Println(carretagrande)
	fmt.Println(carretasedan)
	fmt.Println(carretagrande.cor)
	fmt.Println(carretasedan.modeloluxo)

}

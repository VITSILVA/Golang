package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {

	jamesbond := pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000000,
	}
	darthvader := pessoa{"Anakin", "Skywalker", 50, "Vilão de Filme", 5000000}

	j, err := json.Marshal(jamesbond)
	if err != nil {
		fmt.Println(err)
	}
	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))

}

package main

import (
	"fmt"
)

func main() {

	esportefavorito := "futebol"
	switch esportefavorito {

	case "futebol":
		fmt.Println("futebol é o esporte favorito")

	case "volei":
		fmt.Println("volei é o esporte favorito")

	case "basquete":
		fmt.Println("basquete é o esporte favorito")
	}
}

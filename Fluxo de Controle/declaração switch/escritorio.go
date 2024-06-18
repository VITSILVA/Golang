package main

import (
	"fmt"
)

func main() {

	quemestanoescritoriohoje := "vitor"

	switch quemestanoescritoriohoje {

	case "zezinho":
		fmt.Println("quem esta no escritorio é o zezinho")

	case "marquinhos":
		fmt.Println("quem esta no escritorio é o marquinhos")

	case "joana":
		fmt.Println("quem esta no escritorio é o joana")

	case "maria":
		fmt.Println("quem esta no escritorio é o maria")

	default:
		fmt.Println("ta vazio ou é feriado")
	}
}

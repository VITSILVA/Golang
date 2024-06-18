package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"alfredo": 4444444,
		"joana":   2222222,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 123456
	fmt.Println(amigos)
	fmt.Println(amigos["gopher"], "\n")

	sera, ok := amigos["esmeralda"]
	fmt.Println(sera, ok)

	if sera, ok := amigos["esmeralda"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(sera)
	}
}

package main

import (
	"fmt"
)

func maina() {
	slice := []string{"banana", "maçã", "jaca"}

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}
	//slice[3] = "melancia"
	//for índice, valor := range slice {
	//	fmt.Println("No índice", índice, "temos o valor:", valor)  <- fora do alcance do array = erro

	//}
	slice = append(slice, "melancia")

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}
}

package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for indice, valor := range slice {
		fmt.Println("indice", indice, "vale", valor)
	}
	fmt.Printf("%T", slice)
}

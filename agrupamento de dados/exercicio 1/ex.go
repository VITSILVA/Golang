package main

import (
	"fmt"
)

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	for indice, valor := range array {
		fmt.Println("array", indice, "vale", valor)
		//fmt.Printf("\n%T", array)
	}
	fmt.Printf("%T", array)
}

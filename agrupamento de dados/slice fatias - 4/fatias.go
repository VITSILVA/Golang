package main

import (
	"fmt"
)

func main() {
	sabores := []string{"portuguesa", "pepperoni", "abacaxi", "queijo", "marg"}

	fatia := sabores[0:3]
	fmt.Println(fatia)

	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}

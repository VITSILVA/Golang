package main

import (
	"fmt"
)

func main() {
	x := 0
	for {
		if x < 10 {
			x++
			fmt.Println("chis é menor que déis")
		} else {
			fmt.Println("chis é maior que déis oxi")
			break
		}
	}
	fmt.Println("o loop está brecado!")
}

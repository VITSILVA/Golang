package main

import (
	"fmt"
)

func maina() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("chis é menor que cinco")

	case x == 5:
		fmt.Println("chis é igual a cinco")

	case x > 5:
		fmt.Println("chis é maior que cinco")
	}

}

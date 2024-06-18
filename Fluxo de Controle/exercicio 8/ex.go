package main

import (
	"fmt"
)

func main() {

	x := 10
	switch {
	case x < 10:
		fmt.Println("xis é menor que deiz")

	case x > 10:
		fmt.Println("xis é maior que deiz")

	case x == 10:
		fmt.Println("xis é igual a deiz")
	}
}

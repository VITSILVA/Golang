package main

import (
	"fmt"
)

func main() {

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("\n\nmes:", mes)
		for x := 1; x <= 30; x++ {
			fmt.Print("dia: ", x, ",  ")
		}

	}

}

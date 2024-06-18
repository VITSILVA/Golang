package main

import (
	"fmt"
)

func maina() {

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("\n\nhora:", horas)
		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}

	}

}

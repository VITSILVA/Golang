package main

import (
	"fmt"
)

func main() {
	x := 2001
	y := 2024

	for {

		if x > y {
			break
		}
		fmt.Println(x)
		x++

	}
}

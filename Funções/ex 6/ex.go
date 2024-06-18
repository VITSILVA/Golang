package main

import (
	"fmt"
)

func main() {
	x := 100
	func(x int) {
		fmt.Println(x, "vezes 3 é:", x*3)
		fmt.Println(x * 3)
	}(x)
}

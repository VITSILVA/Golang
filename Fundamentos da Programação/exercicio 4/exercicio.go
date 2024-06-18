package main

import (
	"fmt"
)

func main() {

	x := 2
	fmt.Printf("%d - %b - %#X", x, x, x)

	fmt.Println("")

	y := x << 1
	fmt.Printf("%d - %b - %#X", y, y, y)

}

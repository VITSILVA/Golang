package main

import (
	"fmt"
)

func main() {
	x := 1
	y := x << 2

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
}

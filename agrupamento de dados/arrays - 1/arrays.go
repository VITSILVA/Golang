package main

import (
	"fmt"
)

var x [5]int

func main() {
	x[0] = 1
	x[1] = 2

	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T", x)
	fmt.Println(len(x))
}

package main

import (
	"fmt"
)

func main() {
	x := 10
	y := func(x int) int {
		//	fmt.Println(x, "vezes 444 é:")
		return x * 12
	}
	fmt.Println(x, "vezes 12 é:", y(x))
}

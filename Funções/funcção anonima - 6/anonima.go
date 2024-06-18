package main

import (
	"fmt"
)

func main() {
	x := 333
	func(x int) {
		fmt.Println(x, "vezes 444 é:")
		fmt.Println(x * 444)
	}(x)
}

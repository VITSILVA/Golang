package main

import (
	"fmt"
)

func main() {
	s := "ascii éã§"

	for _, v := range s {
		fmt.Printf("%b - %v - %T - %#U - %#X\n", v, v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#X\n", s[i], s[i], s[i], s[i])
	}

}

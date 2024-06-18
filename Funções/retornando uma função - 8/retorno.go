package main

import (
	"fmt"
)

func main() {
	a := retornaumafuncao()
	y := a(3)
	fmt.Println(y)
	fmt.Println(retornaumafuncao()(3))
}
func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}

package main

import (
	"fmt"
)

func main() {
	si := []int{1, 2, 3, 4, 5}
	si2 := []int{1, 10, 43, 9}

	fmt.Println(funcao1(si...))
	fmt.Println(funcao2(si2))
}
func funcao1(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
func funcao2(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

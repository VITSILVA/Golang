package main

import (
	"fmt"
)

func main() {
	si := []int{10, 10, 11, 4, 3}
	total := soma(si...) //<- podemos passar zero valores
	fmt.Println(total)
}
func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

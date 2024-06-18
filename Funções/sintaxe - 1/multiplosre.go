package main

import (
	"fmt"
)

func mains() {
	total, quantos, oi := soma(10, 10, 11, 4, 3, 3)
	fmt.Println(total, quantos, oi)
}
func soma(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "bom dia"
}

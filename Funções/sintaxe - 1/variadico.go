package main

import (
	"fmt"
)

func mainp() {
	total, quantos, oi := soma("tarde", 10, 10, 11, 4, 3, 3)
	fmt.Println(total, quantos, oi)
}
func soma(s string, x ...int) (int, int, string) {
	oi := ""

	if s == "manhã" {
		oi = "oi, bom dia"
	} else if s == "tarde" {
		oi = "oi, boa tarde"
	} else {
		oi = "oi, boa noite"
	}

	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), oi

}

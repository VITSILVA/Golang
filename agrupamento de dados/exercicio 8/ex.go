package main

import (
	"fmt"
)

func main() {
	pessoas := map[string][]string{
		"vitor":  {"fazer academia", "levantar peso", "chutar"},
		"yasmin": {"dormir um pouco", "treinar", "gostar do campo"},
	}
	for k, v := range pessoas {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "-", h)
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	if x := 110; x > 100 {
		fmt.Println("chis é maior quer cem")

	} else if x < 10 {
		fmt.Println("chis é menor quer deis")

	} else {
		fmt.Println("chis não é maior que deis nem maior que cem")
	}
}

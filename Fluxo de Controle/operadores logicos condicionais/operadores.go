package main

import (
	"fmt"
)

func main() {
	x := 6

	if (x == 2) || (x == 6) {
		fmt.Println("chis é dois ou seis")
	}
	if x%2 == 0 && x%3 == 0 {
		fmt.Println("chis é multiplo de dois e treis")
	}
	if x%2 == 0 || x%3 == 0 {
		fmt.Println("chis é multiplo de dois e treis")
	}
}

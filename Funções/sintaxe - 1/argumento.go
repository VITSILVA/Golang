package main

import (
	"fmt"
)

func main() {
	argumento("manhã")
	argumento("tarde")
	argumento("jacare")
}

 func basica() {
 fmt.Println("oi, bom dia")
 
func argumento(s string) {
	if s == "manhã" {
		fmt.Println("oi, bom dia")
	} else if s == "tarde" {
		fmt.Println("oi, boa tarde")
	} else {
		fmt.Println("oi, boa noite")
	}
}
}

// func (reciver) nome (paramentro) (returns) {code} <- função tem esse formato

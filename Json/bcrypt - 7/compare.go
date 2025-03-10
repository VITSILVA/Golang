package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashPassword(sb, []byte(senhaerrada)))
}

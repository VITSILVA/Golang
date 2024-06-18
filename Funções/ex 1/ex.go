package main

import (
	"fmt"
)

func main() {
	fmt.Println(retornaumint())
	fmt.Println(retornaumintestring())
}

func retornaumint() int {
	return 10

}

func retornaumintestring() (int, string) {
	return 10, "deiz"

}

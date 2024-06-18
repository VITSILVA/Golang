package main

import "fmt"

type teclado int

var x teclado

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}

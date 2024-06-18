package main

import "fmt"

func main() {

	x := 10
	y := "ola bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20
	fmt.Printf("x: %v, %T", x, x)

}

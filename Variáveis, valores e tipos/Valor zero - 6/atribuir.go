package main

import"fmt"

var x int                <- declaração ( )

func main(){

	x = 10               <- inicialização e atribuição ( primeira atrunição de valor a uma variavel)
	fmt.Printf("%v, %T", x, x)

	x = 20              <- atribuição de valor a uma variavel
	fmt.Printf("%v, %T", x, x)


}
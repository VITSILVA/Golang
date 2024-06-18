package main

import"fmt"

var x int = 10 	

func main(){

	x = 20.5                 <- isso é errado pois "var x int = 10" refere-se a um numero inteiro e não a uma fração
	fmt.Println(x)
}

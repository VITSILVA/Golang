package main

import"fmt"

var x int                 <- variável declarada sem valor, essa atribuição só podera ser feita dentro de um code block

func main(){

	x = 10               <- atribuição feita dentro de um code block
	fmt.Printf("%v, %T, x, x")


}
package main

import (
	"fmt"
)

func maina() {

	s := "hello, playground"
	fmt.Printf("%v\n%T\n\n", s, s)

	t := `hello,    voce 
	                playground`

	fmt.Printf("%v\n%T", t, t)

}

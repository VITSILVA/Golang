package main

import "fmt"

type hotdo int

var c hotdo

func maina() {
	fmt.Printf("%T", c)
}

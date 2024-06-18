package main

import (
	"fmt"
)

const (
	a = iota * 10
	_
	c
	x
	y
	z
)

func main() {
	fmt.Println(a, c, x, y, z)

}

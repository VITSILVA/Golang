package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c

	for i := 0; i < 100; i++ {
		select {
		case v := <-q:
			fmt.Println("Canal A:", v)
		case v := <-c:
			fmt.Println("Canal B:", v)
		}
	}
}

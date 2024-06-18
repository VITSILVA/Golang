package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // buffer colocariamos make(chan int, 1) e tiratiamos a go func

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

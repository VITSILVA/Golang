package main

import (
	"fmt"
)

func main() {
	canal := make(chan int, 1)

	canal <- 40
	fmt.Println(<-canal)
}

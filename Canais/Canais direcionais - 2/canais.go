package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go send(canal)

	receiver(canal)

}

func send(s chan<- int) {
	s <- 42

}

func receiver(r <-chan int) {
	fmt.Println("o valor recebido é:", <-r)

}

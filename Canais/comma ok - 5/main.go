package main

import (
	"fmt"
)

func main() {
	par := make(chan int)
	ímpar := make(chan int)
	quit := make(chan bool)

	go mandanumeros(par, ímpar, quit)

	receive(par, ímpar, quit)
}

func mandanumeros(par, ímpar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- 1
		} else {
			ímpar <- i
		}
	}
	close(par)
	close(ímpar)
	quit <- true
}
func receive(par, ímpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-ímpar:
			fmt.Println("O número", v, "é ímpar.")
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu ruim! ó:", v)
			} else {
				fmt.Println("Encerrando, recebemos:", v)
			}
			return
		}
	}
}

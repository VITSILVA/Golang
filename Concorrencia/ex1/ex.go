package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func1()
	go func2()

	wg.Wait()

}

func func1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Func1:\t", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 5; i++ {
		fmt.Println("Func2:\t", i)
	}
	wg.Done()

}

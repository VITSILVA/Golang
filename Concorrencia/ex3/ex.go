package main

import (
	"fmt"
	"runtime"
	"sync"
)

func maina() {

	contador := 0
	totaldegoroutines := 100

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < 100; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines no meio:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines no final:", runtime.NumGoroutine())
	fmt.Println("Esse é o valor final", contador)
}

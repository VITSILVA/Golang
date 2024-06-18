package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])

	slice[3] = 333333
	fmt.Println(slice[3])

	// slice[20] = 10          <- da erro pois esta fora de alcance do array
	//fmt.Println(slice[20])

}

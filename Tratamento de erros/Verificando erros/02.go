package main

import (
	"fmt"
)

func main() {
	var answer1, answer2, answer3 string

	fmt.Print("name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("fav food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("fav sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer1, answer2, answer3)
}

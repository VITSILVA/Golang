package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"Vitor", "Correu", "Atrás", "Do", "Pipa"}

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)
}

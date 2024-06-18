package main

import (
	"fmt"
)

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		988: "massa",
		455: "bacana",
		478: "legal",
	}
	fmt.Println(qualquercoisa)
	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}
}

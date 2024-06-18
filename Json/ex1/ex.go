package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	ua, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	ud, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
	}

	ub, err := json.Marshal(u2)
	if err != nil {
		fmt.Println(err)
	}

	uc, err := json.Marshal(u3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(ua))
	fmt.Println(string(ud))
	fmt.Println(string(ub))
	fmt.Println(string(uc))

}

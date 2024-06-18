package main

import (
	"fmt"
)

func main() {
	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "waffel",
		sabor:   "doce",
		ondetem: []string{"no brasil", "na mercearia"},
		vaibemcom: map[string]string{
			"no café da manhã": "chá",
			"no almoço":        "cafezinho",
			"no jantar":        "não vai bem",
		},
	}
	fmt.Println(x)

}

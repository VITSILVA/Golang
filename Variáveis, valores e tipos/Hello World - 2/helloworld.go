package main

import "fmt"

func main() {
	fmt.Println("Helllo World")

}

func mainB() {
	numerodebytes, erros := fmt.Println("Helllo World")
	fmt.Println(numerodebytes, erros)

}

func mainC() {
	_, erros := fmt.Println("Helllo World")
	fmt.Println(erros)

}

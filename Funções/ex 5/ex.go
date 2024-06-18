package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("a área do quadrado é:", resultado)
}

func (c circulo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println("a área do circulo é:", resultado)
}

type figura interface {
	area()
}

func info(f figura) {
	f.area()
}
func main() {

	x := quadrado{
		lado: 10.0,
	}
	y := circulo{
		raio: 10.5,
	}

	info(x)
	info(y)

}

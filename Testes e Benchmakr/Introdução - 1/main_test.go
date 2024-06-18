package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	teste := Soma(1, 2, 3)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}
func TestMultiplica(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

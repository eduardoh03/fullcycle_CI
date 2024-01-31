package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtract(t *testing.T) {

	total := subtract(15, 15)

	if total != 0 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 0)
	}
}
func TestMultiply(t *testing.T) {

	total := multiply(15, 15)

	if total != 225 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 225)
	}
}
func TestDivide(t *testing.T) {

	total := divide(15, 15)

	if total != 1 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}

package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}

	// Test with different numbers
	total = soma(10, 20)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(15, 15)
	if total != 0 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 0)
	}

	// Test with different numbers
	total = subtract(20, 10)
	if total != 10 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 10)
	}
}

func TestMultiply(t *testing.T) {
	total := multiply(15, 15)
	if total != 225 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado %d. Esperado: %d", total, 225)
	}

	// Test with different numbers
	total = multiply(5, 5)
	if total != 25 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado %d. Esperado: %d", total, 25)
	}
}

func TestDivide(t *testing.T) {
	total := divide(15, 15)
	if total != 1 {
		t.Errorf("Resultado da divisão é inválido: Resultado %d. Esperado: %d", total, 1)
	}

	// Test with different numbers
	total = divide(20, 5)
	if total != 4 {
		t.Errorf("Resultado da divisão é inválido: Resultado %d. Esperado: %d", total, 4)
	}

	// Test division by 0
	total = divide(10, 1)
	if total != 10 {
		t.Errorf("Resultado da divisão por 0 é inválido: Resultado %d. Esperado: %d", total, 10)
	}
}

func TestRunProgram(t *testing.T) {
	runProgram(112, 10)

}

func TestMainFunction(t *testing.T) {
	runProgram(112, 10)
}

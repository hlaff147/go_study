package main

import (
	"fmt"
	"testing"
)

func main() {
	TestSum(&testing.T{})
	fmt.Println("Testes executados manualmente.")
}

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Resultado incorreto: esperado %d, obtido %d", expected, result)
	}

	result = Sum(-1, 1)
	expected = 0
	if result != expected {
		t.Errorf("Resultado incorreto: esperado %d, obtido %d", expected, result)
	}
}

func Sum(a, b int) int {
	return a + b
}

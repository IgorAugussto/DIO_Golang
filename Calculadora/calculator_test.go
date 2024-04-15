package main

import (
	"testing"
)

func Test_ShouldSumCorrect(t *testing.T) {
	teste := sum(3, 2)

	resultado := 5

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func Test_ShouldSumIncorrect(t *testing.T) { // convenção de nomes para teste negativos (shouldSumIncorrect), sempre deixar para teste errados
	teste := sum(3, 2)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func Test_ShouldMultCorrect(t *testing.T) {
	teste := multi(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func Test_ShouldMultIncorrect(t *testing.T) {
	teste := multi(10, 10)
	resultado := 2500

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func Test_ShouldSubCorrect(t *testing.T) {
	teste := subtract(3, 2)

	resultado := 5

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func Test_ShouldSubIncorrect(t *testing.T) {
	teste := subtract(3, 2)

	resultado := 5

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func Test_ShoudDivisionCorrect(t *testing.T) {
	test := division(10, 2)

	result := 5

	if test != result {
		t.Error("Valor esperado: ", result, "Valor retornado: ", test)
	}
}

func Test_ShoudDivisionIncorrect(t *testing.T) {
	test := division(10, 2)

	result := 2

	if test != result {
		t.Error("Valor esperado: ", result, "Valor retornado: ", test)
	}
}

package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var op string

	// Solicita ao usuário os números e a operação
	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Escolha a operação (+, -, *, /):")
	fmt.Scanln(&op)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	// Realiza a operação
	var result int
	switch op {
	case "+":
		result := sum(num1, num2)
	case "-":
		result := subtract(num1, num2)
	case "*":
		result := multi(num1, num2)
	case "/":
		result := float64(division(float64(num1), float64(num2)))
	}

	// Exibe o resultado
	fmt.Printf("Resultado: %f\n", result)
}

func sum(num1, num2 int) int {
	result := num1 + num2
	return result
}

func subtract(num1, num2 int) int {
	result := num1 - num2
	return result
}

func multi(num1, num2 int) int {
	result := num1 * num2
	return result
}

func division(num1, num2 float64) float64 {
	var result float64

	if num1 == 0 {
		fmt.Println("Operação inválida!\nDivisão por zero")
	} else if num2 == 0 {
		fmt.Println("Operação inválida!\nDivisão por zero")
	} else {
		result = num1 / num2
	}
	return result
}

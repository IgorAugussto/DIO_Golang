package main

import (
	"bufio"
	"fmt"
	"os"
)

const K = 373.0

func ChooseConversion() {

	var tempC float64
	var tempK float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite 'C' para converter de Celsius para Kelvin e 'K' para converter de Kelvin para Celsius.")
	scanner.Scan()
	option := scanner.Text()

	if option == "C" {
		// Chama a função para converter de Celsius para Kelvin
		fmt.Printf("digite o valor a ser convertido: ")
		fmt.Scan(&tempC)
		celsiusToKelvin(tempC)
		fmt.Println("Programa encerrado. Pressione Enter para sair.")
		fmt.Scanln()
	} else if option == "K" {
		// Chama a função para converter de Kelvin para Celsius
		fmt.Printf("digite o valor a ser convertido: ")
		fmt.Scan(&tempK)
		kelvinToCelsius(tempK)
		fmt.Println("Programa encerrado. Pressione Enter para sair.")
		fmt.Scanln()
	} else {
		fmt.Println("Opção inválida. Digite 'C' ou 'K'.")
	}
}

func celsiusToKelvin(tempC float64) float64 {
	tempK := tempC + 273.15
	fmt.Scanln("%.2f Celsius são %.2f Kelvin", tempC, tempK)
	return tempK
}

func kelvinToCelsius(tempK float64) float64 {
	tempC := tempK - 273.15
	fmt.Scanln("%.2f Kelvin são %.2f Celsius", tempK, tempC)
	return tempC

}

func main() {
	tempK := K
	tempC := tempK - 273.15

	fmt.Printf("A temperatura de ebulição da água em Kelvin é de %.2f, a temperatuda de ebulição da água em Celsius é %.2f \n",
		tempK, tempC)

	scanner := bufio.NewScanner(os.Stdin)

	// Pergunta ao usuário se deseja converter outro valor
	for {
		fmt.Println("Digite 'S' caso queira converter outro valor de Kelvin para Celsius e 'N' caso queira encerrar o programa.")
		scanner.Scan()
		resposta := scanner.Text()

		// Verifica a resposta do usuário
		if resposta == "S" {
			// Chama a função para escolher a conversão
			ChooseConversion()
			break
		} else if resposta == "N" {
			fmt.Println("Programa encerrado.")
			break
		} else {
			fmt.Println("Resposta inválida. Digite 'S' ou 'N'.")
		}
	}
	fmt.Scanln()
}

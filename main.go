package main

import "fmt"

func menu() {
	fmt.Println("Selecione a operacão desejada")

	fmt.Println("(+) Soma")
	fmt.Println("(-) Subtração")
	fmt.Println("(*) Multiplicação")
	fmt.Println("(/) Divisão")
}

func main() {
	var num1, num2 float64
	var operador string
	var resultado float64

	fmt.Println("Entrada para valor 1:")
	fmt.Scanln(&num1)

	fmt.Println("Entrada para valor 2:")
	fmt.Scanln(&num2)

	menu()
	fmt.Scanln(&operador)

	switch operador {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 != 0 {
			resultado = num1 / num2

		}
	default:
		fmt.Println("Operação inválida!")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operador, num2, resultado)
}

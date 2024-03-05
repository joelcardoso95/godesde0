package funcoes

import "fmt"

func Calculos() {

	var numero3 int = 32
	var numero4 int = 345

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 50))

	calculo = func(numero1, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	fmt.Println(calculo(10, 50))
}

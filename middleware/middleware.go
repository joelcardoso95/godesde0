package middleware

import "fmt"

func somar(a, b int) int {
	return a + b
}

func subtrair(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func Middleware() {
	fmt.Println("Inicio")

	resultado := operacoesMidd(somar)(2, 3)
	fmt.Println(resultado)
	resultado = operacoesMidd(subtrair)(6, 3)
	fmt.Println(resultado)
	resultado = operacoesMidd(multiplicar)(5, 3)
	fmt.Println(resultado)
}

func operacoesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operação")
		return f(a, b)
	}
}

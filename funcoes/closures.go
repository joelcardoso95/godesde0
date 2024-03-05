package funcoes

import "fmt"

func tabuada(valor int) func() int {
	numero := valor
	sequencia := 0
	return func() int {
		sequencia++
		return numero * sequencia
	}
}

func ChamarClosure() {
	tabuada1 := 2
	MinhaTabuada := tabuada(tabuada1)

	for i := 1; i <= 10; i++ {
		fmt.Println(MinhaTabuada())
	}
}

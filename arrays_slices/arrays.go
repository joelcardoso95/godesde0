package arraysslices

import "fmt"

var tabela [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MostrarArrays() {
	tabela[7] = 33
	tabela[2] = 15

	tabela2 := [10]string{"Pablo", "Chaves"}

	fmt.Println(tabela)
	fmt.Println(tabela2)

	for i := 0; i < len(tabela); i++ {
		fmt.Println(tabela[i])
	}

	matriz[9][25] = 15
	fmt.Println(matriz)
}

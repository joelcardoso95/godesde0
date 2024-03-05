package arraysslices

import "fmt"

var tabelaS []int = []int{2, 5, 4}
var array [10]int = [10]int{6, 9, 3, 2, 134, 555, 999, 222, 444}

func MostrarSlices() {
	fmt.Println(tabelaS)
	porcao := array[3:]  // slice criado de um vetor, desde posição 3
	porcao2 := array[:5] // slice criado de um vector, desde a posição 0 até 5
	porcao3 := array[6:7]

	fmt.Println(porcao)
	fmt.Println(porcao2)
	fmt.Println(porcao3)
}

func Capacidade() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largura %d, Capacidade %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargura %d, Capacidade %d", len(nums), cap(nums))
}

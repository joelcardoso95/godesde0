package funcoes

import "fmt"

func Exponenciao(valor int) {
	if valor > 1000 {
		return
	}

	fmt.Println(valor)
	Exponenciao(valor * 2)
}

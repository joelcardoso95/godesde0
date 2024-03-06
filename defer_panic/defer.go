package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Esta é uma primeira mensagem")
	defer fmt.Println("Esta é uma mensagem final")

	fmt.Println("É uma segunda mensagem")
}

func ExemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocorreu um erro que gerou PANIC \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Encontrou valor 1")
	}
}

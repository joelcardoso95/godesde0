package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MeuNomeLento(nome string, canal1 chan bool) {
	letras := strings.Split(nome, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	canal1 <- true
}

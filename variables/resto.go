package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nome string
var Estado bool
var Saldo float32
var Data time.Time

func OutrasVariaveis() {
	Nome = "Pedro"
	Estado = true
	Saldo = 1526.88
	Data = time.Now()
	fmt.Println(Nome)
	fmt.Println(Estado)
	fmt.Println(Saldo)
	fmt.Println(Data)
}

func ConverterToTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}

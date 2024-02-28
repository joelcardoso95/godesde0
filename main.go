package main

import (
	"fmt"
	"projects/godesde0/variables"
)

func main() {
	estado, texto := variables.ConverterToTexto(500)
	fmt.Println(estado)
	fmt.Println(texto)
}

package exer_interfaces

import (
	"fmt"
	"projects/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Sou um/uma %s, e estou respirando \n", hu.Sexo())
}

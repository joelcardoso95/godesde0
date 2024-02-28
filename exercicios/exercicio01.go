package exercicios

import "strconv"

func ConvNumerico(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "ERROR =" + err.Error()
	}

	if num > 100 {
		return num, "É maior que 100"
	} else {
		return num, "Menor que 100"
	}
}

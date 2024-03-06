package main

import (
	"projects/godesde0/webserver"
)

func main() {
	/*estado, texto := variables.ConverterToTexto(500)
	fmt.Println(estado)
	fmt.Println(texto) */

	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Não é Windows")
	} else {
		fmt.Println("É Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("LINUX")
	case "darwin":
		fmt.Println("DARWIN")
	default:
		fmt.Printf("%s \n", os)
	}*/

	/*numero, texto := exercicios.ConvNumerico("AAAAAA")
	fmt.Println(numero)
	fmt.Println(texto)*/

	//teclado.EntradaNumeros()

	//iteracao.Iterar()

	//fmt.Println(exercicios.Tabuada())

	//files.SalvarTabuada()
	//files.GravarTabuadas()
	//files.LerArquivo()

	//arraysslices.Capacidade()
	//	mapas.MostrarMapas()
	//users.CriaUsuario()

	/*Pedro := new(model.Homem)
	exer_interface.HumanosRespirando(Pedro)

	Maria := new(model.Mulher)
	exer_interface.HumanosRespirando(Maria)*/

	/*canal1 := make(chan bool)
	go goroutines.MeuNomeLento("Joel Cardoso", canal1)
	fmt.Println("Estou aqui")

	<-canal1*/
	webserver.MeuWebServer()
}

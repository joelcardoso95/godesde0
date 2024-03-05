package files

import (
	"bufio"
	"fmt"
	"os"
	"projects/godesde0/exercicios"
)

var filename = "./files/txt/tabuada.txt"

func SalvarTabuada() {
	var texto string = exercicios.Tabuada()
	arquivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Erro ao criar arquivo" + err.Error())
		return
	}

	fmt.Fprintln(arquivo, texto)
	arquivo.Close()

}

func GravarTabuadas() {
	var texto string = exercicios.Tabuada()
	if !Append(filename, texto) {
		fmt.Println("erro ao concatenar arquivo")
	}
}

func Append(filen string, texto string) bool {
	arq, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erro ao concatenar arquivo " + err.Error())
	}

	_, err = arq.WriteString(texto)
	if err != nil {
		fmt.Println("Erro ao concatenar arquivo")
		return false
	}

	arq.Close()
	return true
}

func LerArquivo() {
	arquivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erro ao ler arquivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	arquivo.Close()
}

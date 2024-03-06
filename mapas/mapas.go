package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	/*fmt.Println(paises)
	fmt.Println(paises["Argentina"])*/

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	/*for equipe, pontuacao := range campeonato {
		fmt.Printf("Equipe %s, tem uma pontuação de %d \n", equipe, pontuacao)
	}*/

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	pontuacao, existe := campeonato["Chivas"]
	fmt.Printf("Pontuação capturada é %d, e equipe existe = %t \n", pontuacao, existe)
}

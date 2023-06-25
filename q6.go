package main

import "fmt"

func somarContagens(mapas []map[string]int) map[string]int {
	soma := make(map[string]int)
	for _, mapa := range mapas {
		for palavra, contagem := range mapa {
			soma[palavra] += contagem
		}
	}
	return soma
}

func main() {
	texto1 := map[string]int{
		"Lorem": 2,
		"ipsum": 3,
		"dolor": 1,
	}
	texto2 := map[string]int{
		"Lorem": 1,
		"sit":   2,
		"amet":  2,
	}
	texto3 := map[string]int{
		"dolor": 2,
		"sit":   1,
		"amet":  3,
	}

	listaMapas := []map[string]int{texto1, texto2, texto3}
	resultado := somarContagens(listaMapas)

	fmt.Println("Contagem total das palavras:")
	for palavra, contagem := range resultado {
		fmt.Printf("%s: %d\n", palavra, contagem)
	}
}

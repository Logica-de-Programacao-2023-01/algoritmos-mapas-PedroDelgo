package main

import "fmt"

func MergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	// Copiar todos os elementos do primeiro mapa para o resultado
	for key, value := range map1 {
		result[key] = value
	}

	// Adicionar ou atualizar os elementos do segundo mapa no resultado
	for key, value := range map2 {
		result[key] = value
	}

	return result
}

func main() {
	map1 := map[string]int{
		"chave1": 1,
		"chave2": 2,
	}

	map2 := map[string]int{
		"chave2": 3,
		"chave3": 4,
	}

	mergedMap := MergeMaps(map1, map2)

	// Exibir o mapa resultante
	for key, value := range mergedMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
q
package main

import (
	"fmt"
	"sort"
	"strings"
)

func ordenarPalavra(palavra string) string {
	runes := []rune(palavra)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func agruparAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)
	for _, palavra := range palavras {
		ordem := ordenarPalavra(strings.ToLower(palavra))
		anagramas[ordem] = append(anagramas[ordem], palavra)
	}
	return anagramas
}

func main() {
	palavras := []string{"amor", "roma", "carro", "mar", "ram", "corra", "amigo"}
	resultado := agruparAnagramas(palavras)
	fmt.Println("Grupos de palavras anagramas:")
	for _, grupo := range resultado {
		fmt.Println(grupo)
	}
}

package main

import (
	"fmt"
	"strings"
)

func ContarPalavras(str string) map[string]int {
	// Remover pontuação e converter para minúsculas
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, "!", "")
	str = strings.ReplaceAll(str, "?", "")
	str = strings.ReplaceAll(str, ":", "")

	// Dividir a string em palavras
	palavras := strings.Fields(str)

	// Contar a ocorrência de cada palavra
	contagem := make(map[string]int)
	for _, palavra := range palavras {
		contagem[palavra]++
	}

	return contagem
}

func main() {
	texto := "Olá mundo, olá Go! Olá Go, olá mundo!"

	resultado := ContarPalavras(texto)

	// Exibir o mapa com as palavras e suas ocorrências
	for palavra, ocorrencias := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencias)
	}
}

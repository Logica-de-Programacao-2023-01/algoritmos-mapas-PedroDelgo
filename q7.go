package main

import (
	"fmt"
	"strings"
)

func contarLetras(palavra string) map[rune]int {
	frequencia := make(map[rune]int)
	for _, char := range palavra {
		frequencia[char]++
	}
	return frequencia
}

func contarLetrasPorPalavra(frase string) map[string]map[rune]int {
	palavras := strings.Fields(frase)
	resultado := make(map[string]map[rune]int)
	for _, palavra := range palavras {
		resultado[palavra] = contarLetras(palavra)
	}
	return resultado
}

func main() {
	frase := "A frase de exemplo"
	resultado := contarLetrasPorPalavra(frase)
	fmt.Println("Contagem de letras por palavra:")
	for palavra, contagem := range resultado {
		fmt.Println(palavra, contagem)
	}
}

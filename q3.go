package main

import "fmt"

func somarValoresMapa(m map[string]int) int {
	soma := 0
	for _, valor := range m {
		soma += valor
	}
	return soma
}

func main() {
	mapa := map[string]int{
		"valor1": 10,
		"valor2": 20,
		"valor3": 30,
	}
	resultado := somarValoresMapa(mapa)
	fmt.Println("A soma dos valores é:", resultado)
}

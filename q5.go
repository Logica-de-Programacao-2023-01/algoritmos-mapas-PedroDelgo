package main

import "fmt"

func contarCaracteres(str string) map[rune]int {
	frequencia := make(map[rune]int)
	for _, char := range str {
		frequencia[char]++
	}
	return frequencia
}

func main() {
	str := "abracadabra"
	resultado := contarCaracteres(str)
	fmt.Println("Frequência dos caracteres:")
	for char, freq := range resultado {
		fmt.Printf("%c: %d\n", char, freq)
	}
}

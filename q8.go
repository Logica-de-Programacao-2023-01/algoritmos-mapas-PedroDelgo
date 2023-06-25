package main

import (
	"fmt"
)

func calcularDespesas(contas map[string]float64) map[string]float64 {
	total := 0.0
	for _, valor := range contas {
		total += valor
	}

	media := total / float64(len(contas))
	resultado := make(map[string]float64)
	for nome, valor := range contas {
		resultado[nome] = valor - media
	}

	return resultado
}

func main() {
	despesas := map[string]float64{
		"Alice": 50.0,
		"Bob":   30.0,
		"Carol": 70.0,
		"David": 40.0,
		"Eve":   60.0,
		"Frank": 20.0,
	}

	resultado := calcularDespesas(despesas)

	fmt.Println("Valor a receber/pagar para igualar as despesas:")
	for nome, valor := range resultado {
		fmt.Printf("%s: %.2f\n", nome, valor)
	}
}

package main

import "fmt"

func CalcularFrete(valor float64, peso float64) float64 {
	const taxaFixa = 10.0
	const taxaPorKg = 2.0
	return taxaFixa + taxaPorKg*peso
}

func main() {
	valorPedido := 150.0
	pesoPedido := 4.5
	frete := CalcularFrete(valorPedido, pesoPedido)
	fmt.Printf("Frete calculado para peso %.2f kg: R$%.2f\n", pesoPedido, frete)
}

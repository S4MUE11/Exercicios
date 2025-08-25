package main

import "fmt"

func MaisCaro(precos map[string]float64) (string, float64) {
	var produtoMaisCaro string
	var maiorPreco float64
	primeiro := true
	for produto, preco := range precos {
		if primeiro || preco > maiorPreco {
			maiorPreco = preco
			produtoMaisCaro = produto
			primeiro = false
		}
	}
	return produtoMaisCaro, maiorPreco
}

func main() {
	precos := map[string]float64{"Arroz": 5.99, "Carne": 28.50, "Leite": 4.30, "Queijo": 12.75}
	produto, preco := MaisCaro(precos)
	if produto == "" {
		fmt.Println("Mapa de preços vazio.")
	} else {
		fmt.Printf("Produto mais caro: %s (R$%.2f)\n", produto, preco)
	}
}

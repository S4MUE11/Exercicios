package main

import "fmt"

func AtualizarEstoque(estoque map[string]int, vendas map[string]int) map[string]int {
	for produto, quantidadeVendida := range vendas {
		if quantidadeAtual, ok := estoque[produto]; ok {
			novoTotal := quantidadeAtual - quantidadeVendida
			if novoTotal < 0 {
				novoTotal = 0
			}
			estoque[produto] = novoTotal
		}
	}
	return estoque
}

func main() {
	estoque := map[string]int{"Camiseta": 20, "Calça": 15, "Tênis": 8}
	vendas := map[string]int{"Camiseta": 5, "Tênis": 2, "Boné": 1}
	estoqueAtualizado := AtualizarEstoque(estoque, vendas)
	fmt.Println("Estoque atualizado:", estoqueAtualizado)
}

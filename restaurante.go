package main

import "fmt"

func VerificarPedidos(cardapio []string, vendidos []string) ([]string, []string) {
	vendidosSet := make(map[string]bool)
	for _, prato := range vendidos {
		vendidosSet[prato] = true
	}
	var vendidosNoCardapio []string
	var naoVendidos []string
	for _, prato := range cardapio {
		if vendidosSet[prato] {
			vendidosNoCardapio = append(vendidosNoCardapio, prato)
		} else {
			naoVendidos = append(naoVendidos, prato)
		}
	}
	return vendidosNoCardapio, naoVendidos
}

func main() {
	cardapio := []string{"Lasanha", "Risoto", "Salada", "Sopa"}
	vendidos := []string{"Risoto", "Sopa", "Pizza"}
	vendidosCardapio, naoVendidos := VerificarPedidos(cardapio, vendidos)
	fmt.Printf("Pratos vendidos: %v\n", vendidosCardapio)
	fmt.Printf("Pratos faltando: %v\n", naoVendidos)
}

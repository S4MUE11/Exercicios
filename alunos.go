package main

import "fmt"

func Aprovados(alunos map[string]float64) []string {
	aprovados := []string{}
	for nome, nota := range alunos {
		if nota >= 7.0 {
			aprovados = append(aprovados, nome)
		}
	}
	return aprovados
}

func main() {
	alunos := map[string]float64{"Samuel": 9.0, "Caio": 6.5, "Gustavo": 7.0, "Mahgid": 5.8}
	listaAprovados := Aprovados(alunos)
	fmt.Printf("Alunos aprovados: %v\n", listaAprovados)
}

package main

import "fmt"

func main() {

	// Variáveis
	nomeComprador := ""
	var quantidadeBilhetes uint = 0

	// Pergunta 1
	fmt.Printf("Qual o seu nome ? ")
	fmt.Scanln(&nomeComprador)

	// Pergunta 2
	fmt.Printf("Quantos bilhetes você deseja comprar? ")
	fmt.Scanln(&quantidadeBilhetes)

	// Validação
	if quantidadeBilhetes > 1 {
		fmt.Println(nomeComprador, ", você deseja confirmar a compra de ", quantidadeBilhetes, " bilhetes?")
	} else if quantidadeBilhetes == 1 {
		fmt.Println(nomeComprador, ", você deseja confirmar a compra de ", quantidadeBilhetes, " bilhete?")
	} else {
		fmt.Println("Informe um valor válido(ex: 1)")
	}

}

package main

import "fmt"

func main() {
	// mensagem := ""
	nomeComprador := ""
	quantidadeBilhetes := 0.0

	fmt.Printf("Qual o seu nome ? ")
	fmt.Scanln(&nomeComprador)

	fmt.Printf("Quantos bilhetes você deseja comprar: ")
	fmt.Scanln(&quantidadeBilhetes)

}

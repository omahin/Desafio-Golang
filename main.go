package main

import (
	"fmt"
)

type cliente struct {
	nomeComprador      string
	quantidadeBilhetes uint
}

func printFormattedString(compras []cliente, precoIngresso float32, precoTotal float32) {
	fmt.Printf("%-25s %-15s %-10s\n", "Nome", "Quantidade", "Total")
	for _, c := range compras {
		fmt.Printf("%-25s %-15d R$ %.2f\n", c.nomeComprador, c.quantidadeBilhetes, float32(c.quantidadeBilhetes)*precoIngresso)
	}
	fmt.Printf("\nValor total da compra: %-18s R$ %.2f\n", "", precoTotal)
}

func calcularValorTotal(clientes []cliente, valorUnitario float32) float32 {
	var sum float32 = 0.0
	for _, c := range clientes {
		sum += float32(c.quantidadeBilhetes) * valorUnitario
	}
	return sum
}

func isRespostaValida(s string) bool {
	return s == "s" || s == "n"
}

func main() {

	// Variáveis
	const valorBilhete = 10.50

	var structCliente = cliente{}
	var respostaConfirmacao string
	var respostaContinuarCompra string
	var valorTotal float32
	var clientes = []cliente{}

	for {
		// Pergunta 1
		fmt.Printf("Qual o seu nome ? ")
		fmt.Scanln(&structCliente.nomeComprador)

		// Pergunta 2
		fmt.Printf("Quantos bilhetes você deseja comprar? ")
		fmt.Scanln(&structCliente.quantidadeBilhetes)

		for {
			fmt.Printf("Você deseja confirmar a compra? (s/n): ")
			fmt.Scanln(&respostaConfirmacao)
			if isRespostaValida(respostaConfirmacao) {
				break
			} else {
				fmt.Printf("Digite uma resposta valida\n")
			}
		}

		if respostaConfirmacao == "s" {
			clientes = append(clientes, structCliente)
		} else {
			continue
		}

		for {
			fmt.Printf("Deseja comprar mais bilhetes? (s/n): ")
			fmt.Scanln(&respostaContinuarCompra)
			if isRespostaValida(respostaConfirmacao) {
				break
			} else {
				fmt.Printf("Digite uma resposta valida\n")
			}
		}

		if respostaContinuarCompra == "n" {
			valorTotal = calcularValorTotal(clientes, valorBilhete)
			break
		}
	}

	printFormattedString(clientes, valorBilhete, valorTotal)
	print(clientes)
}

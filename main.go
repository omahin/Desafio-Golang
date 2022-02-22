package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type cliente struct {
	nomeComprador      string
	quantidadeBilhetes int
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

	//var structCliente = cliente{}

	var nome string
	var qnt string

	var respostaConfirmacao string
	var respostaContinuarCompra string
	var valorTotal float32
	var clientes = []cliente{}

	in := bufio.NewReader(os.Stdin)

	for {
		// Pergunta 1
		fmt.Printf("Qual o seu nome ? ")
		nome, _ = in.ReadString('\n')

		// Pergunta 2
		for {
			fmt.Printf("Quantos bilhetes você deseja comprar? ")
			fmt.Scanln(&qnt)
			if _, err := strconv.Atoi(qnt); err == nil {
				break
			} else {
				fmt.Printf("Digite um número\n")
			}
		}

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
			intVar, _ := strconv.Atoi(qnt)
			clientes = append(clientes, cliente{
				nomeComprador:      nome,
				quantidadeBilhetes: intVar,
			})
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
}

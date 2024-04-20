package main

import "fmt"

func main() {

	//refatorando o trabalho escolar de 2020/2021 em linguagem C para GO.

	fmt.Println("****************************************************************************")

	fmt.Println("olá, bem vindo à Lancheria do Batata... Escolha qual item deseja comprar:")
	fmt.Println("********************************|MENU|**************************************")

	fmt.Println("item\t\t", "Produtos\t", "Código\t\t", "Preço Unitário\t")

	fmt.Println("****************************************************************************")

	fmt.Println(" 1 -\t Cachorro-quente\t 100\t\t R$ 5,00 ")
	fmt.Println(" 2 -\t X-Salada\t\t 101\t\t R$ 8,79 ")
	fmt.Println(" 3 -\t X-bacon\t\t 102\t\t R$ 9,99 ")
	fmt.Println(" 4 -\t Misto\t\t\t 103\t\t R$ 6,89 ")
	fmt.Println(" 5 -\t Salada\t\t\t 104\t\t R$ 4,80 ")
	fmt.Println(" 6 -\t Água\t\t\t 105\t\t R$ 3,49 ")
	fmt.Println(" 7 -\t Refrigerante\t\t 106\t\t R$ 4,99 ")

	fmt.Println("")
	fmt.Println("Digite a opção desejada de 1 a 7 ou 0 para finalizar o pedido: ")
	fmt.Println("")

	var cachorroquente float32 = 5.00
	var xsalada float32 = 8.79
	var xbacon float32 = 9.99
	var misto float32 = 6.89
	var salada float32 = 4.80
	var agua float32 = 3.49
	var refrigerante float32 = 4.99

	var qcachorro float32 = 0.0
	var qxsalada float32 = 0.0
	var qbacon float32 = 0.0
	var qmisto float32 = 0.0
	var qsalada float32 = 0.0
	var qagua float32 = 0.0
	var qrefri float32 = 0.0
	var qtd1 float32 = 0.0
	var qtd2 float32 = 0.0
	var qtd3 float32 = 0.0
	var qtd4 float32 = 0.0
	var qtd5 float32 = 0.0
	var qtd6 float32 = 0.0
	var qtd7 float32 = 0.0

	opcao := 0

	for {

		fmt.Scan(&opcao)

		switch opcao {

		case 1:

			fmt.Println("Digite a quantidade de cachorro -quente: ")

			fmt.Scan(&qcachorro)
			qtd1 = qcachorro * cachorroquente

			fmt.Println("Algo mais? ... ou 0 para finalizar o pedido: ")

		case 2:

			fmt.Println("Digite a quantidade de X -Salada:")
			fmt.Scan(&qxsalada)
			qtd2 = qxsalada * xsalada
			fmt.Println("Algo mais? ... ou 0 para finalizar o pedido: ")

		case 3:

			fmt.Println("Digite a quantidade de X -bacon: ")
			fmt.Scan(&qbacon)
			qtd3 = qbacon * xbacon
			fmt.Println("Algo mais? ... ou 0 para finalizar o pedido: ")

		case 4:

			fmt.Println("Digite a quantidade de Misto: ")
			fmt.Scan(&qmisto)
			qtd4 = qmisto * misto
			fmt.Println("Algo mais? ... ou 0 para finalizar o pedido: ")

		case 5:

			fmt.Println("Digite a quantidade de Salada: ")
			fmt.Scan(&qsalada)
			qtd5 = qsalada * salada
			fmt.Println("Algo mais? ... ou 0 para finalizar o pedido: ")

		case 6:

			fmt.Println("Digite a quantidade de Água: ")
			fmt.Scan(&qagua)
			qtd6 = qagua * agua
			fmt.Println("Algo mais? ... ou 0 para finalizar o pedido: ")

		case 7:

			fmt.Println("Digite a quantidade de Refrigerante: ")
			fmt.Scan(&qrefri)
			qtd7 = qrefri * refrigerante
			fmt.Println("Algo mais? ... ou 0 para finalizar o pedido: ")

		case 0:

			fmt.Println("****************************************************************************")

			fmt.Println("\t\t\tEMITINDO A NOTA FISCAL")
			fmt.Println("****************************************************************************")

			fmt.Println(" Qtd\t\t Produto\t Valor Unitário\t\t Valor Total ")

			if qcachorro != 0 {
				fmt.Printf(" %.f -\t Cachorro-quente\t R$ %.2f \t\t R$ %.2f", qcachorro, cachorroquente, qtd1)
			}

			fmt.Println("")

			if qxsalada != 0 {
				fmt.Printf(" %.f -\t X-Salada\t\t R$ %.2f \t\t R$ %.2f", qxsalada, xsalada, qtd2)
			}

			fmt.Println("")

			if qbacon != 0 {
				fmt.Printf(" %.f -\t X-bacon\t\t R$ %.2f \t\t R$ %.2f", qbacon, xbacon, qtd3)
			}

			fmt.Println("")

			if qmisto != 0 {
				fmt.Printf(" %.f -\t Misto\t\t\t R$ %.2f \t\t R$ %.2f", qmisto, misto, qtd4)
			}

			fmt.Println("")

			if qsalada != 0 {
				fmt.Printf(" %.f -\t Salada\t\t\t R$ %.2f \t\t R$ %.2f", qsalada, salada, qtd5)
			}

			fmt.Println("")

			if qagua != 0 {
				fmt.Printf(" %.f -\t Água\t\t\t R$ %.2f\t\t R$ %.2f", qagua, agua, qtd6)
			}

			fmt.Println("")

			if qrefri != 0 {
				fmt.Printf(" %.f -\t Refrigerante\t\t R$ %.2f\t\t R$ %.2f", qrefri, refrigerante, qtd7)
			}

			fmt.Println("")

			total := (qtd1 + qtd2 + qtd3 + qtd4 + qtd5 + qtd6 + qtd7)

			fmt.Println("")
			fmt.Printf(" => \tO total do pedido é:\t\t\t\t%.2f reais\n", total)
			fmt.Println("")

			fmt.Println("Obrigado, volte sempre!!!")

			return

		default:			
				opcao = 7
				fmt.Println("Digite um numero válido!!! ")
			
		}		

	}
	
}

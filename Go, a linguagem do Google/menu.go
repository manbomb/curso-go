package main

import "fmt"

func menu() {
	fmt.Println("\nOlá senhor, escolha uma opção:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi: ", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo.")
	} else {
		fmt.Println("Não conheço este comando!")
	}
}

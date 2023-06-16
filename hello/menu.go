package main

import "fmt"

func main() {
	fmt.Println("\nOlá senhor, escolha uma opção:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi: ", comando)
}

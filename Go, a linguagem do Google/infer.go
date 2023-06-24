package main

import (
	"fmt"
	"reflect"
)

func infer() {
	nome := "Sérgio"
	var idade = 21
	var versao = 1.1
	fmt.Println("Olá, sr. ", nome, ", sua idade é ", idade)
	fmt.Println("Este programa esta na versao: ", versao)

	fmt.Println("O tipo da variavel nome é: ", reflect.TypeOf(nome))
}

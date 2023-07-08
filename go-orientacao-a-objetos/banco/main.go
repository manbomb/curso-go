package main

import (
	"curso-go/go-orientacao-a-objetos/banco/clientes"
	"curso-go/go-orientacao-a-objetos/banco/contas"
	"fmt"
)

func main() {
	silvia := clientes.Titular{Nome: "Silvia", CPF: "123", Profissao: "CEO"}
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = silvia
	contaDaSilvia.Depositar(500)

	maria := clientes.Titular{Nome: "Maria", CPF: "321", Profissao: "CTO"}
	contaDaMaria := contas.ContaCorrente{}
	contaDaMaria.Titular = maria
	contaDaMaria.Depositar(500)

	fmt.Println(contaDaSilvia.Saldo())
	fmt.Println(contaDaMaria.Saldo())

	contaDaSilvia.Transferir(500, &contaDaMaria)

	fmt.Println(contaDaSilvia.Saldo())
	fmt.Println(contaDaMaria.Saldo())
}

package contas

import (
	"curso-go/go-orientacao-a-objetos/banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	saldoPositivo := valorDoDeposito > 0
	if saldoPositivo {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Deposito com valor inválido", c.saldo
}

func (origem *ContaCorrente) Transferir(valorDaTransferencia float64, destino *ContaCorrente) bool {
	if valorDaTransferencia <= origem.saldo && valorDaTransferencia > 0 {
		origem.Sacar(valorDaTransferencia)
		destino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}

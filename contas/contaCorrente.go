package contas

import (
	cliente "aluraBank/pessoas/clientes"
)

type ContaCorrente struct {
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (contaCorrente *ContaCorrente) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= contaCorrente.saldo

	if podeSacar {
		contaCorrente.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (contaCorrente *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		contaCorrente.saldo += valorDeposito
		return "Deposito realizado com sucesso", contaCorrente.saldo
	}

	return "Valor do deposito menor que 0", contaCorrente.saldo
}

func (contaOrigem *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia <= contaOrigem.saldo && valorTransferencia > 0 {
		contaOrigem.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

func (contaCorrente *ContaCorrente) GetSaldo() float64 {
	return contaCorrente.saldo
}

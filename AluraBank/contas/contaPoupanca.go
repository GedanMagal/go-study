package contas

import (
	"aluraBank/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (contaPoupanca *ContaPoupanca) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= contaPoupanca.saldo

	if podeSacar {
		contaPoupanca.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (contaPoupanca *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		contaPoupanca.saldo += valorDeposito
		return "Deposito realizado com sucesso", contaPoupanca.saldo
	}

	return "Valor do deposito menor que 0", contaPoupanca.saldo
}

func (contaPoupanca *ContaPoupanca) GetSaldo() float64 {
	return contaPoupanca.saldo
}

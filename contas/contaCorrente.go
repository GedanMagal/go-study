package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (contaCorrente *ContaCorrente) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= contaCorrente.Saldo

	if podeSacar {
		contaCorrente.Saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (contaCorrente *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		contaCorrente.Saldo += valorDeposito
		return "Deposito realizado com sucesso", contaCorrente.Saldo
	}

	return "Valor do deposito menor que 0", contaCorrente.Saldo
}

func (contaOrigem *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia <= contaOrigem.Saldo && valorTransferencia > 0 {
		contaOrigem.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

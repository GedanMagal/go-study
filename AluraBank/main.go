package main

import (
	"aluraBank/clientes"
	"aluraBank/contas"
	"fmt"
)

func main() {

	titularGedan := clientes.Titular{Nome: "Gedan", Cpf: "456.534.266.10", Profissao: "Eng. Software"}
	titularElen := clientes.Titular{Nome: "Elen", Cpf: "432.965.444.12", Profissao: "Professor"}
	contaGedan := contas.ContaCorrente{Titular: titularGedan}
	contaElen := contas.ContaCorrente{Titular: titularElen}

	contaGedan.Depositar(2000)
	contaGedan.Depositar(3000)

	// transferencia de saldo suficiente
	status := contaGedan.Transferir(150000, &contaElen)

	fmt.Println("Status da transferencia: ", status)
	fmt.Println(contaGedan)
	fmt.Println(contaElen)

	// transferencia de saldo insuficiente
	status = contaGedan.Transferir(1000, &contaElen)

	fmt.Println("Status da transferencia: ", status)
	fmt.Println(contaGedan)
	fmt.Println(contaElen)

	// transferencia de saldo suficiente
	status = contaElen.Transferir(8000, &contaGedan)
	fmt.Println("Status da transferencia: ", status)
	fmt.Println(contaGedan)
	fmt.Println(contaElen)

	// transferencia de saldo suficiente
	fmt.Println("Saldo Elen: ", contaElen.GetSaldo())
	fmt.Println("Saldo Gedan: ", contaGedan.GetSaldo())

	// Conta poupança

	gedanPoupanca := contas.ContaPoupanca{Titular: titularGedan}
	fmt.Println("Conta poupança:", gedanPoupanca)

	gedanPoupanca.Depositar(100)
	fmt.Println("Poupança:", gedanPoupanca)
	fmt.Println("Saque:", gedanPoupanca.Sacar(55))
	fmt.Println("Saldo:", gedanPoupanca.GetSaldo())

	contas.PagarBoleto(&gedanPoupanca, 20)

	fmt.Println("Saldo", gedanPoupanca.GetSaldo())

	gedanPoupanca.Depositar(1000)

	contas.PagarBoleto(&gedanPoupanca, 300)

	fmt.Println("Saldo", gedanPoupanca.GetSaldo())

	fmt.Println("Pagamento conta corrente")

	// Pagamento com conta corrente
	contas.PagarBoleto(&contaGedan, 300)

	fmt.Println("Saldo", contaGedan.GetSaldo())

}

// contaGedan := contas.ContaCorrente{}
// contaGedan.Titular = "Gedan"
// contaGedan.Saldo = 500

// fmt.Println("Saldo:", contaGedan.Saldo)

// fmt.Println("\nTeste saque com Saldo positivo")
// fmt.Println("Operação:", contaGedan.Sacar(valorSaque))
// fmt.Println("Saldo atual:", contaGedan.Saldo)

// fmt.Println("\nTeste saque acima do Saldo")
// fmt.Println("Operação:", contaGedan.Sacar(-200))
// fmt.Println("Saldo atual:", contaGedan.Saldo)

// fmt.Println("\nTeste deposito negativo")
// fmt.Println(contaGedan.Depositar(-500))
// fmt.Println("Saldo", contaGedan.Saldo)

// fmt.Println("\nTeste deposito")
// fmt.Println(contaGedan.Depositar(800))

// fmt.Println("Saldo", contaGedan.Saldo)

// fmt.Println("\nTeste retorno duas variaveis")

// statusMessage, valorAtual := contaGedan.Depositar(150)
// fmt.Println("Mensagem", statusMessage)
// fmt.Println("Saldo", valorAtual)

// fmt.Println("Variadict functions")

// fmt.Println(VariadictFunctionSum(1))
// fmt.Println(VariadictFunctionSum(1, 1))
// fmt.Println(VariadictFunctionSum(1, 1, 1))
// fmt.Println(VariadictFunctionSum(1, 1, 2, 4))

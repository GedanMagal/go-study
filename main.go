package main

import (
	"aluraBank/contas"
	"fmt"
)

func main() {

	contaGedan := contas.ContaCorrente{Titular: "Gedan", Saldo: 150000}
	contaElen := contas.ContaCorrente{Titular: "Elen", Saldo: 3000}

	status := contaGedan.Transferir(150000, &contaElen)

	fmt.Println("Status da transferencia: ", status)
	fmt.Println(contaGedan)
	fmt.Println(contaElen)
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

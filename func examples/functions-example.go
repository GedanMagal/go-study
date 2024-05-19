package fmt

import (
	"aluraBank/contas"
	"fmt"
)

func SemParametro() string {
	return "Function sem parametro"
}

func VariadictFunctionSum(numeros ...int) int {
	resultado := 0

	for _, numero := range numeros {
		resultado += numero
	}

	return resultado
}

func Objects() {
	contaCorrenteGedan := contas.ContaCorrente{Titular: "Gedan", NumeroAgencia: 222, NumeroConta: 123456, Saldo: 125.5}

	contaCorrenteElen := contas.ContaCorrente{Titular: "Elen", NumeroAgencia: 333, NumeroConta: 1235543, Saldo: 222.5}

	fmt.Println(contaCorrenteGedan)
	fmt.Println(contaCorrenteElen)

	var contaDaGabi *contas.ContaCorrente
	contaDaGabi = new(contas.ContaCorrente)
	contaDaGabi.Titular = "Gabriele Vitoria"

	fmt.Println(*contaDaGabi)

}

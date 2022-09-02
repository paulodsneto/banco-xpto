package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {
	// contaDoPaulo := ContaCorrente{
	// 	titular:    "Paulo",
	// 	numAgencia: 333,
	// 	numConta:   12345,
	// 	saldo:      100.5,
	// }

	// contaDaAlessandra := ContaCorrente{"Alessandra", 222, 321654, 200}
	// fmt.Println(contaDoPaulo, contaDaAlessandra)

	// var contaDoLula *ContaCorrente
	// contaDoLula = new(ContaCorrente)
	// contaDoLula.titular = "Lula"

	// fmt.Println(*contaDoLula)
	contaDoPaulo := ContaCorrente{}
	contaDoPaulo.titular = "Paulo"
	contaDoPaulo.saldo = 500
	fmt.Println(contaDoPaulo.saldo)

	fmt.Println(contaDoPaulo.Sacar(300), "\nSaldo atualizado:", contaDoPaulo.saldo)
}

func (c *ContaCorrente) Sacar(valordoSaque float64) string {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

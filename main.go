package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {
	contaDoPaulo := ContaCorrente{
		titular:    "Paulo",
		numAgencia: 333,
		numConta:   12345,
		saldo:      100.5,
	}

	contaDaAlessandra := ContaCorrente{"Alessandra", 222, 321654, 200}
	fmt.Println(contaDoPaulo, contaDaAlessandra)

}

package main

import "fmt"

type Conta struct {
	Titular string
	Saldo   float64
}

func adicionarValor(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := Conta{
		Titular: "João",
		Saldo:   100.0,
	}
	valorAdicional := 50.0
	adicionarValor(&conta, valorAdicional)
	fmt.Printf("Saldo da conta de %s; %.2f \n", conta.Titular, conta.Saldo)
}

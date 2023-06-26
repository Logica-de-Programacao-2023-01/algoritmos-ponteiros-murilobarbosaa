package main

import "fmt"

func atualizarValor2(ptr *int) {
	valor := *ptr
	if valor < 0 {
		valor = -valor
	}
	unidade := valor % 10
	dezena := (valor / 10) % 10
	soma := unidade + dezena
	*ptr = soma
}

func main() {
	numero := 1234
	atualizarValor2(&numero)
	fmt.Println("Novo valor:", numero)
}

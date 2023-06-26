package main

import "fmt"

func atualizarValor1(ptr *int, n int) {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	*ptr = soma
}

func main() {
	valor := 0
	n := 5
	atualizarValor1(&valor, n)
	fmt.Println("Novo valor:", valor)
}

package main

import (
	"fmt"
	"math"
)

func atualizarValor3(ptr *float64) {
	*ptr = (*ptr + math.Pi) / 2
}

func main() {
	valor := 3.14
	atualizarValor3(&valor)
	fmt.Println("Novo valor:", valor)
}

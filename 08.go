package main

import "fmt"

func modificarValor(ptr *int) {
	*ptr = 42
}

func main() {
	var valor int
	ptr := &valor

	fmt.Println("Valor antes da modificação:", *ptr)

	modificarValor(ptr)
	fmt.Println("Valor após a modificação:", *ptr)
}

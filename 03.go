package main

import "fmt"

func inverterString(ptr *string) {
	original := *ptr
	runes := []rune(original)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	*ptr = string(runes)
}

func main() {
	texto := "Olá, mundo"
	inverterString(&texto)
	fmt.Println("Texto invertido:", texto)
}

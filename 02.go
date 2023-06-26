package main

import "fmt"

func vericarParImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}

func main() {
	valor := 7
	vericarParImpar(&valor)
	fmt.Println("Novo valor:", valor)
}

package main

import "fmt"

func isPrimo(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func preencherPrimos(s *[]int, tamanho int) {
	primos := []int{}
	num := 2
	for len(primos) < tamanho {
		if isPrimo(num) {
			primos = append(primos, num)
		}
		num++
	}
	*s = primos
}

func main() {
	var primos []int
	tamanho := 10

	preencherPrimos(&primos, tamanho)
	fmt.Printf("Os %d primeiros números primos são: %v \n", tamanho, primos)
}

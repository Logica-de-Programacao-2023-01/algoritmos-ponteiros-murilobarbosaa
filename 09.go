package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	livro.Preco = livro.Preco * 0.9
}

func main() {
	livro := Livro{
		Titulo: "Livro Corte de Rosas e Espinhos",
		Autor:  "Autor Sarah J. Mass",
		Preco:  39.90,
	}

	fmt.Println("Preço antes do desconto:", livro.Preco)

	aplicarDesconto(&livro)
	fmt.Println("Preço após o desconto:", livro.Preco)
}

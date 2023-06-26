package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func atualizarLivro(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro 1",
		Autor:  "Anônimo",
	}
	atualizarLivro(&livro)
	fmt.Println("Título do livro:", livro.Titulo)
	fmt.Println("Autor do livro:", livro.Autor)
}

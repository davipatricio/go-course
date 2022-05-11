package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) SetNome(nome string) {
	p.nome = nome
}

func main() {
	pessoa1 := pessoa{"Alfredo", 30}

	fmt.Println(pessoa1)
}

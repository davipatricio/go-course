package main

import "fmt"

type PessoaManager interface {
	setNome(nome string)
}

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) setNome(nome string) {
	p.nome = nome
}

func nome(p PessoaManager) {
	p.setNome("Mario")
	fmt.Println(&p)
}

func main() {
	pessoa1 := pessoa{"Alfredo", 30}

	nome(&pessoa1)
}

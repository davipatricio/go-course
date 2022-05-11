package main

import "fmt"

type client struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := client{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}

	cliente2 := client{"Joana", "Pereira", true}

	fmt.Println(cliente1, cliente2)
}

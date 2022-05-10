package main

import "fmt"

func main() {
	amigos := map[string]int{
		"alfredo": 555111222,
		"joana":   111999444,
	}
	fmt.Println(amigos)
	fmt.Println("Número de Joana: ", amigos["joana"])
	amigos["gopher"] = 333222333
	fmt.Println("Número de Gopher: ", amigos["gopher"])

	numerodavi, ok := amigos["davi"]
	if ok {
		fmt.Println("Número de Davi: ", numerodavi)
	} else {
		fmt.Println("Não temos o número de Davi :(")
	}

	delete(amigos, "joana")
	fmt.Println("Número de Joana: ", amigos["joana"])
}

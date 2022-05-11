package main

import "fmt"

func main() {
	Hello()

	CustomHello("Boa noite!")

	somarest, qtd := SomarRest(10, 10, 10)
	fmt.Printf("Total da soma: %d - Quantidade de elementos: %d\n", somarest, qtd)
	soma := Somar(10, 10)
	fmt.Printf("Total da soma: %d", soma)
}

func Hello() {
	fmt.Println("Hello World!")
}

func CustomHello(txt string) {
	fmt.Println(txt)
}

func SomarRest(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

func Somar(x, y int) int {
	return x + y
}

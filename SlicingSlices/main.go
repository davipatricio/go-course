package main

import "fmt"

func main() {
	sabores := []string{"pepperoni", "mussarela", "portuguesa", "quatro queijos", "calabresa"}
	// Pegar o item 1 até o item 3
	fatia := sabores[1:3]
	todossabores := sabores[:]
	fmt.Println(fatia)
	fmt.Println(todossabores)

	// Remover o item portuguesa
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}

package main

import "fmt"

func main() {
	quemtanoescritoriohoje := "zezinho"
	switch quemtanoescritoriohoje {
	case "zezinho":
		fmt.Println("hoje quem tá no escritório é o zezinho")
		fallthrough
	case "marquinhos":
		fmt.Println("sempre que o zezinho vem, o marquinhos vem também")
	case "joana":
		fmt.Println("hoje quem tá no escritório é a joana")
	case "maria":
		fmt.Println("hoje quem tá no escritório é a maria")
	}
}

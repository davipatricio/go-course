package main

import "fmt"

func main() {
	quemtanoescritoriohoje := "zezinho"
	switch quemtanoescritoriohoje {
	case "zezinho", "marquinhos":
		fmt.Println("hoje quem tá no escritório é o time 1")
	case "joana", "maria":
		fmt.Println("hoje quem tá no escritório é o time 2")
	}
}

package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{"Mario", 50}

	fmt.Println(x)
}

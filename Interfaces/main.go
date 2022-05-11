package main

import (
	"fmt"
)

type robot interface {
	gettoken() string
}

type bot struct {
	token string
}

func (b bot) gettoken() string {
	return "ABCDEFGHI.Jklmnopq.rstuvwx.yz"
}

func token(b robot) {
	fmt.Println(b)
}

func main() {
	r := bot{"123456789"}
	fmt.Println(r.gettoken())
	token(r)
}

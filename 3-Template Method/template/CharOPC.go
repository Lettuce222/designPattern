package template

import (
	"fmt"
)

type CharOPC struct {
	ch rune
}

func NewCharOPC(ch rune) *CharOPC {
	return &CharOPC{ch}
}

func (c CharOPC) open() {
	fmt.Print("<<")
}

func (c CharOPC) print() {
	fmt.Print(string(c.ch))
}

func (c CharOPC) close() {
	fmt.Println(">>")
}

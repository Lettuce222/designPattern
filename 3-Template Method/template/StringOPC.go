package template

import (
	"fmt"
)

type StringOPC struct {
	str string
}

func NewStringOPC(str string) *StringOPC {
	return &StringOPC{str}
}

func (s StringOPC) open() {
	s.printLine()
}

func (s StringOPC) close() {
	s.printLine()
}

func (s StringOPC) print() {
	fmt.Printf("|%v|\n", s.str)
}

func (s StringOPC) printLine() {
	fmt.Print("+")
	for i := 0; i < len(s.str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

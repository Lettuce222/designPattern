package adapter

import (
	"fmt"
)

type Banner struct {
	String string
}

func (b *Banner) showWithParen() {
	fmt.Printf("{%v}\n", b.String)
}

func (b *Banner) showWithAster() {
	fmt.Printf("*%v*\n", b.String)
}

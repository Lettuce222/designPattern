package main

import (
	"./cor"
)

func main() {
	alice := &cor.SupportMember{Resolver: &cor.NoSupport{}, Name: "Alice"}
	bob := &cor.SupportMember{Resolver: &cor.LimitSupport{Limit: 100}, Name: "Bob"}

	alice.SetNext(bob)
	for i := 0; i < 500; i += 33 {
		alice.Support(&cor.Trouble{Number: i})
	}
}

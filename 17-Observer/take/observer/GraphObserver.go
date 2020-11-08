package observer

import (
	"fmt"
)

type GraphObserver struct{}

func (gro GraphObserver) update(generator NumberGenerator) {
	fmt.Print("GraphObserver: ")
	for i := 0; i < generator.Number(); i++ {
		fmt.Print("*")
	}
	fmt.Println("")
}

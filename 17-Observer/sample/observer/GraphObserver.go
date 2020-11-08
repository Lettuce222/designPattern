package observer

import (
	"fmt"
)

type GraphObserver struct{}

func (gro *GraphObserver) update(ng NumberGetter) {
	fmt.Print("GraphObserver: ")
	for i := 0; i < ng.Number(); i++ {
		fmt.Print("*")
	}
	fmt.Println("")
}

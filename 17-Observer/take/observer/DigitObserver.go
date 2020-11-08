package observer

import (
	"fmt"
)

type DigitObserver struct{}

func (do DigitObserver) update(generator NumberGenerator) {
	fmt.Printf("DigitObserver: %v\n", generator.Number())
}

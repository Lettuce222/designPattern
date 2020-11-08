package observer

import (
	"fmt"
)

type DigitObserver struct{}

func (do *DigitObserver) update(ng NumberGetter) {
	fmt.Printf("DigitObserver: %v\n", ng.Number())
}

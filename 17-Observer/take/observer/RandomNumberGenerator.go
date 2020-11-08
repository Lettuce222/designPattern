package observer

import (
	"math/rand"
	"time"
)

type RandomNumberExecuter struct {
	number int
}

func (rne *RandomNumberExecuter) Number() int {
	return rne.number
}

func (rne *RandomNumberExecuter) innerExecute(NotifyObservers func()) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		rne.number = rand.Intn(50)
		NotifyObservers()
	}
}

package observer

import (
	"math/rand"
	"time"
)

type RandomNumberGenerator struct {
	number int
	*NumberGenerator
}

func NewRandomNumberGenerator() RandomNumberGenerator {
	return RandomNumberGenerator{number: 0, NumberGenerator: &NumberGenerator{observers: []Observer{}}}
}

func (rng *RandomNumberGenerator) Number() int {
	return rng.number
}

func (rng *RandomNumberGenerator) NotifyObservers() {
	for _, v := range rng.observers {
		v.update(rng)
	}
}

func (rng *RandomNumberGenerator) Execute() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		rng.number = rand.Intn(50)
		rng.NotifyObservers()
	}
}

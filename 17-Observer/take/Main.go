package main

import (
	"./observer"
)

func main() {
	generator := observer.NumberGenerator{Observers: []observer.Observer{}, NumberExecuter: &observer.RandomNumberExecuter{}}
	var observer1 observer.Observer = observer.DigitObserver{}
	var observer2 observer.Observer = observer.GraphObserver{}
	generator.AddObserver(observer1)
	generator.AddObserver(observer2)
	generator.Execute()
}

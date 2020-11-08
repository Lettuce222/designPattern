package main

import (
	"./state"
	"time"
)

func main() {
	frame := state.NewSafeFrame()

	for true {
		for hour := 0; hour < 24; hour++ {
			frame.SetClock(hour)
			time.Sleep(time.Second)
		}
	}
}

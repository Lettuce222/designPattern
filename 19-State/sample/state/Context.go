package state

import (
	"fmt"
)

type Context interface {
	SetClock(hour int)
	ChangeState(state State)
	CallSecurityCenter(msg string)
	RecordLog(msg string)
}

type SafeFrame struct {
	state State
}

func (s *SafeFrame) SetClock(hour int) {
	clockstring := "現在時刻は"
	if hour < 10 {
		clockstring += fmt.Sprintf("0%v:00", hour)
	} else {
		clockstring += fmt.Sprintf("%v:00", hour)
	}

	fmt.Println(clockstring)
	s.state.DoClock(s, hour)
	s.state.DoAlarm(s)
	s.state.DoPhone(s)
	s.state.DoUse(s)
	fmt.Println()
}

func (s *SafeFrame) ChangeState(state State) {
	fmt.Printf("%vから%vへ状態が変化しました。\n", s.state, state)
	s.state = state
}

func (s SafeFrame) CallSecurityCenter(msg string) {
	fmt.Printf("call! %v\n", msg)
}

func (s SafeFrame) RecordLog(msg string) {
	fmt.Printf("record ... %v\n", msg)
}

func NewSafeFrame() *SafeFrame {
	return &SafeFrame{state: GetDayState()}
}

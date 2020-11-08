package state

type LunchState struct{}

func (d *LunchState) DoClock(c Context, hour int) {
	if hour != 12 {
		c.ChangeState(GetDayState())
	}
}

func (d *LunchState) DoUse(c Context) {
	c.CallSecurityCenter("金庫使用（昼食時）")
}

func (d *LunchState) DoAlarm(c Context) {
	c.CallSecurityCenter("非常ベル（昼食時）")
}

func (d *LunchState) DoPhone(c Context) {
	c.RecordLog("通常の通話（昼食時）")
}

func (d *LunchState) String() string {
	return "昼食時"
}

func newLunchState() *LunchState {
	return &LunchState{}
}

func GetLunchState() *LunchState {
	return lunchstate
}

var lunchstate *LunchState = newLunchState()

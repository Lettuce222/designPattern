package state

type DayState struct{}

func (d *DayState) DoClock(c Context, hour int) {
	if hour < 9 || 17 <= hour {
		c.ChangeState(GetNightState())
	}
}

func (d *DayState) DoUse(c Context) {
	c.RecordLog("金庫使用（昼間）")
}

func (d *DayState) DoAlarm(c Context) {
	c.CallSecurityCenter("非常ベル（昼間）")
	c.ChangeState(GetEmergencyState())
}

func (d *DayState) DoPhone(c Context) {
	c.CallSecurityCenter("通常の通話（昼間）")
}

func (d *DayState) String() string {
	return "昼間"
}

func newDayState() *DayState {
	return &DayState{}
}

func GetDayState() *DayState {
	return daystate
}

var daystate *DayState = newDayState()

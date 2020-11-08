package state

type NightState struct{}

func (n NightState) DoClock(c Context, hour int) {
	if 9 <= hour && hour < 17 {
		c.ChangeState(GetDayState())
	}
}

func (n NightState) DoUse(c Context) {
	c.CallSecurityCenter("金庫使用（夜間）")
}

func (n NightState) DoAlarm(c Context) {
	c.CallSecurityCenter("非常ベル（夜間）")
	c.ChangeState(GetEmergencyState())
}

func (n NightState) DoPhone(c Context) {
	c.RecordLog("通常の通話（夜間）")
}

func (n NightState) String() string {
	return "夜間"
}

func newNightState() *NightState {
	return &NightState{}
}

func GetNightState() *NightState {
	return nightstate
}

var nightstate *NightState = newNightState()

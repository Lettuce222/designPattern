package state

type EmergencyState struct{}

func (d *EmergencyState) DoClock(c Context, hour int) {}

func (d *EmergencyState) DoUse(c Context) {
	c.CallSecurityCenter("金庫使用（緊急時）")
}

func (d *EmergencyState) DoAlarm(c Context) {
	c.CallSecurityCenter("非常ベル（緊急時）")
}

func (d *EmergencyState) DoPhone(c Context) {
	c.CallSecurityCenter("通常の通話（緊急時）")
}

func (d *EmergencyState) String() string {
	return "緊急時"
}

func newEmergencyState() *EmergencyState {
	return &EmergencyState{}
}

func GetEmergencyState() *EmergencyState {
	return emergency
}

var emergency *EmergencyState = newEmergencyState()

package gohealth

import "testing"

func Test_MonitorSwitch_OK(t *testing.T) {

	m := NewMonitorSwitch("red light")

	DeepEqual("off", m.GetStatus(), t)
	DeepEqual([]*Alarm{}, m.GetAlarms(), t)

	m.On()

	DeepEqual("on", m.GetStatus(), t)
	DeepEqual("red light", m.GetAlarms()[0].Msg, t)

	m.On()
	m.On()

	DeepEqual("on", m.GetStatus(), t)
	DeepEqual("red light", m.GetAlarms()[0].Msg, t)

	m.Off()

	DeepEqual("off", m.GetStatus(), t)
	DeepEqual([]*Alarm{}, m.GetAlarms(), t)

}

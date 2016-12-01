package gohealth

import "testing"

func Test_MonitorSwitch_OK(t *testing.T) {

	m := NewMonitorSwitch("red light")

	DeepEqual("ok", m.GetStatus(), t)
	DeepEqual([]*Alarm{}, m.GetAlarms(), t)

	m.Error()

	DeepEqual("error", m.GetStatus(), t)
	DeepEqual("red light", m.GetAlarms()[0].Msg, t)

	m.Error()
	m.Error()

	DeepEqual("error", m.GetStatus(), t)
	DeepEqual("red light", m.GetAlarms()[0].Msg, t)

	m.Ok()

	DeepEqual("ok", m.GetStatus(), t)
	DeepEqual([]*Alarm{}, m.GetAlarms(), t)

}

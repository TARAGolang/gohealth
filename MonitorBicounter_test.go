package gohealth

import "testing"

func Test_MonitorBicounter_GetAlarms(t *testing.T) {
	m := NewMonitorBicounter(6, 10)

	// send 8 errs
	for i := 0; i < 8; i++ {
		m.Err()
	}

	// Check
	alarms := m.GetAlarms()
	DeepEqual("Limit at 7 out of 10", alarms[0].Msg, t)
	DeepEqual("Limit at 8 out of 10", alarms[1].Msg, t)

	// Get alarms should empty alarms
	alarms = m.GetAlarms()
	DeepEqual([]*Alarm{}, alarms, t)

	// send 8 oks
	for i := 0; i < 8; i++ {
		m.Ok()
	}

	// Check 1 alarm
	alarms = m.GetAlarms()
	DeepEqual("Limit at 8 out of 10", alarms[0].Msg, t)
	DeepEqual("Limit at 8 out of 10", alarms[1].Msg, t)
	DeepEqual("Limit at 7 out of 10", alarms[2].Msg, t)

}

func Test_MonitorBicounter_GetStatus(t *testing.T) {
	m := NewMonitorBicounter(2, 3)

	// send 2 errs
	for i := 0; i < 2; i++ {
		m.Err()
	}

	expected := map[string]interface{}{
		"ok":    0,
		"error": 2,
		"p":     66,
	}

	DeepEqual(expected, m.GetStatus(), t)
}
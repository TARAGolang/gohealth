package gohealth

import (
	"fmt"
	"testing"
)

func Test_MonitorBicounter_GetAlarms(t *testing.T) {
	m := NewMonitorBicounter(6, 10)

	// send 8 errs
	for i := 0; i < 8; i++ {
		m.Error()
	}

	// Check
	alarms := m.GetAlarms()
	DeepEqual("Limit at 8 out of 10", alarms[0].Msg, t)

	// Get alarms should NOT empty alarms
	alarms = m.GetAlarms()
	DeepEqual("Limit at 8 out of 10", alarms[0].Msg, t)

	// send 8 oks
	for i := 0; i < 8; i++ {
		m.Ok()
	}

	// Check alarm is empty
	alarms = m.GetAlarms()
	DeepEqual([]*Alarm{}, alarms, t)

}

func Test_MonitorBicounter_GetStatus(t *testing.T) {
	m := NewMonitorBicounter(2, 3)

	// send 2 errs
	for i := 0; i < 2; i++ {
		m.Error()
	}

	expected := map[string]interface{}{
		"ok":    0,
		"error": 2,
		"p":     66,
	}

	DeepEqual(expected, m.GetStatus(), t)
}

func ExampleMonitorBicounter() {
	max_allowed_bad_events := 3
	memorize_last_events := 100
	m := NewMonitorBicounter(max_allowed_bad_events, memorize_last_events)

	m.Ok()
	m.Ok()
	m.Error()
	m.Error()
	m.Error()
	m.Ok()
	m.Ok()
	m.Error() // this event will generate alarm
	m.Ok()    // following Ok generate alarms until last 100 errors are under 4
	m.Ok()
	m.Ok()

	for _, alarm := range m.GetAlarms() {
		fmt.Println(alarm.Msg)
	}

	// Output:
	// Limit at 4 out of 100
}

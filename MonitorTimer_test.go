package gohealth

import (
	"testing"
	"time"
)

func Test_MonitorTimer_GetAlarms(t *testing.T) {

	bicounter := NewMonitorBicounter(0, 1)

	m := NewMonitorTimer(1*time.Hour, bicounter, func(m Monitorer) {
		// See something...
		bicounter.Error()
	})

	m.tick()

	if 1 != len(m.GetAlarms()) {
		t.Error("Should have one alarm")
	}

	bicounter.Ok()

	if 0 != len(m.GetAlarms()) {
		t.Error("Should have zero alarms")
	}
}

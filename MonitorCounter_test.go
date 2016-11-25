package gohealth

import (
	"reflect"
	"testing"
)

func Test_MonitorCounter_GetAlarms(t *testing.T) {

	m := NewMonitorCounter(90, 100)

	if !reflect.DeepEqual(m.GetAlarms(), []*Alarm{}) {
		t.Error("Alarms should be empty")
	}

	m.Increment(50)
	if !reflect.DeepEqual(m.GetAlarms(), []*Alarm{}) {
		t.Error("Alarms should be empty under the threshold")
	}

	m.Increment(45)

	a := m.GetAlarms()[0]
	expected := "Counter is reaching the limit (95 out of 100)"
	if expected != a.Msg {
		t.Error("Unexpected alarm message")
	}

}

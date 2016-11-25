package gohealth

import "testing"

func Test_Alarm_New(t *testing.T) {
	a := NewAlarm("alarm message")

	if "alarm message" != a.Msg {
		t.Error("New alarm should set alarm message")
	}

	if SeverityCritical != a.Severity {
		t.Error("Default severity is CRITICAL")
	}
}

package gohealth

import (
	"testing"
	"time"
)

func Test_MonitorWatch_GetStatus(t *testing.T) {

	// Setup
	counter1 := NewMonitorCounter(90, 100)
	counter2 := NewMonitorCounter(20, 100)

	w := NewMonitorWatch()

	w.Register("counter_1", counter1)
	w.Register("counter_2", counter2)

	// Increment counters
	counter1.Increment(3)
	counter2.Increment(50)

	// Simulate time passed
	w.tick()

	// Check
	status := w.GetStatus()
	expected := map[string]interface{}{
		"counter_1": map[string]interface{}{
			"counter":   int32(3),
			"threshold": int32(90),
			"p":         int32(3),
		},
		"counter_2": map[string]interface{}{
			"counter":   int32(50),
			"threshold": int32(20),
			"p":         int32(50),
		},
	}
	DeepEqual(expected, status, t)

}

func Test_MonitorWatch_GetAlarms(t *testing.T) {

	// Setup
	counter1 := NewMonitorCounter(90, 100)

	w := NewMonitorWatch()
	w.Register("counter_1", counter1)

	// Simulate time passed
	counter1.Increment(95)
	w.tick()
	counter1.Increment(1)
	w.tick()

	// Check
	a := w.GetAlarms()

	if 1 != len(a) {
		t.Error("Only 1 alarm is expected")
	}

	expected := "Counter is reaching the limit (96 out of 100)"
	if expected != a[0].Msg {
		t.Error("Unexpected alarm message")
	}

}

// This should clean old alarms and print OK message restoration
func Test_MonitorWatch_CleanAlarmsAfterTick(t *testing.T) {

	w := NewMonitorWatch()

	now := time.Now()

	// Add alarms
	w.alarms["new"] = &lastAlarm{now, NewAlarm("new alarm")}
	w.alarms["old"] = &lastAlarm{now, NewAlarm("old alarm")}

	// Fake old alarm
	w.alarms["old"].Time = w.alarms["old"].Time.Add(time.Duration(-30 * time.Second))

	// Fake print
	last_alarm := NewAlarm("last alarm")
	w.Print = func(a *Alarm) {
		last_alarm = a
	}

	// Simulate time passed
	w.tick()

	// Check new alarm is still in alarms
	if "new alarm" != w.alarms["new"].Alarm.Msg {
		t.Error("New alarm should be inside the monitor watch")
	}

	// Check if old alarm has been restored
	if "old alarm" != last_alarm.Msg {
		t.Error("Old alarm should be last one")
	}

	// Check alarm restoration
	if SeverityOK != last_alarm.Severity {
		t.Error("Old alarm should have been restored")
	}

}

func Test_MonitorWatch_Issue1(t *testing.T) {

	p := &MonitorPrintTest{}

	w := NewMonitorWatch()
	w.TickDelay = 100 * time.Millisecond
	w.CautionDelay = 10 * time.Millisecond
	w.Print = p.Print

	s := NewMonitorSwitch("my switch")
	s.Error()

	w.Register("monitor my switch", s)
	w.Start()

	time.Sleep(1 * time.Second)

	last := len(p.Alarms) - 1

	if "monitor my switch: OK" == p.Alarms[last] {
		t.Error("Alarm status should not be OK after `CautionDelay` time")
	}
}

// MonitorPrintTest is a helper object that provides a mocked `Print` function
// and stores all print messages.
type MonitorPrintTest struct {
	Alarms []string
}

func (m *MonitorPrintTest) Print(a *Alarm) {
	s := a.Name + ": " + a.Severity
	m.Alarms = append(m.Alarms, s)
}

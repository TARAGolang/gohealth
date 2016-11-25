package gohealth

import "fmt"

/**
 * Binary counter
 */
type MonitorBicounter struct {
	*MonitorMulticounter
	size      int
	threshold int
	alarms    []*Alarm
}

func NewMonitorBicounter(threshold, size int) *MonitorBicounter {
	m := &MonitorBicounter{
		MonitorMulticounter: NewMonitorMulticounter(size),
		threshold:           threshold,
		size:                size,
		alarms:              []*Alarm{},
	}

	m.MonitorMulticounter.OnChange = m.onChange

	return m
}

func (m *MonitorBicounter) Ok() {
	m.MonitorMulticounter.Increment("ok")
}

func (m *MonitorBicounter) Err() {
	m.MonitorMulticounter.Increment("error")
}

func (m *MonitorBicounter) onChange(s map[string]int) {

	e := s["error"]
	if e > m.threshold {
		msg := fmt.Sprintf("Limit at %d out of %d", e, m.size)
		a := NewAlarm(msg)

		m.alarms = append(m.alarms, a)
	}
}

func (m *MonitorBicounter) GetStatus() interface{} {
	s := m.MonitorMulticounter.GetCounters()
	oks := s["ok"]
	errors := s["error"]
	return map[string]interface{}{
		"ok":    oks,
		"error": errors,
		"p":     100 * errors / m.size,
	}
}

func (m *MonitorBicounter) GetAlarms() []*Alarm {
	alarms := m.alarms
	m.alarms = []*Alarm{}

	return alarms
}

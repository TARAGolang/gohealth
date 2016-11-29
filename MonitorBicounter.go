package gohealth

import "fmt"

/**
 * Binary counter
 */
type MonitorBicounter struct {
	*MonitorMulticounter
	size      int
	threshold int
	alarm     *Alarm
}

func NewMonitorBicounter(threshold, size int) *MonitorBicounter {
	m := &MonitorBicounter{
		MonitorMulticounter: NewMonitorMulticounter(size),
		threshold:           threshold,
		size:                size,
	}

	m.MonitorMulticounter.OnChange = m.onChange

	return m
}

func (m *MonitorBicounter) Ok() {
	m.MonitorMulticounter.Increment("ok")
}

func (m *MonitorBicounter) Error() {
	m.MonitorMulticounter.Increment("error")
}

func (m *MonitorBicounter) onChange(s map[string]int) {

	e := s["error"]
	if e > m.threshold {
		msg := fmt.Sprintf("Limit at %d out of %d", e, m.size)
		m.alarm = NewAlarm(msg)
	} else {
		m.alarm = nil
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
	if nil == m.alarm {
		return []*Alarm{}
	}

	return []*Alarm{m.alarm}
}

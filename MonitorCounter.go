package gohealth

import "fmt"

type MonitorCounter struct {
	Monitor
	threshold int32
	counter   int32
	max       int32
	alarm     *Alarm
}

func NewMonitorCounter(threshold, max int32) *MonitorCounter {
	return &MonitorCounter{
		threshold: threshold,
		max:       max,
		alarm:     nil,
	}
}

func (m *MonitorCounter) Increment(i int32) {

	v := m.counter + i // value to set

	if v < 0 {
		m.counter = 0 // TODO: make this atomic :D
	} else if v > m.max {
		m.counter = m.max // TODO: make this atomic :D
	} else {
		m.counter = v // TODO: make this atomic :D :D
	}

	if m.counter > m.threshold {
		// TODO: if alarm exists, update instead of create!
		msg := fmt.Sprintf(
			"Counter is reaching the limit (%d out of %d)",
			m.counter, m.max,
		)
		a := NewAlarm(msg)

		m.alarm = a
	} else {
		m.alarm = nil
	}
}

func (m *MonitorCounter) GetAlarms() []*Alarm {
	r := []*Alarm{}

	if nil != m.alarm {
		r = append(r, m.alarm)
	}

	return r
}

func (m *MonitorCounter) Reset() {
	m.counter = 0 // TODO: make this atomic
}

func (m *MonitorCounter) GetStatus() interface{} {
	return map[string]interface{}{
		"counter":   m.counter,
		"threshold": m.threshold,
		"p":         100 * m.counter / m.max,
	}
}

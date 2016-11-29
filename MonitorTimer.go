package gohealth

import "time"

type MonitorTimer struct {
	Monitor
	Interval time.Duration
	callback func(Monitorer)
	monitor  Monitorer
}

func NewMonitorTimer(interval time.Duration, monitor Monitorer, callback func(Monitorer)) *MonitorTimer {

	m := &MonitorTimer{
		Interval: interval,
		monitor:  monitor,
		callback: callback,
	}

	go func() {
		for {
			time.Sleep(m.Interval)
			m.tick()
		}
	}()

	return m
}

func (m *MonitorTimer) tick() {
	m.callback(m.monitor)
}

func (m *MonitorTimer) GetAlarms() []*Alarm {
	return m.monitor.GetAlarms()
}

func (m *MonitorTimer) GetStatus() interface{} {
	return m.monitor.GetStatus()
}

package gohealth

type MonitorSwitch struct {
	Monitor
	Message string
	alarm   *Alarm
}

func NewMonitorSwitch(message string) *MonitorSwitch {
	return &MonitorSwitch{
		Message: message,
	}
}

func (m *MonitorSwitch) Error() {
	m.alarm = NewAlarm(m.Message)
}

func (m *MonitorSwitch) Ok() {
	m.alarm = nil
}

func (m *MonitorSwitch) Toggle() {
	if nil == m.alarm {
		m.Ok()
	} else {
		m.Error()
	}
}

func (m *MonitorSwitch) GetAlarms() []*Alarm {
	if nil == m.alarm {
		return []*Alarm{}
	}

	return []*Alarm{m.alarm}
}

func (m *MonitorSwitch) GetStatus() interface{} {
	if nil == m.alarm {
		return "ok"
	}

	return "error"
}

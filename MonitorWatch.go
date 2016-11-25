package gohealth

import "time"

type MonitorWatch struct {
	Monitor // Inherit from standard monitor

	TickDelay    time.Duration
	CautionDelay time.Duration
	Print        func(*Alarm) // Print callback

	alarms   map[string]*Alarm
	monitors map[string]*Monitorer
	run      bool
}

func NewMonitorWatch() *MonitorWatch {
	return &MonitorWatch{
		alarms:       map[string]*Alarm{},
		monitors:     map[string]*Monitorer{},
		run:          false,
		TickDelay:    1 * time.Second,
		CautionDelay: 20 * time.Second,
		Print:        PrintSmartDigits,
	}
}

func (m *MonitorWatch) Start() {

	if m.run {
		return
	}

	m.run = true
	go func() {
		for m.run {
			m.tick()
			time.Sleep(m.TickDelay)
		}
	}()
}

func (m *MonitorWatch) Stop() {
	m.run = false
}

func (m *MonitorWatch) tick() {

	// Update alarm
	for name, monitor := range m.monitors {
		for _, a := range (*monitor).GetAlarms() {
			if "" == a.Name {
				a.Name = name
			}
			m.alarms[name] = a
			m.Print(a)
		}
	}

	// Clean alarms
	for k, a := range m.alarms {
		if a.OlderThan(m.CautionDelay) {
			a.Severity = SeverityOK
			m.Print(a)
			delete(m.alarms, k)
		}
	}

}

func (m *MonitorWatch) Register(name string, r Monitorer) {
	m.monitors[name] = &r
}

func (m *MonitorWatch) GetAlarms() []*Alarm {
	r := []*Alarm{}
	for _, a := range m.alarms {
		r = append(r, a)
	}

	return r
}

func (m *MonitorWatch) GetStatus() interface{} {
	r := map[string]interface{}{}

	for k, v := range m.monitors {
		r[k] = (*v).GetStatus()
	}

	return r
}

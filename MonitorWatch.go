package gohealth

import (
	"fmt"
	"time"
)

type MonitorWatch struct {
	Monitor // Inherit from standard monitor

	TickDelay    time.Duration
	CautionDelay time.Duration
	Print        func(*Alarm) // Print callback

	alarms   map[string]*lastAlarm
	monitors map[string]*Monitorer
	status   map[string]interface{}

	run bool
}

type lastAlarm struct {
	// Time is when the alarm was collected
	Time time.Time

	// Alarm is a pointer to the alarm itself
	Alarm *Alarm
}

func (a *lastAlarm) OlderThan(duration time.Duration) bool {
	return time.Now().Sub(a.Time) > duration
}

func NewMonitorWatch() *MonitorWatch {
	return &MonitorWatch{
		alarms:       map[string]*lastAlarm{},
		monitors:     map[string]*Monitorer{},
		status:       map[string]interface{}{},
		run:          false,
		TickDelay:    1 * time.Second,
		CautionDelay: 20 * time.Second,
		Print:        PrintSmartDigits,
	}
}

func (m *MonitorWatch) Start() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()

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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()

	// Update alarm
	for name, monitor := range m.monitors {
		for _, a := range (*monitor).GetAlarms() {
			if "" == a.Name {
				a.Name = name
			}
			m.Print(a)
			m.alarms[name] = &lastAlarm{time.Now(), a}
		}
	}

	// Clean alarms
	for k, l := range m.alarms {
		if l.OlderThan(m.CautionDelay) {
			l.Alarm.Severity = SeverityOK
			m.Print(l.Alarm)
			delete(m.alarms, k)
		}
	}

	// Update status
	for k, v := range m.monitors {
		m.status[k] = (*v).GetStatus()
	}

}

func (m *MonitorWatch) Register(name string, r Monitorer) {
	m.monitors[name] = &r
}

func (m *MonitorWatch) GetAlarms() []*Alarm {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()

	r := []*Alarm{}
	for _, l := range m.alarms {
		r = append(r, l.Alarm)
	}

	return r
}

func (m *MonitorWatch) GetStatus() interface{} {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()

	return m.status
}

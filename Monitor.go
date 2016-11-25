package gohealth

type Monitor struct {
}

func (m *Monitor) GetAlarms() []Alarm {
	return []Alarm{}
}

func (m *Monitor) GetStatus() interface{} {
	return nil
}

func (m *Monitor) Start() {
	// Do nothing
}

func (m *Monitor) Stop() {
	// Do nothing
}

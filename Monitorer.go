package gohealth

type Monitorer interface {
	GetAlarms() []*Alarm
	GetStatus() interface{}
}

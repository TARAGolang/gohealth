package gohealth

import "time"

const (
	SeverityCritical = "CRITICAL"
	SeverityWarnig   = "WARNING"
	SeverityOK       = "OK"
)

type Alarm struct {
	Time     time.Time
	Name     string
	Severity string
	Msg      string
}

func NewAlarm(msg string) *Alarm {
	return &Alarm{
		Time:     time.Now(),
		Name:     "",
		Severity: SeverityCritical,
		Msg:      msg,
	}
}

func (a *Alarm) OlderThan(duration time.Duration) bool {
	return time.Now().Sub(a.Time) > duration
}

package gohealth

import (
	"encoding/json"
	"time"
)

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

// Print should print an alarm like this (but in one line):
// {
//	"time": "2016-11-11 15:07:07.854Z",
//	"type": "ALARM",
//	"payload": {
//		"monitor": "idle",
//		"severity": "OK",
//		"msg": "Last event received at 2016-11-11 15:07:07.854307"
//     }
// }
func (a *Alarm) Print() string {
	v := map[string]interface{}{
		"time": a.Time.UTC().Format(time.RFC3339),
		"type": "ALARM",
		"payload": map[string]interface{}{
			"monitor":  a.Name,
			"severity": a.Severity,
			"msg":      a.Msg,
		},
	}

	j, _ := json.Marshal(v) // NO ERROR VALIDATION :D

	return string(j)
}

func (a *Alarm) OlderThan(duration time.Duration) bool {
	return time.Now().Sub(a.Time) > duration
}

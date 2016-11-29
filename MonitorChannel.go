package gohealth

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * Binary counter
 */
type MonitorChannel struct {
	*Monitor

	Interval time.Duration

	size           int
	threshold      int
	lifetime       int64
	callback_usage func() int

	status map[string]interface{}

	alarm *Alarm

	last_channel_use int
	last_time        time.Time
}

func NewMonitorChannel(size int, threshold int, lifetime int, callback_usage func() int) *MonitorChannel {
	m := &MonitorChannel{
		Interval: 1 * time.Second,

		size:           size,
		threshold:      threshold,
		lifetime:       int64(lifetime),
		callback_usage: callback_usage,

		status: map[string]interface{}{},
	}

	m.updateLastValues()

	go func() {
		m.tick()
		time.Sleep(m.Interval)
	}()

	return m
}

func (m *MonitorChannel) tick() {
	// Get channel usage
	channel_size := m.size
	channel_use := m.callback_usage()
	channel_use_p := 100 * channel_use / channel_size
	channel_free := channel_size - channel_use
	channel_free_p := 100 * channel_free / channel_size

	// Check threshold alarm
	alarm_threshold := channel_use > m.threshold // ALARM

	// Check threshold lifetime
	delay_use := channel_use - m.last_channel_use
	delay_timestamp := time.Now().UnixNano() - m.last_time.UnixNano()
	rate := float32(delay_use) / (float32(delay_timestamp) / 1000000000)
	alarm_time := false
	life := "eternity"
	if delay_use > 0 {
		seconds_left := int64(float32(channel_free) / (rate))
		duration := time.Duration(seconds_left * 1000000000)
		life = fmt.Sprintf("%v", duration)

		alarm_time = seconds_left < m.lifetime // ALARM
	}

	if alarm_time || alarm_threshold {
		msg := "Running out of buffer (" + strconv.Itoa(channel_use_p) +
			"%) time to solve: " + life
		m.alarm = NewAlarm(msg)
	} else {
		m.alarm = nil
	}

	// Update status
	m.status["channel_size"] = channel_size
	m.status["channel_use"] = channel_use
	m.status["channel_use_p"] = channel_use_p
	m.status["channel_free"] = channel_free
	m.status["channel_free_p"] = channel_free_p
	m.status["life"] = life

	// Update last values
	m.updateLastValues()
}

func (m *MonitorChannel) updateLastValues() {
	m.last_channel_use = m.callback_usage()
	m.last_time = time.Now()
}

func (m *MonitorChannel) GetStatus() interface{} {
	return "TODO: implement this"
}

func (m *MonitorChannel) GetAlarms() []*Alarm {
	if nil == m.alarm {
		return []*Alarm{}
	}

	return []*Alarm{m.alarm}
}

package gohealth

import (
	"fmt"
	"testing"
)

func Test_MonitorChannel_GetStatus(t *testing.T) {

	size := 10

	c := make(chan int, size)

	threshold := 6
	lifetime := 100

	m := NewMonitorChannel(size, threshold, lifetime, func() int {
		return len(c)
	})

	// Check status
	m.tick()
	expected_status := map[string]interface{}{
		"channel_free_p": 100,
		"life":           "eternity",
		"channel_size":   10,
		"channel_use":    0,
		"channel_use_p":  0,
		"channel_free":   10,
	}
	DeepEqual(expected_status, m.status, t)
	DeepEqual((*Alarm)(nil), m.alarm, t)

	// Get threshold alarm
	for i := 0; i < 7; i++ {
		c <- i
	}
	m.tick()

	fmt.Println("status:", m.status)

	fmt.Println("alarms:", m.alarm)

	// t.FailNow()

}

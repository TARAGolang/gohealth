package gohealth

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_MonitorMulticounter_Increment(t *testing.T) {

	// Simulation, window size: 4, values: A, B, C
	//
	//                  11111111
	// #:     012345678901234567
	input := "ABCAACCBBBBBABBBBB"
	cou_a := "111222210000111100"
	cou_b := "011110012344333344"
	cou_c := "001112222100000000"
	chang := "^^^^_^_^^^^_^___^_" // (^ means event is triggered)
	//
	// NOTICE 1: steps 4 and 5 has the same status, so the Change event
	// should be triggered only once.
	// NOTICE 2: step 6 B is out of queue (CAAC), status change (B:0)
	// NOTICE 3: step 7 does not trigger a change event (same as notice 1)

	m := NewMonitorMulticounter(4)

	last_status := map[string]int{}
	m.OnChange = func(s map[string]int) {
		last_status = s
	}

	// Initial check
	DeepEqual(map[string]int{}, m.GetStatus(), t)

	// Cases

	for step := 0; step < 18; step++ {

		// Get step values
		i := string(input[step])
		e := string(chang[step])
		a, _ := strconv.Atoi(string(cou_a[step]))
		b, _ := strconv.Atoi(string(cou_b[step]))
		c, _ := strconv.Atoi(string(cou_c[step]))

		// Run
		last_status = nil
		m.Increment(i)

		// Check
		if "_" == e {
			DeepEqual(map[string]int(nil), last_status, t)
		} else {
			DeepEqual(a, last_status["A"], t)
			DeepEqual(b, last_status["B"], t)
			DeepEqual(c, last_status["C"], t)
		}

		// Uncomment this for debugging:
		// fmt.Println("step:", step, "\tinput:", i, "-->", last_status)
	}

}

func Test_MonitorMulticounter_Performance(t *testing.T) {
	m := NewMonitorMulticounter(100)

	tags := []string{"ok", "warning", "error"}
	tags_len := len(tags)

	t1 := time.Now()

	for i := 0; i < 100000; i++ {
		tag := tags[rand.Intn(tags_len)]
		m.Increment(tag)
	}

	delay := time.Now().Sub(t1)

	if delay.Seconds() > float64(1) {
		t.Error("MonitorMulticounter is too slow:", delay)
	}
}

func Test_MonitorMulticounter_Concurrency(t *testing.T) {
	// TODO: how do I test this???

}

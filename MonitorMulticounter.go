package gohealth

import (
	"container/list"
	"runtime"
	"sync"
)

type MonitorMulticounter struct {
	Monitor
	OnChange func(map[string]int)
	size     int
	list     *list.List
	counters map[string]int
	mutex    *sync.Mutex
}

func NewMonitorMulticounter(size int) *MonitorMulticounter {

	return &MonitorMulticounter{
		size:     size,
		list:     list.New(),
		counters: map[string]int{},
		mutex:    &sync.Mutex{},
	}
}

func (m *MonitorMulticounter) Increment(tag string) {
	m.mutex.Lock()

	m.incrementUnsafe(tag)

	m.mutex.Unlock()
	runtime.Gosched()
}

func (m *MonitorMulticounter) incrementUnsafe(tag string) {

	m.list.PushFront(tag)

	if m.list.Len() > m.size {
		last := m.list.Back()
		m.list.Remove(last)

		last_tag := last.Value.(string)

		if last_tag == tag {
			// Counters does not change
			return
		}

		m.counters[last_tag] = m.counters[last_tag] - 1
	}

	m.counters[tag] = m.counters[tag] + 1

	// Change !
	if nil != m.OnChange {
		m.OnChange(m.counters)
	}
}

func (m *MonitorMulticounter) GetStatus() interface{} {
	return m.counters
}

func (m *MonitorMulticounter) GetCounters() map[string]int {
	return m.counters
}

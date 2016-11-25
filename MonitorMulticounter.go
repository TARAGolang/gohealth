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
	list     list.List
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

	if nil != m.OnChange {
		m.OnChange(m.counters)
	}

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
			return
		}

		m.counters[last_tag] = m.counters[last_tag] - 1
	}

	m.counters[tag] = m.counters[tag] + 1
}

func (m *MonitorMulticounter) GetStatus() interface{} {
	return m.counters
}

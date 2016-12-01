package gohealth

import "runtime"

type MonitorMemory struct {
	Monitor
	stats *runtime.MemStats
}

func NewMonitorMemory() *MonitorMemory {
	return &MonitorMemory{
		stats: &runtime.MemStats{},
	}
}

func (m *MonitorMemory) GetStatus() interface{} {
	runtime.ReadMemStats(m.stats)

	return map[string]interface{}{
		"sys":      m.stats.Sys,
		"heap_sys": m.stats.HeapSys,
		"gc_next":  m.stats.NextGC,
		"gc_last":  m.stats.LastGC,
		"gc_num":   m.stats.NumGC,
		"gc_time":  m.stats.GCCPUFraction,
	}
}

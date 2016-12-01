package gohealth

import (
	"encoding/json"
	"net/http"
)

func NewHandler(m Monitorer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		alarms := m.GetAlarms()
		status := m.GetStatus()

		if len(alarms) > 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		v := map[string]interface{}{
			"alarms": alarms,
			"status": status,
		}

		json.NewEncoder(w).Encode(v)
	}
}

package gohealth

import (
	"encoding/json"
	"fmt"
	"time"
)

// PrintSmartDigits print an alarm to stdout following the SmartDigits standard
// spec. Example:
//
// {
//	"time": "2016-11-11 15:07:07.854Z",
//	"type": "ALARM",
//	"payload": {
//		"monitor": "idle",
//		"severity": "OK",
//		"msg": "Last event received at 2016-11-11 15:07:07.854307"
//     }
// }
//
func PrintSmartDigits(a *Alarm) {
	v := map[string]interface{}{

		"time": a.Time.UTC().Format(time.RFC3339Nano),
		"type": "ALARM",
		"payload": map[string]interface{}{
			"monitor":  a.Name,
			"severity": a.Severity,
			"msg":      a.Msg,
		},
	}

	j, _ := json.Marshal(v) // NO ERROR VALIDATION :D

	fmt.Println(string(j))
}

package utility

import (
	"PillarsPhenomVFXWeb/pillarsLog"
	"encoding/json"
	"fmt"
	"net/http"
)

func OutputJsonLog(w http.ResponseWriter, ret int, reason string, i interface{}, logInfo string) {
	var str string
	if i != nil {
		rs, _ := json.Marshal(i)
		str = string(rs)
	}

	out := &FeedbackMessage{ret, reason, str}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(b))

	if ret != 0 && logInfo != "" {
		pillarsLog.PillarsLogger.Print(logInfo)
	}
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	var str string
	if i != nil {
		rs, _ := json.Marshal(i)
		str = string(rs)
	}

	out := &FeedbackMessage{ret, reason, str}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(b))
}

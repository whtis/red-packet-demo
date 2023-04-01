package utils

import (
	"encoding/json"
)

func Json2String(val interface{}) string {
	jByte, err := json.Marshal(val)
	if err != nil {
		Error("Json2String error")
		return ""
	}
	return string(jByte)
}

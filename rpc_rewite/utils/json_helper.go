package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func Json2String(val interface{}) string {
	jByte, err := json.Marshal(val)
	if err != nil {
		logrus.Error("Json2String error")
		return ""
	}
	return string(jByte)
}

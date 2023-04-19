package convutil

import (
	"encoding/json"
	"fmt"
)

func MapToString(m map[string]interface{}) (string, error) {
	r, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("marshal map to string failed. err: %s", err.Error())
	}
	return string(r), nil
}

func StringToMap(str string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return m, fmt.Errorf("unmarshal string to map failed. err: %s", err.Error())
	}

	return m, nil
}

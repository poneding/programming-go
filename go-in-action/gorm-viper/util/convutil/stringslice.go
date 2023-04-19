package convutil

import (
	"encoding/json"
	"fmt"
)

func SliceToString(sli []interface{}) (string, error) {
	r, err := json.Marshal(sli)
	if err != nil {
		return "", fmt.Errorf("marshal slice to string failed. err: %s", err.Error())
	}
	return string(r), nil
}

func StringToSlice(str string) ([]interface{}, error) {
	var m []interface{}
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return m, fmt.Errorf("unmarshal string to map failed. err: %s", err.Error())
	}

	return m, nil
}

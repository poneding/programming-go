package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// SweepString remove specials strings in text.
func SweepString(str, spec string, ashes ...string) string {
	return strings.Replace(str, spec, "", -1)
}

// ShortenString
func ShortenString(v string, n int) string {
	if len(v) <= n {
		return v
	}
	return v[:n]
}

// JsonMarshal serialize object to json string.
func JsonMarshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		fmt.Errorf("json marshal failed.\n")
		return ""
	}
	return string(bytes)
}

// JsonUnmarshal deserialize json string to object.
func JsonUnmarshal(str string, v interface{}) {
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		fmt.Errorf("json unmarshal failed.\n")
	}
}

func GetOsEnv(key string, def ...string) string {
	if v := os.Getenv(key); len(v) > 0 {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

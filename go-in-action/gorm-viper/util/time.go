package util

import "time"

func SecondsTimeVal(v *int64) time.Time {
	if v != nil {
		return time.Unix(*v/1000, 0)
	}
	return time.Time{}
}

func MillisecondsTimeVal(v int64) time.Time {
	return time.Unix(0, v*int64(time.Millisecond))
}

func UnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

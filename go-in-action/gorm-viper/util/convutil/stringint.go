package convutil

import "strconv"

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

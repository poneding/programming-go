package util

func IntIf(condition bool, v1, v2 int) int {
	if condition {
		return v1
	}
	return v2
}

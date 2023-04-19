package util

import (
	"runtime"
	"strings"
)

func CurrentFuncFullName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	return runtime.FuncForPC(pc[0]).Name()
}

func CurrentFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)

	fullName := runtime.FuncForPC(pc[0]).Name()
	strs := strings.Split(fullName, ".")
	return strs[len(strs)-1]
}

package orm

import "errors"

var (
	RecordNotExist error = errors.New("record not exist")
)

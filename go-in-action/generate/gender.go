// gender.go
package generate

//go:generate stringer -type=Gender
type Gender int8

const (
	Female Gender = iota
	Male
	Unknown
)

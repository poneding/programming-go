package util

import "time"

func StringOrDefault(v, def string) string {
	if len(v) == 0 {
		v = def
	}
	return v
}

func IntOrDefault(v, def int) int {
	if v == 0 {
		v = def
	}
	return v
}
func Uint8OrDefault(v, def uint8) uint8 {
	if v == 0 {
		v = def
	}
	return v
}
func Uint16OrDefault(v, def uint16) uint16 {
	if v == 0 {
		v = def
	}
	return v
}
func Uint32OrDefault(v, def uint32) uint32 {
	if v == 0 {
		v = def
	}
	return v
}

func Uint64OrDefault(v, def uint64) uint64 {
	if v == 0 {
		v = def
	}
	return v
}

func Int8OrDefault(v, def int8) int8 {
	if v == 0 {
		v = def
	}
	return v
}

func Int16OrDefault(v, def int16) int16 {
	if v == 0 {
		v = def
	}
	return v
}

func Int32OrDefault(v, def int32) int32 {
	if v == 0 {
		v = def
	}
	return v
}

func Int64OrDefault(v, def int64) int64 {
	if v == 0 {
		v = def
	}
	return v
}

func StringPtr(v string) *string {
	return &v
}

func StringVal(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func BoolPtr(v bool) *bool {
	return &v
}

func BoolVal(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

func TimePtr(v time.Time) *time.Time {
	return &v
}

func TimeVal(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

// Int
func IntPtr(v int) *int {
	return &v
}

func IntVal(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// Uint
func UintPtr(v uint) *uint {
	return &v
}

func UintVal(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// Int8
func Int8Ptr(v int8) *int8 {
	return &v
}

func Int8Val(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint8
func Uint8Ptr(v uint8) *uint8 {
	return &v
}

func Uint8Val(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

// Int16
func Int16Ptr(v int16) *int16 {
	return &v
}

func Int16Val(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint16
func Uint16Ptr(v uint16) *uint16 {
	return &v
}

func Uint16Val(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

// Int32
func Int32Ptr(v int32) *int32 {
	return &v
}

func Int32Val(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint32
func Uint32Ptr(v uint32) *uint32 {
	return &v
}

func Uint32Val(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64
func Int64Ptr(v int64) *int64 {
	return &v
}

func Int64Val(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint64
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

func Uint64Val(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32
func Float32Ptr(v float32) *float32 {
	return &v
}

func Float32Val(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64
func Float64Ptr(v float64) *float64 {
	return &v
}

func Float64Val(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

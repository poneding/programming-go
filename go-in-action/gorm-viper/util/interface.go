package util

import (
	"math/rand"
	"time"
)

// var seed int64

// func init() {
// 	seed = time.Now().UnixNano()
// }

func Shuffle(slice []interface{}) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(time.Second.Microseconds()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := rand.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

package util

import (
	"fmt"
	"math/rand"
	"time"

	"log"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GeneratePureUUID() string {
	return SweepString(GenerateUUID(), "-")
}

// Password generate random string
// len must greater then or equal to 8
func GenerateRandomString(length uint8) string {
	res := ""
	if length < 8 {
		log.Println("The password length must be greater than or equal to 8.")
		return res
	}
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*~-?_"

	rand.NewSource(time.Now().UnixNano())

	// rand.Seed(time.Now().UnixNano())
	// time.Sleep(1 * time.Millisecond) // Sleep 1 millisecond to make sure the  time.Now().UnixNano() is refreshed.

	lowerChar := string(chars[rand.Intn(26)])
	upperChar := string(chars[rand.Intn(26)+26])
	digitChar := string(chars[rand.Intn(10)+52])
	speChar := string(chars[rand.Intn(len(chars)-62)+62])

	preSelectedChars := []string{lowerChar, upperChar, digitChar, speChar}

	for i := 4; i < int(length); i++ {
		res += string(chars[rand.Intn(len(chars))])
	}
	for _, c := range preSelectedChars {
		insertIndex := rand.Intn(len(res))
		res = fmt.Sprintf("%s%s%s", res[:insertIndex], c, res[insertIndex:])
	}

	return res
}

func GeneratePureRandomString(length uint8) string {
	res := ""
	if length == 0 {
		return res
	}
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Millisecond) // Sleep 1 millisecond to make sure the  time.Now().UnixNano() is refreshed.

	for i := 0; i < int(length); i++ {
		res += string(chars[rand.Intn(len(chars))])
	}

	return res
}

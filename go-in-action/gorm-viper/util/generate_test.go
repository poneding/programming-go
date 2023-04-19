package util

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	t.Log(GenerateRandomString(16))
	t.Log(GenerateRandomString(16))
	t.Log(GenerateRandomString(10))
}

func TestGeneratePureRandomString(t *testing.T) {
	t.Log(GeneratePureRandomString(10))
	t.Log(GeneratePureRandomString(10))
	t.Log(GeneratePureRandomString(10))
}

func TestGenUUID(t *testing.T) {
	t.Log(GenerateUUID())
	t.Log(GeneratePureUUID())
}

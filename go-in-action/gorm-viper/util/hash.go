package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

type HashMode uint

const (
	_ HashMode = iota
	HashMd5
	HashSha1
	HashSha256
	HashSha512
)

func Hash(str string, mode HashMode) string {
	h := getHash(mode)
	h.Write([]byte(str))
	bytes := h.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}

func DoubleHash(str string, mode HashMode) string {
	h := getHash(mode)
	h.Write([]byte(str))
	bytes := h.Sum(nil)
	h.Reset()
	h.Write(bytes)
	bytes = h.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}

func getHash(mode HashMode) hash.Hash {
	var h hash.Hash
	switch mode {
	case HashMd5:
		h = md5.New()
	case HashSha1:
		h = sha1.New()
	case HashSha256:
		h = sha256.New()
	case HashSha512:
		h = sha512.New()
	default:
		panic("unsupported hash mode.")
	}
	return h
}

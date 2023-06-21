package main

import (
	"crypto/rand"
	"crypto/rsa"
)

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	prikey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return prikey, &prikey.PublicKey, nil
}

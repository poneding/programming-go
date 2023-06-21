package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func Sign(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.New()
	hash.Write(data)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash.Sum(nil))
}

func Verify(data []byte, sig []byte, publicKey *rsa.PublicKey) error {
	hash := sha256.New()
	hash.Write(data)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash.Sum(nil), sig)
}

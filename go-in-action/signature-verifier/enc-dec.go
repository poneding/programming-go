package main

import (
	"crypto/rand"
	"crypto/rsa"
)

func Encrypt(data []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

func Decrypt(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
}

package main

import (
	"crypto/rsa"
	"crypto/x509"
)

func GetPrivateKey(b []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(b)
}

func GetPublicKey(b []byte) (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(b)
}

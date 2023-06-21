package main

import (
	"bytes"
	"testing"
)

func TestEncDec(t *testing.T) {
	prikey, pubkey, err := GenerateKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	data := []byte("hello")
	enc, err := Encrypt(data, pubkey)
	if err != nil {
		t.Fatal(err)
	}

	dec, err := Decrypt(enc, prikey)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, dec) {
		t.Fatalf("decrypted data is not equal to original data")
	}

	t.Log("decrypted data is equal to original data")
}

func TestSignVerify(t *testing.T) {
	prikey, pubkey, err := GenerateKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	data := []byte("hello")
	sig, err := Sign(data, prikey)
	if err != nil {
		t.Fatal(err)
	}

	err = Verify(data, sig, pubkey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("signature verified")
}

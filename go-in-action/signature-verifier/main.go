package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

var sigKey = []byte("HelloWorld")

func main() {
	config := controllerruntime.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)

	s, err := clientset.CoreV1().Secrets("default").Get(context.Background(), "hello-sig-secret", metav1.GetOptions{})
	if err != nil {
		panic("get secret err: " + err.Error())
	}

	certData := s.Data["cert.crt"]
	sig := s.Data["hello.sig"]
	if certData == nil || sig == nil {
		panic("cert.crt or hello.sig is nil")
	}

	b, _ := pem.Decode(certData)

	// 验证证书有效期
	cert, err := x509.ParseCertificate(b.Bytes)
	if err != nil {
		panic("parse cert err: " + err.Error())
	}

	if cert.NotAfter.Before(time.Now()) || cert.NotBefore.After(time.Now()) {
		panic("cert is expired")
	}

	pubkey := cert.PublicKey.(*rsa.PublicKey)

	// 公钥验证签名
	err = Verify(sigKey, sig, pubkey)
	if err != nil {
		panic("verify err: " + err.Error())
	}
	log.Println("signature verified")
}

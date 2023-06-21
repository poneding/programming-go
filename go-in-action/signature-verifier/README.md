# Golang 签名验证

1. 使用 openssl 升成私钥和公钥

```bash
openssl genrsa -out pri.key 2048
openssl rsa -in pri.key -pubout -out pub.key
```

2. 使用私钥生成证书

```bash
openssl req -new -x509 -key pri.key -out cert.crt -days 365 -subj /C=CN/ST=Hunan/L=Changsha/O=/OU=/CN=poneding.com
```

3. 对文件签名

```bash
echo -n "HelloWorld" > hello
openssl dgst -sha256 -sign pri.key -out hello.sig hello
```

4. 使用公钥验证签名

```bash
openssl dgst -sha256 -verify pub.key -signature hello.sig hello
```

5. 使用公钥和签名文件创建 Kubernetes Secret

```bash
kubectl create secret generic hello-sign-secret --from-file=cert.crt=cert.crt --from-file=hello.sig=hello.sig
```

6. Go 代码验证

```go
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
```

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/golang/glog"
)

const (
	port = 10000
)

var (
	demoKeyPair  *tls.Certificate
	demoCertPool *x509.CertPool
	demoAddr     string
)

func init() {
	var err error
	pair, err := tls.X509KeyPair([]byte(Cert), []byte(Key))
	if err != nil {
		glog.Fatal(err)
	}
	demoKeyPair = &pair
	demoCertPool = x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM([]byte(Cert))
	if !ok {
		glog.Fatal("bad certs")
	}
	demoAddr = fmt.Sprintf("localhost:%d", port)
}

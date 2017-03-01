package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
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
		log.Fatal(err.Error())
	}
	demoKeyPair = &pair
	demoCertPool = x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM([]byte(Cert))
	if !ok {
		log.Fatal("Certificado inválido")
	}
	demoAddr = fmt.Sprintf("localhost")
}

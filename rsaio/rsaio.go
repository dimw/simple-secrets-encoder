package rsaio

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func LoadPublicKey(filename string) (*rsa.PublicKey, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	p, _ := pem.Decode(data)

	key, err := x509.ParsePKCS1PublicKey(p.Bytes)

	return key, err
}

func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	p, _ := pem.Decode(data)

	if x509.IsEncryptedPEMBlock(p) {
		return nil, fmt.Errorf("encrypted PEM blocks are not supported")
	}

	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)

	return key, err
}

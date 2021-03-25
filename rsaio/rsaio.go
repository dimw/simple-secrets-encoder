package rsaio

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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

	// TODO use passkey if needed

	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)

	return key, err
}

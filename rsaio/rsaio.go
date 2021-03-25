package rsaio

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

var errEncryptedPEMBlockNotSupported = errors.New("encrypted PEM blocks are not supported")

func EncryptedPEMBlockNotSupportedError() error {
	return fmt.Errorf("%w", errEncryptedPEMBlockNotSupported)
}

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
		return nil, EncryptedPEMBlockNotSupportedError()
	}

	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)

	return key, err
}

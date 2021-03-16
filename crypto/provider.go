package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"regexp"
)

type ProcessFunc func(string) (string, error)

type Provider struct {
	secretKeysRegexp            *regexp.Regexp
	encryptedValueRegexp        *regexp.Regexp
	encryptedValueParsingRegexp *regexp.Regexp
	publicKey                   *rsa.PublicKey
	privateKey                  *rsa.PrivateKey
	processFunc                 ProcessFunc
	Strategy                    string
}

func (p *Provider) IsEncrypted(val string) bool {
	return p.encryptedValueRegexp.MatchString(val)
}

func (p *Provider) Encrypt(val string) (string, error) {
	if !p.IsEncrypted(val) {
		encVal, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, p.publicKey, []byte(val), []byte(""))
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("ENC[rsa,data:%v]", base64.StdEncoding.EncodeToString(encVal)), nil
	}

	return val, nil
}

func (p *Provider) Decrypt(val string) (string, error) {
	if p.IsEncrypted(val) {
		subMatches := p.encryptedValueParsingRegexp.FindStringSubmatch(val)
		methodSubmatchIndex := p.encryptedValueParsingRegexp.SubexpIndex("method")
		method := subMatches[methodSubmatchIndex]

		if method != "rsa" {
			return "", fmt.Errorf(`RSA encryption method is supported only, found "%v"`, method)
		}

		dataSubmatchIndex := p.encryptedValueParsingRegexp.SubexpIndex("data")
		dataEnc := subMatches[dataSubmatchIndex]
		data, err := base64.StdEncoding.DecodeString(dataEnc)
		if err != nil {
			return "", fmt.Errorf(`data cannot be encoded from Base64`)
		}

		decodedVal, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, p.privateKey, data, []byte(""))
		if err != nil {
			return "", err
		}

		return string(decodedVal), nil
	}

	return val, nil
}

func (p *Provider) IsSecretKey(key string) bool {
	return p.secretKeysRegexp.MatchString(key)
}

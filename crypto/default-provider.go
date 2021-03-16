package crypto

import (
	"crypto/rsa"
	"regexp"
)

const (
	StrategyEncrypt = "encrypt"
	StrategyDecrypt = "decrypt"
)

var secretKeysRegexp = regexp.MustCompile("(?i)(secret|token|password)$")
var encryptedValueRegexp = regexp.MustCompile("(?i)^ENC\\[.*]$")
var encryptedValueParsingRegexp = regexp.MustCompile("^ENC\\[(?P<method>\\w+),data:(?P<data>.*)]$")

func CreateEncryptionProvider(publicKey *rsa.PublicKey) *Provider {
	return &Provider{
		secretKeysRegexp:     secretKeysRegexp,
		encryptedValueRegexp: encryptedValueRegexp,
		publicKey:            publicKey,
		Strategy:             StrategyEncrypt,
	}
}

func CreateDecryptionProvider(privateKey *rsa.PrivateKey) *Provider {
	return &Provider{
		secretKeysRegexp:            secretKeysRegexp,
		encryptedValueRegexp:        encryptedValueRegexp,
		encryptedValueParsingRegexp: encryptedValueParsingRegexp,
		privateKey:                  privateKey,
		Strategy:                    StrategyDecrypt,
	}
}

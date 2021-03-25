package process

import (
	"errors"
	"fmt"
	"log"

	"github.com/dimw/simple-secrets-encryptor/crypto"
)

var errUnsupportedKeyType = errors.New("unsupportedKeyTypeError")

func UnsupportedKeyTypeError(keyType interface{}) error {
	return fmt.Errorf(`%w: "%v"`, errUnsupportedKeyType, keyType)
}

var errUnknownStrategy = errors.New("unknownStrategyError")

func UnknownStrategyError(strategy string) error {
	return fmt.Errorf(`%w: "%v"`, errUnknownStrategy, strategy)
}

func Walk(data map[string]interface{}, encryptionProvider *crypto.Provider) (map[string]interface{}, error) {
	processedData := make(map[string]interface{})

	for key, val := range data {
		if encryptionProvider.IsSecretKey(key) || encryptionProvider.Strategy == "decrypt" {
			var err error
			switch t := val.(type) {
			case string:
				log.Printf("↳ Processing key: %v", key)
				switch encryptionProvider.Strategy {
				case "encrypt":
					processedData[key], err = encryptionProvider.Encrypt(val.(string))
				case "decrypt":
					processedData[key], err = encryptionProvider.Decrypt(val.(string))
				default:
					return nil, UnknownStrategyError(encryptionProvider.Strategy)
				}
			default:
				return nil, UnsupportedKeyTypeError(t)
			}

			if err != nil {
				return nil, err
			}
		} else {
			processedData[key] = val
		}
	}

	return processedData, nil
}

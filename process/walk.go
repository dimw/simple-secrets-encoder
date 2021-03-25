package process

import (
	"fmt"
	"log"

	"github.com/dimw/simple-secrets-encryptor/crypto"
)

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
					return nil, fmt.Errorf(`unknown strategy "%v"`, encryptionProvider.Strategy)
				}
			default:
				return nil, fmt.Errorf(`unsupported type "%v"`, t)
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

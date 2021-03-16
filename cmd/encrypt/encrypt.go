package encrypt

import (
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/glob"
	"github.com/dimw/simple-secrets-encryptor/io"
	"github.com/dimw/simple-secrets-encryptor/rsaio"
	"github.com/dimw/simple-secrets-encryptor/secrets"
	"log"
)

type Args struct {
	PublicKeyFilename string
	Workdir           string
	FilenamePattern   string
}

func Encrypt(args Args) error {
	publicKey, err := rsaio.LoadPublicKey(args.PublicKeyFilename)
	if err != nil {
		return err
	}

	files, err := glob.Glob(args.Workdir, args.FilenamePattern)

	for _, filename := range files {
		log.Printf(`Encoding: %v`, filename)
		data, err := io.ReadYaml(filename)

		if err != nil {
			log.Fatal("Error loading", filename, err)
		}

		encryptedData, err := secrets.ProcessSecrets(data, crypto.CreateEncryptionProvider(publicKey))
		if err != nil {
			return err
		}

		err = io.WriteYaml(filename, encryptedData)
		if err != nil {
			return err
		}
	}

	return err
}

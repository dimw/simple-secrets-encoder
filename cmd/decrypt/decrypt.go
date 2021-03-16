package decrypt

import (
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/glob"
	"github.com/dimw/simple-secrets-encryptor/io"
	"github.com/dimw/simple-secrets-encryptor/rsaio"
	"github.com/dimw/simple-secrets-encryptor/secrets"
	"log"
)

type Args struct {
	PrivateKeyFilename string
	Workdir            string
	FilenamePattern    string
}

func Decrypt(args Args) error {
	privateKey, err := rsaio.LoadPrivateKey(args.PrivateKeyFilename)
	if err != nil {
		return err
	}

	files, err := glob.Glob(args.Workdir, args.FilenamePattern)
	if err != nil {
		return err
	}

	for _, filename := range files {
		log.Printf(`Decoding: %v`, filename)
		data, err := io.ReadYaml(filename)

		if err != nil {
			log.Fatalf("Error loading %v", filename)
		}

		decryptedData, err := secrets.ProcessSecrets(data, crypto.CreateDecryptionProvider(privateKey))
		if err != nil {
			return err
		}

		err = io.WriteYaml(filename, decryptedData)
		if err != nil {
			return err
		}
	}

	return err
}

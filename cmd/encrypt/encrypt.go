package encrypt

import (
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/fileutils"
	"github.com/dimw/simple-secrets-encryptor/rsaio"
)

type Args struct {
	PublicKeyFilename string
	Workdir           string
	FilenamePattern   string
	Outdir            string
}

func Encrypt(args Args) error {
	publicKey, err := rsaio.LoadPublicKey(args.PublicKeyFilename)
	if err != nil {
		return err
	}

	provider := crypto.CreateEncryptionProvider(publicKey)

	err = fileutils.IterateFiles(args.Workdir, args.FilenamePattern, args.Outdir, provider)
	if err != nil {
		return err
	}

	return err
}

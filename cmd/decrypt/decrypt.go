package decrypt

import (
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/fileutils"
	"github.com/dimw/simple-secrets-encryptor/rsaio"
)

type Args struct {
	PrivateKeyFilename string
	Workdir            string
	FilenamePattern    string
	Outdir             string
}

func Decrypt(args Args) error {
	privateKey, err := rsaio.LoadPrivateKey(args.PrivateKeyFilename)
	if err != nil {
		return err
	}

	provider := crypto.CreateDecryptionProvider(privateKey)

	err = fileutils.IterateFiles(args.Workdir, args.FilenamePattern, args.Outdir, provider)
	if err != nil {
		return err
	}

	return nil
}

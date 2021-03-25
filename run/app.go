package run

import (
	"github.com/urfave/cli/v2"
)

const (
	flagOutdir          = "outdir"
	flagWorkdir         = "workdir"
	flagFilenamePattern = "filename-pattern"
	flagPublicKeyFile   = "public-key-file"
	flagPrivateKeyFile  = "private-key-file"
	flagKeySize         = "key-size"
	flagReplaceKeys     = "replace-keys"
	flagOutputFormat    = "output-format"
)

func CreateApp() *cli.App {
	app := &cli.App{
		Name:  "Simple Secret Encryptor",
		Usage: "Tool for asymmetric encryption (RSA) of structured secret files",

		Commands: []*cli.Command{
			NewEncryptCommand(),
			NewDecryptCommand(),
			NewGenerateKeysCommand(),
		},
	}

	return app
}

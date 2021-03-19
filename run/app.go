package run

import (
	"github.com/urfave/cli/v2"
)

const flagOutdir = "outdir"
const flagWorkdir = "workdir"
const flagFilenamePattern = "filename-pattern"
const flagPublicKeyFile = "public-key-file"
const flagPrivateKeyFile = "private-key-file"
const flagKeySize = "key-size"
const flagReplaceKeys = "replace-keys"
const flagOutputFormat = "output-format"

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

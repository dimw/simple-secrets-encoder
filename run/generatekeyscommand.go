package run

import (
	generate_keys "github.com/dimw/simple-secrets-encryptor/cmd/generate-keys"
	"github.com/urfave/cli/v2"
)

func NewGenerateKeysCommand() *cli.Command {
	return &cli.Command{
		Name: "generate-keys",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  flagPrivateKeyFile,
				Value: "private.key",
				Usage: "Location of the private key",
			},
			&cli.StringFlag{
				Name:  flagPublicKeyFile,
				Value: "public.pem",
				Usage: "Location of the public key",
			},
			&cli.IntFlag{
				Name:  flagKeySize,
				Value: 2048,
				Usage: "Key size",
			},
			&cli.BoolFlag{
				Name:  flagReplaceKeys,
				Usage: "Replace already existing private and public key files",
			},
		},
		Aliases: []string{"g"},
		Usage:   "Generate a pair of RSA keys",
		Action:  generateKeysAction,
	}
}

func generateKeysAction(c *cli.Context) error {
	args := generate_keys.GenerateRSAArgs{
		PrivateKeyFilename: c.String(flagPrivateKeyFile),
		PublicKeyFilename:  c.String(flagPublicKeyFile),
		KeySize:            c.Int(flagKeySize),
		ReplaceKeys:        c.Bool(flagReplaceKeys),
	}

	return generate_keys.GenerateRSA(args)
}

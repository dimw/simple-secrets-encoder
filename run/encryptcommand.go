package run

import (
	"github.com/dimw/simple-secrets-encryptor/cmd/encrypt"
	"github.com/urfave/cli/v2"
)

func NewEncryptCommand() *cli.Command {
	return &cli.Command{
		Name: "encrypt",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  flagPublicKeyFile,
				Value: "public.pem",
				Usage: "Location of the public key",
			},
			&cli.StringFlag{
				Name:  flagWorkdir,
				Value: "./",
				Usage: "Location of files to encrypt",
			},
			&cli.StringFlag{
				Name:  flagFilenamePattern,
				Value: "**/*.{yml,yaml,json}",
				Usage: "Glob pattern of files to include",
			},
			&cli.StringFlag{
				Name:  flagOutdir,
				Value: "",
				Usage: "Output location for files (workdir)",
			},
			&cli.StringFlag{
				Name:  flagOutputFormat,
				Value: "",
				Usage: "Output format: yaml, json (default: same as input file)",
			},
		},
		Aliases: []string{"e"},
		Usage:   "Encrypt secrets",
		Action:  encryptAction,
	}
}

func encryptAction(c *cli.Context) error {
	args := encrypt.Args{
		PublicKeyFilename: c.String(flagPublicKeyFile),
		Workdir:           c.String(flagWorkdir),
		FilenamePattern:   c.String(flagFilenamePattern),
		Outdir:            c.String(flagOutdir),
		OutputFormat:      c.String(flagOutputFormat),
	}

	return encrypt.Encrypt(args)
}

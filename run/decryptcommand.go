package run

import (
	"github.com/dimw/simple-secrets-encryptor/cmd/decrypt"
	"github.com/urfave/cli/v2"
)

func NewDecryptCommand() *cli.Command {
	return &cli.Command{
		Name: "decrypt",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  flagPrivateKeyFile,
				Value: "private.key",
				Usage: "Location of the private key",
			},
			&cli.StringFlag{
				Name:  flagWorkdir,
				Value: "./",
				Usage: "Location files to decrypt",
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
		Aliases: []string{"d"},
		Usage:   "Decrypt secrets",
		Action:  decryptAction,
	}
}

func decryptAction(c *cli.Context) error {
	args := decrypt.Args{
		PrivateKeyFilename: c.String(flagPrivateKeyFile),
		Workdir:            c.String(flagWorkdir),
		FilenamePattern:    c.String(flagFilenamePattern),
		Outdir:             c.String(flagOutdir),
		OutputFormat:       c.String(flagOutputFormat),
	}

	return decrypt.Decrypt(args)
}

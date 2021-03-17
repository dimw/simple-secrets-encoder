package main

import (
	"github.com/dimw/simple-secrets-encryptor/cmd/decrypt"
	"github.com/dimw/simple-secrets-encryptor/cmd/encrypt"
	generate_keys "github.com/dimw/simple-secrets-encryptor/cmd/generate-keys"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	const flagOutdir = "outdir"
	const flagWorkdir = "workdir"
	const flagFilenamePattern = "filename-pattern"
	const flagPublicKeyFile = "public-key-file"
	const flagPrivateKeyFile = "private-key-file"
	const flagKeySize = "key-size"
	const flagReplaceKeys = "replace-keys"
	const flagOutputFormat = "output-format"

	app := &cli.App{
		Name:  "Simple Secret Encryptor",
		Usage: "Tool for asymmetric encryption (RSA) of secrets",

		Commands: []*cli.Command{
			{
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
				Action: func(c *cli.Context) error {
					args := encrypt.Args{
						PublicKeyFilename: c.String(flagPublicKeyFile),
						Workdir:           c.String(flagWorkdir),
						FilenamePattern:   c.String(flagFilenamePattern),
						Outdir:            c.String(flagOutdir),
						OutputFormat:      c.String(flagOutputFormat),
					}
					return encrypt.Encrypt(args)
				},
			},

			{
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
						Value: "**/*.{yml,yaml,json}.enc",
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
				Action: func(c *cli.Context) error {
					args := decrypt.Args{
						PrivateKeyFilename: c.String(flagPrivateKeyFile),
						Workdir:            c.String(flagWorkdir),
						FilenamePattern:    c.String(flagFilenamePattern),
						Outdir:             c.String(flagOutdir),
						OutputFormat:       c.String(flagOutputFormat),
					}
					return decrypt.Decrypt(args)
				},
			},

			{
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
				Action: func(c *cli.Context) error {
					args := generate_keys.GenerateRSAArgs{
						PrivateKeyFilename: c.String(flagPrivateKeyFile),
						PublicKeyFilename:  c.String(flagPublicKeyFile),
						KeySize:            c.Int(flagKeySize),
						ReplaceKeys:        c.Bool(flagReplaceKeys),
					}
					return generate_keys.GenerateRSA(args)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

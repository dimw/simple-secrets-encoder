package main

import (
	"github.com/dimw/simple-secrets-encryptor/cmd"
	"github.com/dimw/simple-secrets-encryptor/cmd/decrypt"
	"github.com/dimw/simple-secrets-encryptor/cmd/encrypt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Simple Secret Encryptor",
		Usage: "Tool for asymmetric encryption (RSA) of secrets",

		Commands: []*cli.Command{
			{
				Name: "encrypt",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "public-key-file",
						Value: "public.pem",
						Usage: "Location of the public key",
					},
					&cli.StringFlag{
						Name:  "workdir",
						Value: "./",
						Usage: "Location of YAML files to encode",
					},
					&cli.StringFlag{
						Name:  "filename-pattern",
						Value: "**/*.{yml,yaml}",
						Usage: "Glob pattern of files to include",
					},
				},
				Aliases: []string{"e"},
				Usage:   "Encrypt secrets",
				Action: func(c *cli.Context) error {
					args := encrypt.Args{
						PublicKeyFilename: c.String("public-key-file"),
						Workdir:           c.String("workdir"),
						FilenamePattern:   c.String("filename-pattern"),
					}
					return encrypt.Encrypt(args)
				},
			},

			{
				Name: "decrypt",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "private-key-file",
						Value: "private.key",
						Usage: "Location of the private key",
					},
					&cli.StringFlag{
						Name:  "workdir",
						Value: "./",
						Usage: "Location of YAML files to decode",
					},
					&cli.StringFlag{
						Name:  "filename-pattern",
						Value: "**/*.{yml,yaml}.enc",
						Usage: "Glob pattern of files to include",
					},
				},
				Aliases: []string{"d"},
				Usage:   "Decrypt secrets",
				Action: func(c *cli.Context) error {
					args := decrypt.Args{
						PrivateKeyFilename: c.String("private-key-file"),
						Workdir:            c.String("workdir"),
						FilenamePattern:    c.String("filename-pattern"),
					}
					return decrypt.Decrypt(args)
				},
			},

			{
				Name: "generate-keys",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "private-key-file",
						Value: "private.key",
						Usage: "Location of the private key",
					},
					&cli.StringFlag{
						Name:  "public-key-file",
						Value: "public.pem",
						Usage: "Location of the public key",
					},
					&cli.IntFlag{
						Name:  "bit-size",
						Value: 2048,
						Usage: "Key size",
					},
					&cli.BoolFlag{
						Name:  "replace-keys",
						Usage: "Replace already existing private and public key files",
					},
				},
				Aliases: []string{"g"},
				Usage:   "Generate a pair of RSA keys",
				Action: func(c *cli.Context) error {
					args := cmd.GenerateRSAArgs{
						PrivateKeyFilename: c.String("private-key-file"),
						PublicKeyFilename:  c.String("public-key-file"),
						KeySize:            c.Int("key-size"),
						ReplaceKeys:        c.Bool("replace-keys"),
					}
					return cmd.GenerateRSA(args)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

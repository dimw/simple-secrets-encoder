package main

import (
	"log"
	"os"

	"github.com/dimw/simple-secrets-encryptor/run"
)

func main() {
	app := run.CreateApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

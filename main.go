package main

import (
	"github.com/dimw/simple-secrets-encryptor/run"
	"log"
	"os"
)

func main() {
	app := run.CreateApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

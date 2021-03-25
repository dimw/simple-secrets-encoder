package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dimw/simple-secrets-encryptor/run"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := run.CreateApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

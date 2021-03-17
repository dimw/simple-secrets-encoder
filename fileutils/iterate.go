package fileutils

import (
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/io"
	"github.com/dimw/simple-secrets-encryptor/process"
	"log"
)

func IterateFiles(workdir string, filenamePattern string, provider *crypto.Provider) error {
	files, err := Glob(workdir, filenamePattern)
	if err != nil {
		return err
	}

	for _, filename := range files {
		log.Printf(`Decoding: %v`, filename)
		data, err := io.ReadYaml(filename)

		if err != nil {
			log.Fatalf("Error loading %v", filename)
		}

		encodedData, err := process.Walk(data, provider)
		if err != nil {
			return err
		}

		err = io.WriteYaml(filename, encodedData)
		if err != nil {
			return err
		}
	}
	return nil
}

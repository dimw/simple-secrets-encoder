package fileutils

import (
	"fmt"
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/io"
	"github.com/dimw/simple-secrets-encryptor/process"
	"log"
	"os"
	"path/filepath"
)

func IterateFiles(workdir string, filenamePattern string, outdir string, provider *crypto.Provider) error {
	files, err := Glob(workdir, filenamePattern)
	if err != nil {
		return err
	}

	validOutdir := outdir
	if validOutdir == "" {
		validOutdir = workdir
	} else {
		_ = os.Mkdir(outdir, os.ModeDir)
	}

	for _, filename := range files {
		log.Printf(`Reading: %v`, filename)
		data, err := io.Read(filename)
		if err != nil {
			return fmt.Errorf("error reading file: %v", filename)
		}

		encodedData, err := process.Walk(data, provider)
		if err != nil {
			return err
		}

		outputFilename := filepath.ToSlash(validOutdir + "/" + filename[len(workdir)+1:])
		err = io.Write(outputFilename, encodedData)
		if err != nil {
			return err
		}
	}
	return nil
}

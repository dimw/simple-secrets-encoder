package fileutils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/io"
	"github.com/dimw/simple-secrets-encryptor/process"
)

var errReadFile = errors.New("readFileError")

func ReadFileError(filename string) error {
	return fmt.Errorf(`%w: %v`, errReadFile, filename)
}

func IterateFiles(workdir string, filenamePattern string, outdir string, format string, provider *crypto.Provider) error {
	workdir = filepath.Clean(workdir)
	files, err := Glob(workdir, filenamePattern)
	if err != nil {
		return err
	}

	if outdir == "" {
		outdir = workdir
	} else {
		outdir = filepath.Clean(outdir)
		_ = os.Mkdir(outdir, 0o770)
	}

	for _, filename := range files {
		log.Printf(`Reading: %v`, filename)
		data, err := io.Read(filename)
		if err != nil {
			return ReadFileError(filename)
		}

		encodedData, err := process.Walk(data, provider)
		if err != nil {
			return err
		}

		outputFilename := updateOutdir(workdir, outdir, filename)
		outputFilename = updateFilename(outputFilename, format)
		err = io.Write(outputFilename, encodedData)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateOutdir(workdir string, outdir string, filename string) string {
	return filepath.ToSlash(outdir + "/" + filename[len(workdir)+1:])
}

func updateFilename(filename string, format string) string {
	if format != "" {
		filename = filename[0:len(filename)-len(filepath.Ext(filename))] + "." + format
	}

	return filename
}

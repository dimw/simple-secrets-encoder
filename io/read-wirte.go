package io

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var errUnsupportedFile = errors.New("unsupportedFileError")

func UnsupportedFileError(filename string) error {
	return fmt.Errorf(`%w: %v`, errUnsupportedFile, filename)
}

var errParsing = errors.New("parsingError")

func ParsingError(filename string) error {
	return fmt.Errorf(`%w: %v`, errParsing, filename)
}

var errFolderCreation = errors.New("folderCreationError")

func FolderCreationError(path string) error {
	return fmt.Errorf(`%w: %v`, errFolderCreation, path)
}

var errFileCreation = errors.New("fileCreationError")

func FileCreationError(path string) error {
	return fmt.Errorf(`%w: %v`, errFileCreation, path)
}

func Read(filename string) (map[string]interface{}, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})

	switch filepath.Ext(filename) {
	case ".yml", ".yaml":
		err = yaml.Unmarshal(fileData, &data)
	case ".json":
		err = json.Unmarshal(fileData, &data)
	default:
		err = UnsupportedFileError(filename)
	}

	return data, err
}

func Write(filename string, data map[string]interface{}) error {
	var err error
	var bytes []byte
	switch filepath.Ext(filename) {
	case ".yml", ".yaml":
		bytes, err = yaml.Marshal(data)
	case ".json":
		bytes, err = json.MarshalIndent(data, "", "  ")
	default:
		return UnsupportedFileError(filename)
	}

	if err != nil {
		return ParsingError(filename)
	}

	outputDir := filepath.Dir(filename)
	err = os.MkdirAll(outputDir, 0o770)
	if err != nil {
		return FolderCreationError(outputDir)
	}

	err = ioutil.WriteFile(filename, bytes, 0o644)
	if err != nil {
		return FileCreationError(filename)
	}

	return nil
}

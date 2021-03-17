package io

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func Read(filename string) (map[string]interface{}, error) {
	fileData, err := ioutil.ReadFile(filename)
	data := make(map[string]interface{})

	switch filepath.Ext(filename) {
	case ".yml", ".yaml":
		err = yaml.Unmarshal(fileData, &data)
	default:
		err = fmt.Errorf("unsupported file extension: %v", filename)
	}

	return data, err
}

func Write(filename string, data map[string]interface{}) error {
	var err error
	var bytes []byte
	switch filepath.Ext(filename) {
	case ".yml", ".yaml":
		bytes, err = yaml.Marshal(data)
	default:
		return fmt.Errorf("unsupported file extension: %v", filename)
	}

	if err != nil {
		return fmt.Errorf("could not marshal data")
	}

	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return fmt.Errorf("cannot write to file: %v", filename)
	}

	return nil
}

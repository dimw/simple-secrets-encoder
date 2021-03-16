package io

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ReadYaml(filename string) (map[string]interface{}, error) {
	fileData, err := ioutil.ReadFile(filename)
	data := make(map[string]interface{})

	err = yaml.Unmarshal(fileData, &data)

	return data, err
}

func WriteYaml(filename string, data map[string]interface{}) error {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("Could not marshal data %v!", data)
	}

	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

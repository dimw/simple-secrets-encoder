package io

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestReadYaml(t *testing.T) {
	tmpYmlFile, _ := ioutil.TempFile("./", "foo.*.yml")
	tmpYmlFile.WriteString("foo: bar")
	tmpYmlFile.Close()

	yaml, err := ReadYaml(tmpYmlFile.Name())

	assert.Nil(t, err)
	assert.Equal(t, "bar", yaml["foo"])

	_ = os.Remove(tmpYmlFile.Name())
}

func TestWriteYaml(t *testing.T) {
	tmpYmlFile, _ := ioutil.TempFile("./", "foo.*.yml")
	tmpYmlFile.Close()

	data := make(map[string]interface{})
	data["foo"] = "bar"

	err := WriteYaml(tmpYmlFile.Name(), data)
	assert.Nil(t, err)

	content, err := ioutil.ReadFile(tmpYmlFile.Name())

	assert.Equal(t, "foo: bar", strings.TrimSpace(string(content)))
	_ = os.Remove(tmpYmlFile.Name())
}

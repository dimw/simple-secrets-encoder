package io

import (
	"github.com/dimw/simple-secrets-encryptor/testhelper/tempfile"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestReadYaml(t *testing.T) {
	tmpYmlFile := tempfile.New("./", "foo.*.yml", "foo: bar")
	defer tmpYmlFile.Remove()

	yaml, err := ReadYaml(tmpYmlFile.Name)

	assert.NoError(t, err)
	assert.Equal(t, "bar", yaml["foo"])
}

func TestWriteYaml(t *testing.T) {
	tmpYmlFile := tempfile.New("./", "foo.*.yml", "")
	defer tmpYmlFile.Remove()

	data := make(map[string]interface{})
	data["foo"] = "bar"

	err := WriteYaml(tmpYmlFile.Name, data)
	assert.NoError(t, err)

	content, err := ioutil.ReadFile(tmpYmlFile.Name)

	assert.Equal(t, "foo: bar", strings.TrimSpace(string(content)))
	_ = os.Remove(tmpYmlFile.Name)
}

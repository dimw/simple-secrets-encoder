package io

import (
	"github.com/dimw/simple-secrets-encryptor/test/tempfile"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestShouldNotReadUnsupportedFileFormat(t *testing.T) {
	tmpFile := tempfile.Create("./", "foo.*.txt", "")
	defer tmpFile.Remove()

	_, err := Read(tmpFile.Name)

	assert.NotNil(t, err)
}

func TestShouldReadFileFormat(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		content  string
	}{
		{name: "Should read .yaml", filename: "foo.*.yaml", content: "foo: bar"},
		{name: "Should read .yml", filename: "foo.*.yml", content: "foo: bar"},
		{name: "Should read .json", filename: "foo.*.json", content: `{"foo": "bar"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := tempfile.Create("./", tt.filename, tt.content)
			defer tmpFile.Remove()

			data, err := Read(tmpFile.Name)

			assert.NotNil(t, data)
			assert.Nil(t, err)
		})
	}
}

func TestShouldNotWriteUnsupportedFileFormat(t *testing.T) {
	tmpFile := tempfile.Create("./", "foo.*.txt", "boo")
	defer tmpFile.Remove()

	data := make(map[string]interface{})
	data["foo"] = "bar"

	err := Write(tmpFile.Name, data)
	assert.NotNil(t, err)
}

func TestShouldWriteFileFormat(t *testing.T) {
	tests := []struct {
		name            string
		filename        string
		expectedContent string
	}{
		{name: "Should write to .yaml", filename: "foo.*.yaml", expectedContent: "foo: bar"},
		{name: "Should write to .yml", filename: "foo.*.yml", expectedContent: "foo: bar"},
		{name: "Should write to .json", filename: "foo.*.json", expectedContent: `{"foo":"bar"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := tempfile.Create("./", tt.filename, "")
			defer tmpFile.Remove()

			data := make(map[string]interface{})
			data["foo"] = "bar"

			err := Write(tmpFile.Name, data)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedContent, tmpFile.Content())
		})
	}
}

func TestShouldCreateSubFolders(t *testing.T) {
	tmpFile := tempfile.Create("./", "foo.*.yaml", "foo: bar")
	defer tmpFile.Remove()

	data := make(map[string]interface{})
	data["foo"] = "bar"

	filename := "booh/" + tmpFile.Name
	err := Write(filename, data)
	assert.Nil(t, err)

	contentBytes, err := ioutil.ReadFile(filename)
	assert.Nil(t, err)
	assert.Equal(t, "foo: bar", strings.TrimSpace(string(contentBytes)))
}

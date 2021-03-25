package io

import (
	"github.com/dimw/simple-secrets-encryptor/testhelper/ossafe"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/dimw/simple-secrets-encryptor/testhelper/tempfile"
	"github.com/stretchr/testify/assert"
)

func TestShouldNotReadUnsupportedFileFormat(t *testing.T) {
	tmpFile := tempfile.NewT(t, "./", "foo.*.txt", "")

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
			tmpFile := tempfile.NewT(t, "./", tt.filename, tt.content)

			data, err := Read(tmpFile.Name)

			assert.NotNil(t, data)
			assert.NoError(t, err)
		})
	}
}

func TestShouldNotWriteUnsupportedFileFormat(t *testing.T) {
	tmpFile := tempfile.NewT(t, "./", "foo.*.txt", "boo")

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
		{name: "Should write to .json", filename: "foo.*.json", expectedContent: "{\n  \"foo\": \"bar\"\n}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := tempfile.NewT(t, "./", tt.filename, "")

			data := make(map[string]interface{})
			data["foo"] = "bar"

			err := Write(tmpFile.Name, data)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedContent, tmpFile.Content())
		})
	}
}

func TestShouldCreateSubFolders(t *testing.T) {
	tmpFile := tempfile.NewT(t, "./", "foo.*.yaml", "foo: bar")

	data := make(map[string]interface{})
	data["foo"] = "bar"

	filename := "booh/" + tmpFile.Name
	defer ossafe.Remove(filename)
	err := Write(filename, data)
	assert.NoError(t, err)

	contentBytes, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, "foo: bar", strings.TrimSpace(string(contentBytes)))
}

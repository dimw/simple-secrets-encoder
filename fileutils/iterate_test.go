package fileutils

import (
	crypto_rand "crypto/rand"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/dimw/simple-secrets-encryptor/testhelper/ossafe"

	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/testhelper/tempfile"
	"github.com/stretchr/testify/assert"
)

func TestShouldIterate(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	tmpDir2, _ := ioutil.TempDir(tmpDir, "tmp2-*")
	defer ossafe.RemoveAll(tmpDir)
	_ = tempfile.NewT(t, tmpDir2, "foo.*.yml", "")

	err := IterateFiles(tmpDir, "*.yml", "", "", nil)
	assert.NoError(t, err)
}

func TestShouldOutputToOutdir(t *testing.T) {
	// create folder with secrets
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer ossafe.RemoveAll(tmpDir)
	tmpDir2, _ := ioutil.TempDir(tmpDir, "subdir-*")
	defer ossafe.RemoveAll(tmpDir2)
	tmpFile := tempfile.NewT(t, tmpDir2, "foo.*.yml", "foo-secret: bar")

	// define output folder
	outdir := fmt.Sprintf("tmp-out-%v", rand.Int())
	defer ossafe.RemoveAll(outdir)

	privateKey, _ := rsa.GenerateKey(crypto_rand.Reader, 2048)
	err := IterateFiles(tmpDir, "*.yml", outdir, "", crypto.CreateEncryptionProvider(&privateKey.PublicKey))
	assert.NoError(t, err)

	outputFilename := tmpDir2 + "/" + filepath.Base(tmpFile.Name)
	_, err = os.Stat(outputFilename)
	assert.NoError(t, err)
}

func TestShouldCreateOutdir(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer ossafe.RemoveAll(tmpDir)
	tmpOutDir := fmt.Sprintf("tmp-out-%v", rand.Int())
	defer ossafe.RemoveAll(tmpOutDir)
	_ = tempfile.NewT(t, tmpDir, "foo.*.yml", "")

	err := IterateFiles(tmpDir, "*.yml", tmpOutDir, "", nil)
	assert.NoError(t, err)

	_, err = os.Stat(tmpOutDir)
	assert.NoError(t, err)
}

func TestShouldIgnoreOutdirCreationIfExists(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer ossafe.RemoveAll(tmpDir)
	tmpOutDir, _ := ioutil.TempDir("./", "tmp-out-*")
	defer ossafe.RemoveAll(tmpOutDir)
	_ = tempfile.NewT(t, tmpDir, "foo.*.yml", "")

	err := IterateFiles(tmpDir, "*.yml", tmpOutDir, "", nil)
	assert.NoError(t, err)

	_, err = os.Stat(tmpOutDir)
	assert.NoError(t, err)
}

func TestShouldChangeFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		format   string
		out      string
	}{
		{name: "Should replace ext", filename: "foo/bar/test.yml", format: "", out: "foo/bar/test.yml"},
		{name: "Should replace ext", filename: "foo/bar/test.yml", format: "json", out: "foo/bar/test.json"},
		{name: "Should replace ext", filename: "test.yml", format: "json", out: "test.json"},
		{name: "Should replace ext", filename: "test.json", format: "yaml", out: "test.yaml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := updateFilename(tt.filename, tt.format)

			assert.Equal(t, tt.out, filename)
		})
	}
}

func TestShouldChangeFilepathToOutdir(t *testing.T) {
	tests := []struct {
		name     string
		workdir  string
		outdir   string
		filename string
		out      string
	}{
		{name: "Should replace ext", workdir: "tmp", outdir: "tmp-out", filename: "tmp/bar/test.yml", out: "tmp-out/bar/test.yml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := updateOutdir(tt.workdir, tt.outdir, tt.filename)

			assert.Equal(t, tt.out, filename)
		})
	}
}

package fileutils

import (
	crypto_rand "crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/dimw/simple-secrets-encryptor/test/tempfile"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestShouldIterate(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	tmpDir2, _ := ioutil.TempDir(tmpDir, "tmp2-*")
	defer os.RemoveAll(tmpDir)
	_ = tempfile.Create(tmpDir2, "foo.*.yml", "")

	err := IterateFiles(tmpDir, "*.yml", "", nil)
	assert.Nil(t, err)
}

func TestShouldOutputToOutdir(t *testing.T) {
	// create folder with secrets
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer os.RemoveAll(tmpDir)
	tmpDir2, _ := ioutil.TempDir(tmpDir, "subdir-*")
	defer os.RemoveAll(tmpDir2)
	tmpFile := tempfile.Create(tmpDir2, "foo.*.yml", "foo-secret: bar")
	defer tmpFile.Remove()

	// define output folder
	outdir := fmt.Sprintf("tmp-out-%v", rand.Int())
	defer os.RemoveAll(outdir)

	privateKey, _ := rsa.GenerateKey(crypto_rand.Reader, 2048)
	err := IterateFiles(tmpDir, "*.yml", outdir, crypto.CreateEncryptionProvider(&privateKey.PublicKey))
	assert.Nil(t, err)

	outputFilename := tmpDir2 + "/" + filepath.Base(tmpFile.Name)
	_, err = os.Stat(outputFilename)
	assert.Nil(t, err)
}

func TestShouldCreateOutdir(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer os.RemoveAll(tmpDir)
	tmpOutDir := fmt.Sprintf("tmp-out-%v", rand.Int())
	defer os.RemoveAll(tmpOutDir)
	_ = tempfile.Create(tmpDir, "foo.*.yml", "")

	err := IterateFiles(tmpDir, "*.yml", tmpOutDir, nil)
	assert.Nil(t, err)

	_, err = os.Stat(tmpOutDir)
	assert.Nil(t, err)
}

func TestShouldIgnoreOutdirCreationIfExists(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer os.RemoveAll(tmpDir)
	tmpOutDir, _ := ioutil.TempDir("./", "tmp-out-*")
	defer os.RemoveAll(tmpOutDir)
	_ = tempfile.Create(tmpDir, "foo.*.yml", "")

	err := IterateFiles(tmpDir, "*.yml", tmpOutDir, nil)
	assert.Nil(t, err)

	_, err = os.Stat(tmpOutDir)
	assert.Nil(t, err)
}

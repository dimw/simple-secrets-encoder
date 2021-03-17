package fileutils

import (
	"github.com/dimw/simple-secrets-encryptor/test/tempfile"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestShouldIterate(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer os.RemoveAll(tmpDir)
	_ = tempfile.Create(tmpDir, "foo.*.yml", "")

	err := IterateFiles(tmpDir, "*.yml", nil)
	assert.Nil(t, err)
}

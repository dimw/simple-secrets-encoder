package fileutils

import (
	"github.com/dimw/simple-secrets-encryptor/testhelper/tempfile"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestGlob(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer os.RemoveAll(tmpDir)
	tmpYmlFile := tempfile.Create(tmpDir, "foo.*.yml", "")
	_ = tempfile.Create(tmpDir, "bar.*.txt", "")

	files, err := Glob(tmpDir, "**/*.{yml,yaml}")
	assert.Nil(t, err)

	assert.Equal(t, 1, len(files))
	assert.Equal(t, tmpYmlFile.Name, files[0])
}

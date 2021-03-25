package fileutils

import (
	"io/ioutil"
	"testing"

	"github.com/dimw/simple-secrets-encryptor/testhelper/ossafe"

	"github.com/dimw/simple-secrets-encryptor/testhelper/tempfile"
	"github.com/stretchr/testify/assert"
)

func TestGlob(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	defer ossafe.RemoveAll(tmpDir)
	tmpYmlFile := tempfile.NewT(t, tmpDir, "foo.*.yml", "")
	_ = tempfile.NewT(t, tmpDir, "bar.*.txt", "")

	files, err := Glob(tmpDir, "**/*.{yml,yaml}")
	assert.NoError(t, err)

	assert.Equal(t, 1, len(files))
	assert.Equal(t, tmpYmlFile.Name, files[0])
}

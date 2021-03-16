package glob

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestGlob(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-*")
	tmpYmlFile, _ := ioutil.TempFile(tmpDir, "foo.*.yml")
	tmpYmlFile.Close()
	tmpTxtFile, _ := ioutil.TempFile(tmpDir, "bar.*.txt")
	tmpTxtFile.Close()

	files, err := Glob(tmpDir, "**/*.{yml,yaml}")
	assert.Nil(t, err)

	assert.Equal(t, 1, len(files))
	assert.Equal(t, tmpYmlFile.Name(), files[0])

	//os.Remove(tmpDir + "/" + tmpYmlFile.Name())
	//os.Remove(tmpDir + "/" + tmpTxtFile.Name())
	err = os.RemoveAll(tmpDir)
	assert.Nil(t, err)
}

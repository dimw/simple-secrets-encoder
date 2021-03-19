package tempfile

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type TempFile struct {
	Name string
}

func New(dir string, pattern string, content string) *TempFile {
	tmpFile, _ := ioutil.TempFile(dir, pattern)
	_, _ = tmpFile.WriteString(content)
	_ = tmpFile.Close()

	return &TempFile{
		Name: tmpFile.Name(),
	}
}

func NewT(t *testing.T, dir string, pattern string, content string) *TempFile {
	tmp, _ := ioutil.TempFile(dir, pattern)
	_, _ = tmp.WriteString(content)
	_ = tmp.Close()

	tempFile := &TempFile{
		Name: tmp.Name(),
	}

	t.Cleanup(tempFile.Remove)

	return tempFile
}

func (tf *TempFile) Remove() {
	_ = os.Remove(tf.Name)
}

func (tf *TempFile) Content() string {
	data, _ := ioutil.ReadFile(tf.Name)

	return strings.TrimSpace(string(data))
}

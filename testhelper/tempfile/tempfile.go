package tempfile

import (
	"io/ioutil"
	"os"
	"strings"
)

type TempFile struct {
	Name string
}

func Create(dir string, pattern string, content string) *TempFile {
	tmpFile, _ := ioutil.TempFile(dir, pattern)
	_, _ = tmpFile.WriteString(content)
	_ = tmpFile.Close()

	return &TempFile{
		Name: tmpFile.Name(),
	}
}

func (tf *TempFile) Remove() {
	_ = os.Remove(tf.Name)
}

func (tf *TempFile) Content() string {
	data, _ := ioutil.ReadFile(tf.Name)

	return strings.TrimSpace(string(data))
}

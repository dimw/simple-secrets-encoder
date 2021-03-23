package tempfile

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestCreate(t *testing.T) {
	tmpFile := NewT(t, "./", "foo.*.txt", "bar")

	assert.Regexp(t, regexp.MustCompile(`foo\.\d*\.txt`), tmpFile.Name)
	data, err := ioutil.ReadFile(tmpFile.Name)
	assert.NoError(t, err)

	content := string(data)
	assert.Equal(t, "bar", content)
}

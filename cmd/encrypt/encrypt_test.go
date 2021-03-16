package encrypt

import (
	"fmt"
	"github.com/dimw/simple-secrets-encryptor/cmd"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func TestShouldEncrypt(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-secrets-*")
	tmpFile, _ := ioutil.TempFile(tmpDir, "foo.*.yml")
	_ = tmpFile.Close()

	generateRsaArgs := cmd.GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("private.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("public.%v.pem", rand.Int()),
		KeySize:            2048,
	}
	defer os.Remove(generateRsaArgs.PublicKeyFilename)
	defer os.Remove(generateRsaArgs.PrivateKeyFilename)
	err := cmd.GenerateRSA(generateRsaArgs)
	assert.Nil(t, err)

	args := Args{
		PublicKeyFilename: generateRsaArgs.PublicKeyFilename,
		Workdir:           tmpDir,
		FilenamePattern:   "*.yml",
	}

	err = Encrypt(args)
	assert.Nil(t, err)

	_ = os.RemoveAll(tmpDir)
}

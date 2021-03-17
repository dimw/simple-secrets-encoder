package encrypt

import (
	"fmt"
	generate_keys "github.com/dimw/simple-secrets-encryptor/cmd/generate-keys"
	"github.com/dimw/simple-secrets-encryptor/test/tempfile"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func TestShouldEncrypt(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-secrets-*")
	defer os.RemoveAll(tmpDir)
	_ = tempfile.Create(tmpDir, "foo.*.yml", "")

	generateRsaArgs := generate_keys.GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("private.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("public.%v.pem", rand.Int()),
		KeySize:            2048,
	}
	defer os.Remove(generateRsaArgs.PublicKeyFilename)
	defer os.Remove(generateRsaArgs.PrivateKeyFilename)

	err := generate_keys.GenerateRSA(generateRsaArgs)
	assert.Nil(t, err)

	args := Args{
		PublicKeyFilename: generateRsaArgs.PublicKeyFilename,
		Workdir:           tmpDir,
		FilenamePattern:   "*.yml",
	}

	err = Encrypt(args)
	assert.Nil(t, err)
}

func TestShouldFailDueToMissingPublicKey(t *testing.T) {
	args := Args{
		PublicKeyFilename: "",
	}

	err := Encrypt(args)
	assert.NotNil(t, err)
}

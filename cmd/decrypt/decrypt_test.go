package decrypt

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/dimw/simple-secrets-encryptor/testhelper/ossafe"

	generate_keys "github.com/dimw/simple-secrets-encryptor/cmd/generate-keys"
	"github.com/dimw/simple-secrets-encryptor/testhelper/tempfile"
	"github.com/stretchr/testify/assert"
)

func TestShouldDecrypt(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("./", "tmp-secrets-*")
	defer ossafe.RemoveAll(tmpDir)
	_ = tempfile.NewT(t, tmpDir, "foo.*.yml", "")

	generateRsaArgs := generate_keys.GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("private.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("public.%v.pem", rand.Int()),
		KeySize:            2048,
	}
	defer ossafe.Remove(generateRsaArgs.PublicKeyFilename)
	defer ossafe.Remove(generateRsaArgs.PrivateKeyFilename)
	err := generate_keys.GenerateRSA(generateRsaArgs)
	assert.NoError(t, err)

	args := Args{
		PrivateKeyFilename: generateRsaArgs.PrivateKeyFilename,
		Workdir:            tmpDir,
		FilenamePattern:    "*.yml",
	}

	err = Decrypt(args)
	assert.NoError(t, err)
}

func TestShouldFailDueToMissingPublicKey(t *testing.T) {
	args := Args{
		PrivateKeyFilename: "",
	}

	err := Decrypt(args)
	assert.NotNil(t, err)
}

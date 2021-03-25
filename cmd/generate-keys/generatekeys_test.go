package generatekeys

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/dimw/simple-secrets-encryptor/testhelper/tempfile"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateKeyFiles(t *testing.T) {
	args := GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("tmp-private.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("tmp-public.%v.pem", rand.Int()),
		KeySize:            512,
	}

	err := GenerateRSA(args)
	assert.NoError(t, err)

	privateKey, err := ioutil.ReadFile(args.PrivateKeyFilename)
	assert.Regexp(t, regexp.MustCompile("-----BEGIN RSA PRIVATE KEY-----\n[-A-Za-z0-9+=/\n]+\n-----END RSA PRIVATE KEY-----"), strings.TrimSpace(string(privateKey)))

	publicKey, err := ioutil.ReadFile(args.PublicKeyFilename)
	assert.Regexp(t, regexp.MustCompile("-----BEGIN RSA PUBLIC KEY-----\n[-A-Za-z0-9+=/\n]+\n-----END RSA PUBLIC KEY-----"), strings.TrimSpace(string(publicKey)))

	tearDown(args)
}

func TestShouldNotCreateKeyFilesToAvoidOverwritingPrivateKey(t *testing.T) {
	tmpFile := tempfile.NewT(t, "./", "tmp-private.*.key", "")

	args := GenerateRSAArgs{
		PrivateKeyFilename: tmpFile.Name,
		PublicKeyFilename:  fmt.Sprintf("tmp-public.%v.pem", rand.Int()),
		KeySize:            512,
	}

	err := GenerateRSA(args)
	assert.NotNil(t, err)

	tearDown(args)
}

func TestShouldNotCreateKeyFilesToAvoidOverwritingPublicKey(t *testing.T) {
	tempFile := tempfile.NewT(t, "./", "tmp-private.*.pem", "")

	args := GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("tmp-private-%v.key", rand.Int()),
		PublicKeyFilename:  tempFile.Name,
		KeySize:            512,
	}

	err := GenerateRSA(args)
	assert.Error(t, err)

	tearDown(args)
}

func tearDown(args GenerateRSAArgs) {
	_ = os.Remove(args.PrivateKeyFilename)
	_ = os.Remove(args.PublicKeyFilename)
}

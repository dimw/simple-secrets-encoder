package rsaio

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	generate_keys "github.com/dimw/simple-secrets-encryptor/cmd/generate-keys"
	"github.com/stretchr/testify/assert"
)

func TestShouldLoadPrivateKey(t *testing.T) {
	args := generate_keys.GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("tmp-private.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("tmp-public.%v.pem", rand.Int()),
		KeySize:            2048,
	}

	err := generate_keys.GenerateRSA(args)
	assert.NoError(t, err)

	privateKey, _ := LoadPrivateKey(args.PrivateKeyFilename)
	assert.Equal(t, 256, privateKey.Size())

	_ = os.Remove(args.PublicKeyFilename)
	_ = os.Remove(args.PrivateKeyFilename)
}

func TestShouldLoadPublicKey(t *testing.T) {
	args := generate_keys.GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("tmp-private2.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("tmp-public2.%v.pem", rand.Int()),
		KeySize:            2048,
	}

	err := generate_keys.GenerateRSA(args)
	assert.NoError(t, err)

	publicKey, _ := LoadPublicKey(args.PublicKeyFilename)
	assert.Equal(t, 256, publicKey.Size())

	_ = os.Remove(args.PublicKeyFilename)
	_ = os.Remove(args.PrivateKeyFilename)
}

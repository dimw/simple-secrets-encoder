package rsaio

import (
	"fmt"
	"github.com/dimw/simple-secrets-encryptor/cmd"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

func TestShouldLoadPrivateKey(t *testing.T) {
	args := cmd.GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("tmp-private.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("tmp-public.%v.pem", rand.Int()),
		KeySize:            2048,
	}

	err := cmd.GenerateRSA(args)
	assert.Nil(t, err)

	privateKey, _ := LoadPrivateKey(args.PrivateKeyFilename)
	assert.Equal(t, 256, privateKey.Size())

	_ = os.Remove(args.PublicKeyFilename)
	_ = os.Remove(args.PrivateKeyFilename)
}

func TestShouldLoadPublicKey(t *testing.T) {
	args := cmd.GenerateRSAArgs{
		PrivateKeyFilename: fmt.Sprintf("tmp-private2.%v.key", rand.Int()),
		PublicKeyFilename:  fmt.Sprintf("tmp-public2.%v.pem", rand.Int()),
		KeySize:            2048,
	}

	err := cmd.GenerateRSA(args)
	assert.Nil(t, err)

	publicKey, _ := LoadPublicKey(args.PublicKeyFilename)
	assert.Equal(t, 256, publicKey.Size())

	_ = os.Remove(args.PublicKeyFilename)
	_ = os.Remove(args.PrivateKeyFilename)
}

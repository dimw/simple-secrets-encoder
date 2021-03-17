package process

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/dimw/simple-secrets-encryptor/crypto"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestEncryptDecryptSecrets(t *testing.T) {
	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["foo-secret"] = "secret-bar"
	data["fooToken"] = "token-bar"
	data["foo-s-e-c-r-e-t"] = "dummy value"
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	encryptionProvider := crypto.CreateEncryptionProvider(&privateKey.PublicKey)

	got, err := Walk(data, encryptionProvider)

	assert.Nil(t, err)

	assert.Equal(t, "bar", got["foo"])
	encodedStringRegexp := regexp.MustCompile(`ENC\[rsa,data:[-A-Za-z0-9+=/]+]`)
	assert.Regexp(t, encodedStringRegexp, got["foo-secret"])
	assert.Regexp(t, encodedStringRegexp, got["fooToken"])

	// Setting a non-secret key to a secret
	got["foo-s-e-c-r-e-t"] = got["foo-secret"]

	decryptionProvider := crypto.CreateDecryptionProvider(privateKey)
	got, err = Walk(got, decryptionProvider)
	assert.Equal(t, "bar", got["foo"])
	assert.Equal(t, "secret-bar", got["foo-secret"])
	assert.Equal(t, "token-bar", got["fooToken"])
	assert.Equal(t, "secret-bar", got["foo-s-e-c-r-e-t"])
}

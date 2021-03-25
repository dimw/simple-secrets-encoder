package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	provider := Provider{
		encryptedValueRegexp: regexp.MustCompile(`ENC\[.+]`),
		publicKey:            &privateKey.PublicKey,
	}

	encryptedVal, err := provider.Encrypt("foo")
	assert.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile(`ENC\[rsa,data:.+]`), encryptedVal)
}

func TestEncryptShouldFailDueToBrokenPublicKey(t *testing.T) {
	publicKey := rsa.PublicKey{}

	provider := Provider{
		encryptedValueRegexp: regexp.MustCompile(`ENC\[.+]`),
		publicKey:            &publicKey,
	}

	_, err := provider.Encrypt("foo")
	assert.NotNil(t, err)
}

func TestEncryptShouldReturnOriginalValueIfAlreadyEncrypted(t *testing.T) {
	provider := Provider{
		encryptedValueRegexp: regexp.MustCompile(`ENC\[.+]`),
	}

	val, err := provider.Encrypt("ENC[foo]")
	assert.NoError(t, err)
	assert.Equal(t, "ENC[foo]", val)
}

func TestDecryptShouldReturnUndecrypted(t *testing.T) {
	provider := Provider{
		encryptedValueRegexp: regexp.MustCompile(`ENC\[.*]`),
	}

	val, err := provider.Decrypt("foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", val)
}

func TestDecryptShouldComplainAboutUnsupportedKey(t *testing.T) {
	provider := Provider{
		encryptedValueRegexp:        regexp.MustCompile(`ENC\[.*]`),
		encryptedValueParsingRegexp: regexp.MustCompile(`ENC\[(?P<method>\w+),.*]`),
	}

	_, err := provider.Decrypt("ENC[plain,foo]")
	assert.NotNil(t, err)
}

func TestDecryptShouldComplainAboutBrokenBase64Data(t *testing.T) {
	provider := Provider{
		encryptedValueRegexp:        regexp.MustCompile(`ENC\[.*]`),
		encryptedValueParsingRegexp: regexp.MustCompile(`ENC\[(?P<method>\w+),data:(?P<data>.+)]`),
	}

	_, err := provider.Decrypt("ENC[rsa,data:§$§]")
	assert.NotNil(t, err)
}

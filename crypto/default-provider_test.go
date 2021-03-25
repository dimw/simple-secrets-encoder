package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSecretKey(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{name: "Should not encode value 'foo'", key: "foo", want: false},
		{name: "Should encode 'bar-secret' value", key: "bar-secret", want: true},
		{name: "Should encode 'barToken' value", key: "barToken", want: true},
		{name: "Should encode 'password' value", key: "password", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
			p := CreateEncryptionProvider(&privateKey.PublicKey)

			assert.Equal(t, tt.want, p.IsSecretKey(tt.key))
		})
	}
}

func TestIsEncrypted(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "Not encoded value", value: "foo", want: false},
		{name: "Not encoded value", value: "ENC(foo)", want: false},
		{name: "Not encoded value", value: "xENC[barToken]", want: false},
		{name: "Not encoded value", value: "ENC[barToken]x", want: false},
		{name: "Encoded value", value: "ENC[]", want: true},
		{name: "Encoded value", value: "ENC[barToken]", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateEncryptionProvider(nil)

			assert.Equal(t, tt.want, p.IsEncrypted(tt.value), fmt.Sprintf("%v", tt.value))
		})
	}
}

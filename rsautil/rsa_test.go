package rsautil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRsaKey(t *testing.T) {
	priKey, pubKey, err := GenerateRsaKey(1024)
	assert.NotEmpty(t, priKey)
	assert.NotEmpty(t, pubKey)
	assert.Nil(t, err)
}

func TestGenerateRsaKeyBase64(t *testing.T) {
	rsaKey, err := GenerateRsaKeyBase64(2048)
	assert.NotEmpty(t, rsaKey)
	assert.Nil(t, err)
}

func TestGenerateRsaKeyHex(t *testing.T) {
	rsaKey, err := GenerateRsaKeyHex(1024)
	assert.NotEmpty(t, rsaKey)
	assert.Nil(t, err)
}

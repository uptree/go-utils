package encrypt

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestAes256EcbPkcs7EncryptBase64(t *testing.T) {
	encrypted, err := Aes256EcbPkcs7EncryptBase64("hello", "bc4b2a76b9719d91")
	assert.Equalf(t, "i5MJO9ExGESP5akFLJ7EQA==", encrypted, "FAIL")
	assert.Nil(t, err)
}

func TestAesAes256EcbPkcs7DecryptBase64(t *testing.T) {
	msg := "i5MJO9ExGESP5akFLJ7EQA=="
	decrypted, err := Aes256EcbPkcs7DecryptBase64(msg, "bc4b2a76b9719d91")
	assert.Equalf(t, "hello", decrypted, "FAIL")
	assert.Nil(t, err)
}

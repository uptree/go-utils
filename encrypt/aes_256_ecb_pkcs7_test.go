package encrypt

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestAes256EcbPkcs7EncryptBase64(t *testing.T) {
	encrypted, err := Aes256EcbPkcs7EncryptBase64("hello", "b0baee9d279d34fa1dfd71aadb908c3f")
	assert.Equalf(t, "OisW/4Y22wVjY/0dNZbjOg==", encrypted, "FAIL")
	assert.Nil(t, err)
}

func TestAesAes256EcbPkcs7DecryptBase64(t *testing.T) {
	msg := "OisW/4Y22wVjY/0dNZbjOg=="
	decrypted, err := Aes256EcbPkcs7DecryptBase64(msg, "b0baee9d279d34fa1dfd71aadb908c3f")
	assert.Equalf(t, "hello", decrypted, "FAIL")
	assert.Nil(t, err)
}

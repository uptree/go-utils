package encrypt

import (
	"encoding/base64"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestDesEbcPkcs5EncryptBase64(t *testing.T) {
	key := "12345678"
	encrypted, err := DesEbcPkcs5EncryptBase64("hello", key)
	assert.Equalf(t, "uhbGoCVxJa8=", encrypted, "FAIL")
	assert.Nil(t, err, "FAIL")
}

func TestDesEbcPkcs5DecryptBase64(t *testing.T) {
	key := "12345678"
	decrypted, err := DesEbcPkcs5DecryptBase64("uhbGoCVxJa8=", key)
	assert.Equalf(t, "hello", decrypted, "FAIL")
	assert.Nil(t, err, "FAIL")
}

func TestDesEdeEcbEncrypt(t *testing.T) {
	key := []byte("12345678bc4b2a76b9719d91")
	ciphertext, _ := DesEdeEcbEncrypt([]byte("hello"), key)
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	assert.Equalf(t, "Qa6CzunG4x4=", base64Ciphertext, "FAIL")
}

func TestDesEdeEcbDecrypt(t *testing.T) {
	key := []byte("12345678bc4b2a76b9719d91")
	ciphertext, _ := base64.StdEncoding.DecodeString("Qa6CzunG4x4=")
	plaintext, _ := DesEdeEcbDecrypt(ciphertext, key)
	assert.Equalf(t, "hello", string(plaintext), "FAIL")
}

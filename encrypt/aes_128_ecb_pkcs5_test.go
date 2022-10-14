package encrypt

import (
	"encoding/base64"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestAes128EcbPkcs5Encrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext := Aes128EcbPkcs5Encrypt([]byte("hello"), key)
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	assert.Equalf(t, "jRHUh558C/r1p+WWsJ1TbA==", base64Ciphertext, "FAIL")
}

func TestAes128EcbPkcs5Decrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	base64Ciphertext := "jRHUh558C/r1p+WWsJ1TbA=="
	ciphertext, _ := base64.StdEncoding.DecodeString(base64Ciphertext)
	plaintext := Aes128EcbPkcs5Decrypt(ciphertext, key)
	assert.Equalf(t, "hello", string(plaintext), "FAIL")
}

func TestAes128EcbPkcs5EncryptBase64(t *testing.T) {
	encrypted := Aes128EcbPkcs5EncryptBase64("hello", "bc4b2a76b9719d91")
	assert.Equalf(t, "jRHUh558C/r1p+WWsJ1TbA==", encrypted, "FAIL")
}

func TestAes128EcbPkcs5DecryptBase64(t *testing.T) {
	msg := "jRHUh558C/r1p+WWsJ1TbA=="
	decrypted := Aes128EcbPkcs5DecryptBase64(msg, "bc4b2a76b9719d91")
	assert.Equalf(t, "hello", decrypted, "FAIL")
}

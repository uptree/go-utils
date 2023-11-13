package encrypt

import (
	"encoding/base64"
	"testing"

	"github.com/uptree/go-utils/assert"
	"github.com/uptree/go-utils/hashutil"
)

func TestAesCbcPkcs5Encrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext, err := AesCbcPkcs5Encrypt([]byte("hello"), key)
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	assert.Equalf(t, "W4ro9Ffii0pqFZBWYB+QZw==", base64Ciphertext, "FAIL")
	assert.Nil(t, err)
}

func TestAesCbcPkcs5Decrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext := "W4ro9Ffii0pqFZBWYB+QZw=="
	bytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	text, err := AesCbcPkcs5Decrypt(bytes, key)
	assert.Equalf(t, "hello", string(text), "FAIL")
	assert.Nil(t, err)
}

func TestAesCbcPkcs5EncryptBase64(t *testing.T) {
	key := hashutil.Md5Hex("hello")[8:24]
	ciphertext, err := AesCbcPkcs5EncryptBase64("hello", key)
	assert.Equalf(t, "W4ro9Ffii0pqFZBWYB+QZw==", ciphertext, "FAIL")
	assert.Nil(t, err)
}

func TestAesCbcPkcs5DecryptBase64(t *testing.T) {
	key := hashutil.Md5Hex("hello")[8:24]
	ciphertext := "W4ro9Ffii0pqFZBWYB+QZw=="
	text, err := AesCbcPkcs5DecryptBase64(ciphertext, key)
	assert.Equalf(t, "hello", text, "FAIL")
	assert.Nil(t, err)
}

func TestAesCbcEncryptWithSalt(t *testing.T) {
	key := hashutil.Md5Hex("hello")[8:24]
	ciphertext, err := AesCbcEncryptWithSalt([]byte("hello"), []byte(key))
	text, err := AesCbcDecryptWithSalt(ciphertext, []byte(key))
	assert.Equalf(t, "hello", string(text), "FAIL")
	assert.Nil(t, err)
}

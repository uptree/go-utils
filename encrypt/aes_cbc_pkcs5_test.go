package encrypt

import (
	"encoding/base64"
	"testing"

	"github.com/uptree/go-utils/assert"
	"github.com/uptree/go-utils/hashutil"
)

func TestAesEncrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext, err := AesEncrypt([]byte("hello"), key)
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	assert.Equalf(t, "W4ro9Ffii0pqFZBWYB+QZw==", base64Ciphertext, "FAIL")
	assert.Nil(t, err)
}

func TestAesDecrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext := "W4ro9Ffii0pqFZBWYB+QZw=="
	bytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	text, err := AesDecrypt(bytes, key)
	assert.Equalf(t, "hello", string(text), "FAIL")
	assert.Nil(t, err)
}

func TestAesEncryptMsg(t *testing.T) {
	key := hashutil.Md5Hex("hello")[8:24]
	ciphertext, err := AesEncryptMsg("hello", key)
	assert.Equalf(t, "W4ro9Ffii0pqFZBWYB+QZw==", ciphertext, "FAIL")
	assert.Nil(t, err)
}

func TestAesDecryptMsg(t *testing.T) {
	key := hashutil.Md5Hex("hello")[8:24]
	ciphertext := "W4ro9Ffii0pqFZBWYB+QZw=="
	text, err := AesDecryptMsg(ciphertext, key)
	assert.Equalf(t, "hello", text, "FAIL")
	assert.Nil(t, err)
}

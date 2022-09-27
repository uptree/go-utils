package encrypt

import (
	"encoding/base64"
	"github.com/uptree/go-utils/assert"
	"testing"
)

func TestDesEBCEncryptMsg(t *testing.T) {
	key := "12345678"
	encrypted, err := DesEBCEncryptMsg("hello", key)
	assert.Equalf(t, "uhbGoCVxJa8=", encrypted, "FAIL")
	assert.Nil(t, err, "FAIL")
}

func TestDesEBCDecryptMsg(t *testing.T) {
	key := "12345678"
	decrypted, err := DesEBCDecryptMsg("uhbGoCVxJa8=", key)
	assert.Equalf(t, "hello", decrypted, "FAIL")
	assert.Nil(t, err, "FAIL")
}

func TestDESedeECBEncrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext, _ := DESedeECBEncrypt([]byte("hello"), key)
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	assert.Equalf(t, "nIyAIyI3PQY=", base64Ciphertext, "FAIL")
}

func TestDESedeECBDecrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext, _ := base64.StdEncoding.DecodeString("nIyAIyI3PQY=")
	plaintext, _ := DESedeECBDecrypt(ciphertext, key)
	assert.Equalf(t, "hello", string(plaintext), "FAIL")
}

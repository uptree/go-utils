package encrypt

import (
	"encoding/base64"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestAesECBEncrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	ciphertext := AesECBEncrypt([]byte("hello"), key)
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	assert.Equalf(t, "jRHUh558C/r1p+WWsJ1TbA==", base64Ciphertext, "FAIL")
}

func TestAesECBDecrypt(t *testing.T) {
	key := []byte("bc4b2a76b9719d91")
	base64Ciphertext := "jRHUh558C/r1p+WWsJ1TbA=="
	ciphertext, _ := base64.StdEncoding.DecodeString(base64Ciphertext)
	plaintext := AesECBDecrypt(ciphertext, key)
	assert.Equalf(t, "hello", string(plaintext), "FAIL")
}

func TestAesECBEncryptMsg(t *testing.T) {
	encrypted := AesECBEncryptMsg("hello", "bc4b2a76b9719d91")
	assert.Equalf(t, "jRHUh558C/r1p+WWsJ1TbA==", encrypted, "FAIL")
}

func TestAesECBDecryptMsg(t *testing.T) {
	msg := "jRHUh558C/r1p+WWsJ1TbA=="
	decrypted := AesECBDecryptMsg(msg, "bc4b2a76b9719d91")
	assert.Equalf(t, "hello", decrypted, "FAIL")
}

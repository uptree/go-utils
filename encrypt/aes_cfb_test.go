package encrypt

import (
	"testing"

	"github.com/uptree/go-utils/assert"
	"github.com/uptree/go-utils/hashutil"
)

func TestAesCfdEncryptWithSalt(t *testing.T) {
	key := hashutil.Md5Hex("hello")[8:24]
	ciphertext, err := AesCfbEncryptWithSalt([]byte("hello"), []byte(key))
	text, err := AesCfbDecryptWithSalt(ciphertext, []byte(key))
	assert.Equalf(t, "hello", string(text), "FAIL")
	assert.Nil(t, err)
}

func TestAesCfdEncryptBase64(t *testing.T) {
	key := hashutil.Md5Hex("hello")[8:24]
	ciphertext, err := AesCfbEncryptBase64("hello", key)
	text, err := AesCfbDecryptBase64(ciphertext, key)
	assert.Equalf(t, "hello", text, "FAIL")
	assert.Nil(t, err)
}

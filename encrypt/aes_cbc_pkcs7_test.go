package encrypt

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestAesCbcPkcs7EncryptBase64(t *testing.T) {
	encodingAesKey := "YjBiYWVlOWQyNzlkMzRmYTFkZmQ3MWFhZGI5MDhjM2Y"
	resMsg := "x7tgJYMVOp/O+zLYw4uosZFFWYCNmO1V7Ty99Gm3BjU="
	base64EncryptMsg, err := AesCbcPkcs7EncryptBase64("hello", encodingAesKey)
	assert.Equalf(t, resMsg, base64EncryptMsg, "FAIL")
	assert.Nil(t, err)
}

func TestAesCbcPkcs7DecryptBase64(t *testing.T) {
	encodingAesKey := "YjBiYWVlOWQyNzlkMzRmYTFkZmQ3MWFhZGI5MDhjM2Y"
	base64EncryptMsg := "x7tgJYMVOp/O+zLYw4uosZFFWYCNmO1V7Ty99Gm3BjU="
	msg, err := AesCbcPkcs7DecryptBase64(base64EncryptMsg, encodingAesKey)
	assert.Equalf(t, "hello", msg, "FAIL")
	assert.Nil(t, err)
}

func TestGenerateEncodingAesKey(t *testing.T) {
	encodingAesKey := GenerateEncodingAesKey("b0baee9d279d34fa1dfd71aadb908c3f")
	assert.Equalf(t, 43, len(encodingAesKey), "FAIL")
}

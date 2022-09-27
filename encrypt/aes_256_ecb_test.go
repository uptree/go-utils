package encrypt

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestAesECBEncrypt256Msg(t *testing.T) {
	encrypted, err := AesECBEncrypt256Msg("hello", "bc4b2a76b9719d91")
	assert.Equalf(t, "i5MJO9ExGESP5akFLJ7EQA==", encrypted, "FAIL")
	assert.Nil(t, err)
}

func TestAesECBDecrypt256Msg(t *testing.T) {
	msg := "i5MJO9ExGESP5akFLJ7EQA=="
	decrypted, err := AesECBDecrypt256Msg(msg, "bc4b2a76b9719d91")
	assert.Equalf(t, "hello", decrypted, "FAIL")
	assert.Nil(t, err)
}

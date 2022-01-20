package encode

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	t.Logf("Base64Encode: %v", Base64Encode([]byte("hello")))
}

func TestBase64Decode(t *testing.T) {
	b, err := Base64Decode("hello")
	t.Logf("Base64Decode: %v, %v", string(b), err)
}

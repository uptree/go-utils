package hashes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	res := len(MD5("text"))
	assert.Equalf(t, 32, res, "failed")
}

func TestSHA1(t *testing.T) {
	res := len(SHA1("text"))
	assert.Equalf(t, 40, res, "failed")
}

func TestSHA256(t *testing.T) {
	res := len(SHA256("text"))
	assert.Equalf(t, 64, res, "failed")
}

func TestSHA512(t *testing.T) {
	res := len(SHA512("text"))
	assert.Equalf(t, 128, res, "failed")
}

func TestHmacSHA256(t *testing.T) {
	res := len(HmacSHA256("text", "secret"))
	assert.Equalf(t, 64, res, "failed")
}

package hashes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	res := len(MD5("1"))
	assert.Equalf(t, 32, res, "failed")
}

func TestSHA1(t *testing.T) {
	res := len(SHA1("1"))
	assert.Equalf(t, 40, res, "failed")
}

func TestSHA256(t *testing.T) {
	res := len(SHA256("1"))
	assert.Equalf(t, 64, res, "failed")
}

func TestSHA512(t *testing.T) {
	res := len(SHA512("1"))
	assert.Equalf(t, 128, res, "failed")
}

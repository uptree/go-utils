package hashutil

import (
	"encoding/hex"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestMd5(t *testing.T) {
	res := len(Md5([]byte("text")))
	assert.Equalf(t, 16, res, "FAIL")
}

func TestMd5Hex(t *testing.T) {
	s := []byte("text")
	assert.Equalf(t, hex.EncodeToString(Md5(s)), Md5Hex(s), "FAIL")
}

func TestSha1(t *testing.T) {
	res := len(Sha1([]byte("text")))
	assert.Equalf(t, 20, res, "FAIL")
}

func TestSha1Hex(t *testing.T) {
	s := []byte("text")
	assert.Equalf(t, hex.EncodeToString(Sha1(s)), Sha1Hex(s), "FAIL")
}

func TestSha256(t *testing.T) {
	res := len(Sha256([]byte("text")))
	assert.Equalf(t, 32, res, "FAIL")
}

func TestSha256Hex(t *testing.T) {
	s := []byte("text")
	assert.Equalf(t, hex.EncodeToString(Sha256(s)), Sha256Hex(s), "FAIL")
}

func TestSha512(t *testing.T) {
	res := len(Sha512([]byte("text")))
	assert.Equalf(t, 64, res, "FAIL")
}

func TestSha512Hex(t *testing.T) {
	s := []byte("text")
	assert.Equalf(t, hex.EncodeToString(Sha512(s)), Sha512Hex(s), "FAIL")
}

func TestHmacMd5(t *testing.T) {
	res := len(HmacMd5([]byte("text"), []byte("secret")))
	assert.Equalf(t, 16, res, "FAIL")
}

func TestHmacMd5Hex(t *testing.T) {
	s := []byte("text")
	key := []byte("secret")
	assert.Equalf(t, hex.EncodeToString(HmacMd5(s, key)), HmacMd5Hex(s, key), "FAIL")
}

func TestHmacSha1(t *testing.T) {
	res := len(HmacSha1([]byte("text"), []byte("secret")))
	assert.Equalf(t, 20, res, "FAIL")
}

func TestHmacSha1Hex(t *testing.T) {
	s := []byte("text")
	key := []byte("secret")
	assert.Equalf(t, hex.EncodeToString(HmacSha1(s, key)), HmacSha1Hex(s, key), "FAIL")
}

func TestHmacSha256(t *testing.T) {
	res := len(HmacSha256([]byte("text"), []byte("secret")))
	assert.Equalf(t, 32, res, "FAIL")
}

func TestHmacSha256Hex(t *testing.T) {
	s := []byte("text")
	key := []byte("secret")
	assert.Equalf(t, hex.EncodeToString(HmacSha256(s, key)), HmacSha256Hex(s, key), "FAIL")
}

func TestHmacSha512(t *testing.T) {
	res := len(HmacSha512([]byte("text"), []byte("secret")))
	assert.Equalf(t, 64, res, "FAIL")
}

func TestHmacSha512Hex(t *testing.T) {
	s := []byte("text")
	key := []byte("secret")
	assert.Equalf(t, hex.EncodeToString(HmacSha512(s, key)), HmacSha512Hex(s, key), "FAIL")
}

func TestTime33(t *testing.T) {
	num := Time33("hello")
	t.Log(num)
}

func TestCrc32(t *testing.T) {
	num := Crc32([]byte("hello"))
	t.Log(num)
}

func TestCrc64(t *testing.T) {
	num := Crc64([]byte("hello"))
	t.Log(num)
}

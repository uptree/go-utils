package hashutil

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/crc32"
	"hash/crc64"
)

// Md5 hashes using md5 algorithm
func Md5(b []byte) []byte {
	return Hashes(md5.New(), b)
}

// Md5Hex hashes using md5 algorithm
func Md5Hex(b []byte) string {
	return StringHashes(md5.New(), b)
}

// Sha1 hashes using sha1 algorithm
func Sha1(b []byte) []byte {
	return Hashes(sha1.New(), b)
}

// Sha1Hex hashes using sha1 algorithm
func Sha1Hex(b []byte) string {
	return StringHashes(sha1.New(), b)
}

// Sha256 hashes using sha256 algorithm
func Sha256(b []byte) []byte {
	return Hashes(sha256.New(), b)
}

// Sha256Hex hashes using sha256 algorithm
func Sha256Hex(b []byte) string {
	return StringHashes(sha256.New(), b)
}

// Sha512 hashes using sha512 algorithm
func Sha512(b []byte) []byte {
	return Hashes(sha512.New(), b)
}

// Sha512Hex hashes using sha512 algorithm
func Sha512Hex(b []byte) string {
	return StringHashes(sha512.New(), b)
}

// HmacMd5 hashes using md5 algorithm with a secret
func HmacMd5(b []byte, secret []byte) []byte {
	algorithm := hmac.New(md5.New, secret)
	return Hashes(algorithm, b)
}

// HmacMd5Hex hashes using md5 algorithm with a secret
func HmacMd5Hex(b []byte, secret []byte) string {
	return StringHashes(hmac.New(md5.New, secret), b)
}

// HmacSha1 hashes using sha1 algorithm with a secret
func HmacSha1(b []byte, secret []byte) []byte {
	return Hashes(hmac.New(sha1.New, secret), b)
}

// HmacSha1Hex hashes using sha1 algorithm with a secret
func HmacSha1Hex(b []byte, secret []byte) string {
	return StringHashes(hmac.New(sha1.New, secret), b)
}

// HmacSha256 hashes using sha256 algorithm with a secret
func HmacSha256(b []byte, secret []byte) []byte {
	return Hashes(hmac.New(sha256.New, secret), b)
}

// HmacSha256Hex hashes using sha256 algorithm with a secret
func HmacSha256Hex(b []byte, secret []byte) string {
	return StringHashes(hmac.New(sha256.New, secret), b)
}

// HmacSha512 hashes using sha512 algorithm with a secret
func HmacSha512(b []byte, secret []byte) []byte {
	return Hashes(hmac.New(sha512.New, secret), b)
}

// HmacSha512Hex hashes using sha512 algorithm with a secret
func HmacSha512Hex(b []byte, secret []byte) string {
	return StringHashes(hmac.New(sha512.New, secret), b)
}

// StringHashes hashes string
func StringHashes(algorithm hash.Hash, b []byte) string {
	return hex.EncodeToString(Hashes(algorithm, b))
}

// Hashes hashes
func Hashes(algorithm hash.Hash, b []byte) []byte {
	algorithm.Write(b)
	return algorithm.Sum(nil)
}

// Time33 ...
func Time33(text string) int {
	hash := 5381
	length := len(text)
	for i := 0; i < length; i++ {
		hash += (hash << 5) + int(text[i])
	}
	// other: hash ^ (hash >> 32)
	return hash & 0x7FFFFFFF
}

// Crc32 ...
func Crc32(b []byte) uint32 {
	return crc32.ChecksumIEEE(b)
}

// Crc64 ... 可用于腾讯云/阿里云对象存储校验
func Crc64(b []byte) uint64 {
	tab := crc64.MakeTable(crc64.ECMA)
	return crc64.Checksum(b, tab)
}

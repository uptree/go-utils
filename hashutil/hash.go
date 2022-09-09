package hashutil

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// Md5 hashes using md5 algorithm
func Md5(b []byte) []byte {
	return Hashes(md5.New(), b)
}

// Md5Hex hashes using md5 algorithm
func Md5Hex(text string) string {
	return StringHashes(md5.New(), text)
}

// Sha1 hashes using sha1 algorithm
func Sha1(b []byte) []byte {
	return Hashes(sha1.New(), b)
}

// Sha1Hex hashes using sha1 algorithm
func Sha1Hex(text string) string {
	return StringHashes(sha1.New(), text)
}

// Sha256 hashes using sha256 algorithm
func Sha256(b []byte) []byte {
	return Hashes(sha256.New(), b)
}

// Sha256Hex hashes using sha256 algorithm
func Sha256Hex(text string) string {
	return StringHashes(sha256.New(), text)
}

// Sha512 hashes using sha512 algorithm
func Sha512(b []byte) []byte {
	return Hashes(sha512.New(), b)
}

// Sha512Hex hashes using sha512 algorithm
func Sha512Hex(text string) string {
	algorithm := sha512.New()
	return StringHashes(algorithm, text)
}

// HmacSha256 hashes using sha256 algorithm with a secret
func HmacSha256(b []byte, secret []byte) []byte {
	algorithm := hmac.New(sha256.New, secret)
	return Hashes(algorithm, b)
}

// HmacSha256Hex hashes using sha256 algorithm with a secret
func HmacSha256Hex(text, secret string) string {
	algorithm := hmac.New(sha256.New, []byte(secret))
	return StringHashes(algorithm, text)
}

// HmacSha512 hashes using sha512 algorithm with a secret
func HmacSha512(b []byte, secret []byte) []byte {
	algorithm := hmac.New(sha512.New, secret)
	return Hashes(algorithm, b)
}

// HmacSha512Hex hashes using sha512 algorithm with a secret
func HmacSha512Hex(text, secret string) string {
	algorithm := hmac.New(sha512.New, []byte(secret))
	return StringHashes(algorithm, text)
}

// StringHashes hashes string
func StringHashes(algorithm hash.Hash, text string) string {
	return hex.EncodeToString(Hashes(algorithm, []byte(text)))
}

// Hashes hashes
func Hashes(algorithm hash.Hash, b []byte) []byte {
	algorithm.Write(b)
	return algorithm.Sum(nil)
}

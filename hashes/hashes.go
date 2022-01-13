package hashes

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// MD5 hashes using md5 algorithm
func MD5(text string) string {
	algorithm := md5.New()
	return stringHasher(algorithm, text)
}

// SHA1 hashes using sha1 algorithm
func SHA1(text string) string {
	algorithm := sha1.New()
	return stringHasher(algorithm, text)
}

// SHA256 hashes using sha256 algorithm
func SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

// SHA512 hashes using sha512 algorithm
func SHA512(text string) string {
	algorithm := sha512.New()
	return stringHasher(algorithm, text)
}

// HmacSHA256 hashes using sha256 algorithm with a secret
func HmacSHA256(text, secret string) string {
	algorithm := hmac.New(sha256.New, []byte(secret))
	return stringHasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

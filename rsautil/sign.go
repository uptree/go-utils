package rsautil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"

	"github.com/uptree/go-utils/hashutil"
)

// SignSha256Base64 signs using sha256 algorithm base64
func SignSha256Base64(msg, priKey string) (string, error) {
	priBytes, err := base64.StdEncoding.DecodeString(priKey)
	if err != nil {
		return "", err
	}
	sign, err := signSha256([]byte(msg), priBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

// VerifySignSha256Base64 verifies sign using sha256 algorithm base64
func VerifySignSha256Base64(msg, sign, pubKey string) error {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	pubBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return err
	}
	return verifySignSha256([]byte(msg), signBytes, pubBytes)
}

// SignSha256Hex signs using sha256 algorithm hex
func SignSha256Hex(msg, priKey string) (string, error) {
	priBytes, err := hex.DecodeString(priKey)
	if err != nil {
		return "", err
	}
	sign, err := signSha256([]byte(msg), priBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sign), nil

}

// VerifySignSha256Hex verifies sign using sha256 algorithm hex
func VerifySignSha256Hex(msg, sign, pubKey string) error {
	signBytes, err := hex.DecodeString(sign)
	if err != nil {
		return err
	}
	pubBytes, err := hex.DecodeString(pubKey)
	if err != nil {
		return err
	}
	return verifySignSha256([]byte(msg), signBytes, pubBytes)
}

// signSha256 签名
func signSha256(msg, priKey []byte) ([]byte, error) {
	privateKey, err := x509.ParsePKCS8PrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	hashed := hashutil.Sha256(msg)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// verifySignSha256 验签
func verifySignSha256(msg, sign, pubKey []byte) error {
	publicKey, err := x509.ParsePKIXPublicKey(pubKey)
	if err != nil {
		return err
	}
	hashed := hashutil.Sha256(msg)
	return rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed, sign)
}

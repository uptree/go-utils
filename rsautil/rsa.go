package rsautil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

var (
	RSABitsErr                = errors.New("bits 1024 or 2048")
	RSAPrivateKeyPemDecodeErr = errors.New("private key pem.decode error")
	RSAPublicKeyPemDecodeErr  = errors.New("public key pem.decode error")
)

type RsaKey struct {
	PrivateKey string // PKCS#8
	PublicKey  string // PKCS#8
}

// GenerateRsaKey 生成公钥私钥
func GenerateRsaKey(bits int) ([]byte, []byte, error) {
	if bits != 1024 && bits != 2048 {
		return nil, nil, RSABitsErr
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	priKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}
	pubKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return priKey, pubKey, nil
}

// GenerateRsaKeyBase64 生成公钥私钥 Base64
func GenerateRsaKeyBase64(bits int) (RsaKey, error) {
	priKey, pubKey, err := GenerateRsaKey(bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: base64.StdEncoding.EncodeToString(priKey),
		PublicKey:  base64.StdEncoding.EncodeToString(pubKey),
	}, nil
}

// GenerateRsaKeyHex 生成公钥私钥 Hex
func GenerateRsaKeyHex(bits int) (RsaKey, error) {
	priKey, pubKey, err := GenerateRsaKey(bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: hex.EncodeToString(priKey),
		PublicKey:  hex.EncodeToString(pubKey),
	}, nil
}

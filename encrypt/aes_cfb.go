package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// AesCfbEncryptWithSalt 加盐加密
func AesCfbEncryptWithSalt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// AesCfbDecryptWithSalt 加盐解密
func AesCfbDecryptWithSalt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if (len(ciphertext) % aes.BlockSize) != 0 {
		return nil, errors.New("invalid ciphertext block size")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return PKCS5UnPadding(ciphertext), nil
}

// AesCfbEncryptBase64 加密
func AesCfbEncryptBase64(msg, encodingAesKey string) (string, error) {
	ciphertext, err := AesCfbEncryptWithSalt([]byte(msg), []byte(encodingAesKey))
	if err != nil {
		return msg, err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), err
}

// AesCfbDecryptBase64 解密
func AesCfbDecryptBase64(base64EncryptMsg, encodingAesKey string) (string, error) {
	encryptMsg, err := base64.StdEncoding.DecodeString(base64EncryptMsg)
	if err != nil {
		return "", err
	}
	bs, err := AesCfbDecryptWithSalt(encryptMsg, []byte(encodingAesKey))
	if err != nil {
		return "", err
	}
	return string(bs), err
}

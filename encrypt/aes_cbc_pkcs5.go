package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// AesCbcPkcs5Encrypt AES加密
func AesCbcPkcs5Encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(PasswdPadding16(key))
	if err != nil {
		return nil, err
	}
	// AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)
	// IV必须等于块block的长度16字节
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cipherText := make([]byte, len(plaintext))
	blockMode.CryptBlocks(cipherText, plaintext)
	return cipherText, nil
}

// AesCbcPkcs5Decrypt AES解密
func AesCbcPkcs5Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(PasswdPadding16(key))
	if err != nil {
		return nil, err
	}
	// AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	// IV必须等于块block的长度16字节
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	plainText := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(plainText, ciphertext)
	return PKCS5UnPadding(plainText), nil
}

// AesCbcPkcs5EncryptBase64 加密+base64编码
func AesCbcPkcs5EncryptBase64(plaintext, key string) (string, error) {
	ciphertext, err := AesCbcPkcs5Encrypt([]byte(plaintext), []byte(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesCbcPkcs5DecryptBase64 解密+base64编码
func AesCbcPkcs5DecryptBase64(ciphertext, key string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := AesCbcPkcs5Decrypt(cipherBytes, []byte(key))
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func AesCbcEncryptWithSalt(plaintext, key []byte) ([]byte, error) {
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[0:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func AesCbcDecryptWithSalt(ciphertext, key []byte) ([]byte, error) {
	var block cipher.Block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("iciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)
	return ciphertext, nil
}

package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesEncrypt AES加密
func AesEncrypt(plaintext, key []byte) ([]byte, error) {
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

// AesDecrypt AES解密
func AesDecrypt(ciphertext, key []byte) ([]byte, error) {
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

// AesEncryptMsg 加密+base64编码
func AesEncryptMsg(plaintext, key string) (string, error) {
	ciphertext, err := AesEncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesDecryptMsg 解密+base64编码
func AesDecryptMsg(ciphertext, key string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := AesDecrypt(cipherBytes, []byte(key))
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

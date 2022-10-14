// MySQL中默认的aes加密算法是aes-128-ecb

package encrypt

import (
	"crypto/aes"
	"encoding/base64"
)

// Aes128EcbPkcs5Encrypt 加密 => MYSQL AES_ENCRYPT
func Aes128EcbPkcs5Encrypt(plaintext, key []byte) []byte {
	block, _ := aes.NewCipher(PasswdPadding16(key))
	plaintext = PKCS5Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	size := block.BlockSize()
	for bs, be := 0, size; bs < len(plaintext); bs, be = bs+size, be+size {
		block.Encrypt(ciphertext[bs:be], plaintext[bs:be])
	}
	return ciphertext
}

// Aes128EcbPkcs5Decrypt 解密 => MYSQL AES_DECRYPT
func Aes128EcbPkcs5Decrypt(ciphertext, key []byte) []byte {
	block, _ := aes.NewCipher(PasswdPadding16(key))
	plaintext := make([]byte, len(ciphertext))
	size := block.BlockSize()
	for bs, be := 0, size; bs < len(ciphertext); bs, be = bs+size, be+size {
		block.Decrypt(plaintext[bs:be], ciphertext[bs:be])
	}
	return PKCS5UnPadding(plaintext)
}

// Aes128EcbPkcs5EncryptBase64 加密+base64编码 => MYSQL TO_BASE64(AES_ENCRYPT)
func Aes128EcbPkcs5EncryptBase64(msg, key string) string {
	encrypted := Aes128EcbPkcs5Encrypt([]byte(msg), []byte(key))
	if encrypted == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(encrypted)
}

// Aes128EcbPkcs5DecryptBase64 解密+base64编码
func Aes128EcbPkcs5DecryptBase64(msg, key string) string {
	encrypted, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return ""
	}
	decrypted := Aes128EcbPkcs5Decrypt(encrypted, []byte(key))
	return string(decrypted)
}

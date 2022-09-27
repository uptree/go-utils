package encrypt

import (
	"bytes"
	"errors"
)

const (
	PKCS7AesKeyLen = 32 //AES算法的密钥，长度为32字节
	PKCS5AesKeyLen = 16
)

// PKCS5Padding 填充补齐
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	// 计算最后一个分组,缺少多少个字节
	padding := blockSize - len(plaintext)%blockSize
	// 创建一个大小为padding的切片, 每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 将padText添加到原始数据的后边, 将最后一个分组缺少的字节数补齐
	return append(plaintext, padText...)
}

// PKCS5UnPadding 去除填充数据
func PKCS5UnPadding(plaintext []byte) []byte {
	// 计算数据的总长度
	length := len(plaintext)
	// 根据填充的字节值得到填充的次数
	padding := int(plaintext[length-1])
	// 将尾部填充的padding个字节去掉
	return plaintext[:(length - padding)]
}

// PKCS7Padding 填充
func PKCS7Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - (len(plaintext) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...)
}

// PKCS7UnPadding 去除
func PKCS7UnPadding(plaintext []byte, blockSize int) ([]byte, error) {
	plaintextLen := len(plaintext)
	if nil == plaintext || plaintextLen == 0 {
		return nil, errors.New("PKCS7UnPadding error nil or zero")
	}
	if plaintextLen%blockSize != 0 {
		return nil, errors.New("PKCS7UnPadding text not a multiple of the block size")
	}
	paddingLen := int(plaintext[plaintextLen-1])
	if len(plaintext) < plaintextLen-paddingLen || plaintextLen-paddingLen < 0 {
		return nil, errors.New("PKCS7UnPadding plaintext len too small")
	}
	return plaintext[:plaintextLen-paddingLen], nil
}

func ZeroPadding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padText := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(plaintext, padText...)
}

func ZeroUnPadding(plaintext []byte) []byte {
	return bytes.TrimFunc(plaintext,
		func(r rune) bool {
			return r == rune(0)
		})
}

// PasswdPadding8 Fill 0x00 if the length of key less than 8
func PasswdPadding8(key []byte) []byte {
	newKey := make([]byte, 8)
	copy(newKey, key)
	return newKey
}

// PasswdPadding16 秘钥补齐
func PasswdPadding16(key []byte) []byte {
	newKey := make([]byte, 16)
	copy(newKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			newKey[j] ^= key[i]
		}
	}
	return newKey
}

// PasswdPadding24 秘钥补齐
func PasswdPadding24(key []byte) []byte {
	newKey := make([]byte, 24)
	copy(newKey, key)
	for i := 24; i < len(key); {
		for j := 0; j < 24 && i < len(key); j, i = j+1, i+1 {
			newKey[j] ^= key[i]
		}
	}
	return newKey
}

// PasswdPadding32 Fill 0x00 if the length of key less than 32
func PasswdPadding32(key []byte) []byte {
	newKey := make([]byte, 32)
	copy(newKey, key)
	return newKey
}

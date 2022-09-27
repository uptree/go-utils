// 参照企业微信实现AES加解密
// AES采用CBC模式，数据采用PKCS#7填充至32字节的倍数；
// IV初始向量大小为16字节，取AESKey前16字节。
// EncodingAESKey：用于消息体的加密，
// 长度固定为43个字符，从a-z, A-Z, 0-9共62个字符中选取，是AESKey的Base64编码。
// 解码后即为32字节长的AESKey
// https://developer.work.weixin.qq.com/devtool/introduce?id=36388

package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"strings"
)

// AesCBCEncrypt 加密
func AesCBCEncrypt(msg, encodingAesKey string) ([]byte, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAesKey + "=")
	if nil != err {
		return nil, err
	}
	// 数据采用PKCS#7填充至32字节的倍数
	padMsg := PKCS7Padding([]byte(msg), PKCS7AesKeyLen)
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(padMsg))
	// IV初始向量大小为16字节，取AESKey前16字节
	iv := aesKey[:aes.BlockSize]
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, padMsg)

	base64Msg := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(base64Msg, ciphertext)

	return base64Msg, nil
}

// AesCBCDecrypt 解密
func AesCBCDecrypt(base64EncryptMsg, encodingAesKey string) ([]byte, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAesKey + "=")
	if nil != err {
		return nil, err
	}
	encryptMsg, err := base64.StdEncoding.DecodeString(base64EncryptMsg)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	if len(encryptMsg) < aes.BlockSize {
		return nil, errors.New("encrypt_msg size is not valid")
	}
	iv := aesKey[:aes.BlockSize]
	if len(encryptMsg)%aes.BlockSize != 0 {
		return nil, errors.New("encrypt_msg not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptMsg, encryptMsg)

	encryptMsg, err = PKCS7UnPadding(encryptMsg, PKCS7AesKeyLen)
	return encryptMsg, err
}

// AesCBCEncryptMsg 加密
func AesCBCEncryptMsg(msg, encodingAesKey string) (string, error) {
	bs, err := AesCBCEncrypt(msg, encodingAesKey)
	if err != nil {
		return msg, err
	}
	return string(bs), err
}

// AesCBCDecryptMsg 解密
func AesCBCDecryptMsg(base64EncryptMsg, encodingAesKey string) (string, error) {
	bs, err := AesCBCDecrypt(base64EncryptMsg, encodingAesKey)
	if err != nil {
		return base64EncryptMsg, err
	}
	return string(bs), err
}

// GenerateEncodingAesKey 生成43位EncodingAesKey，len(key)最好等于32
func GenerateEncodingAesKey(key string) string {
	var aseKey []byte
	if len(key) >= PKCS7AesKeyLen {
		aseKey = []byte(key[:PKCS7AesKeyLen])
	} else {
		aseKey = PKCS7Padding([]byte(key), PKCS7AesKeyLen)
	}
	return strings.Trim(base64.StdEncoding.EncodeToString(aseKey), "=")
}

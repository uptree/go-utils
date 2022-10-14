// 微信支付退款结果通知
// 解密步骤如下：
//（1）对加密串A做base64解码，得到加密串B
//（2）对商户key做md5，得到32位小写key* ( key设置路径：微信商户平台(pay.weixin.qq.com)-->账户设置-->API安全-->密钥设置 )
//（3）用key*对加密串B做AES-256-ECB解密（PKCS7Padding）
// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10

package encrypt

import (
	"crypto/aes"
	"encoding/base64"
)

// Aes256EcbPkcs7Encrypt 加密
func Aes256EcbPkcs7Encrypt(plaintext, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(PasswdPadding32(key))
	if err != nil {
		return nil, err
	}
	plaintext = PKCS7Padding(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	for st, ed := 0, aes.BlockSize; ed <= len(plaintext); st, ed = st+aes.BlockSize, ed+aes.BlockSize {
		c.Encrypt(ciphertext[st:ed], plaintext[st:ed])
	}
	return ciphertext, nil
}

// Aes256EcbPkcs7Decrypt 解密
func Aes256EcbPkcs7Decrypt(ciphertext, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(PasswdPadding32(key))
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext))
	for st, ed := 0, aes.BlockSize; ed <= len(plaintext); st, ed = st+aes.BlockSize, ed+aes.BlockSize {
		c.Decrypt(plaintext[st:ed], ciphertext[st:ed])
	}
	return PKCS7UnPadding(plaintext, aes.BlockSize)
}

// Aes256EcbPkcs7EncryptBase64 加密+base64编码
func Aes256EcbPkcs7EncryptBase64(msg, key string) (string, error) {
	encrypted, err := Aes256EcbPkcs7Encrypt([]byte(msg), []byte(key))
	if encrypted == nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Aes256EcbPkcs7DecryptBase64 解密+base64编码
func Aes256EcbPkcs7DecryptBase64(msg, key string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}
	decrypted, err := Aes256EcbPkcs7Decrypt(encrypted, []byte(key))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// Java默认DES算法使用DES/ECB/PKCS5Padding
// 而golang认为这种方式是不安全的，所以故意没有提供这种加密方式
// 3DES，也称为3DESede或TripleDES，是三重数据加密，且可以逆推的一种算法方案

package encrypt

import (
	"crypto/des"
	"encoding/base64"
	"errors"
)

//DesEBCEncrypt 加密
func DesEBCEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := des.NewCipher(PasswdPadding8(key))
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(plaintext)%bs != 0 {
		return nil, errors.New("padding is invalid")
	}
	ciphertext := make([]byte, len(plaintext))
	dst := ciphertext
	for len(plaintext) > 0 {
		block.Encrypt(dst, plaintext[:bs])
		plaintext = plaintext[bs:]
		dst = dst[bs:]
	}
	return ciphertext, nil
}

//DesEBCDecrypt 解密
func DesEBCDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := des.NewCipher(PasswdPadding8(key))
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(ciphertext)%bs != 0 {
		return nil, errors.New("ciphertext size is invalid")
	}
	plaintext := make([]byte, len(ciphertext))
	dst := plaintext
	for len(ciphertext) > 0 {
		block.Decrypt(dst, ciphertext[:bs])
		ciphertext = ciphertext[bs:]
		dst = dst[bs:]
	}
	return plaintext, nil
}

// DesEBCEncryptMsg 加密+Base64 兼容java默认
func DesEBCEncryptMsg(msg, key string) (string, error) {

	data := PKCS5Padding([]byte(msg), 8)
	encrypted, err := DesEBCEncrypt(data, []byte(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DesEBCDecryptMsg 解密+Base64 兼容java默认
func DesEBCDecryptMsg(msg, key string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}
	decrypted, err := DesEBCDecrypt(encrypted, []byte(key))
	if err != nil {
		return "", err
	}
	return string(PKCS5UnPadding(decrypted)), nil
}

// DESedeECBEncrypt 3DES加密desede-ECB
func DESedeECBEncrypt(plaintext, key []byte) ([]byte, error) {
	desKey := make([]byte, 24, 24)
	copy(desKey, PasswdPadding24(key))
	k1 := desKey[:8]
	k2 := desKey[8:16]
	k3 := desKey[16:]
	plaintext = PKCS7Padding(plaintext, 8)

	buf1, err := DesEBCEncrypt(plaintext, k1)
	if err != nil {
		return nil, err
	}
	buf2, err := DesEBCDecrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	ciphertext, err := DesEBCEncrypt(buf2, k3)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// DESedeECBDecrypt 3DES解密desede-ECB
func DESedeECBDecrypt(src, key []byte) ([]byte, error) {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]
	buf1, err := DesEBCDecrypt(src, k3)
	if err != nil {
		return nil, err
	}
	buf2, err := DesEBCEncrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	out, err := DesEBCDecrypt(buf2, k1)
	if err != nil {
		return nil, err
	}
	return PKCS7UnPadding(out, 8)
}


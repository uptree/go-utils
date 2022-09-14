package rsautil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strings"

	"github.com/uptree/go-utils/fileutil"
)

// Pem Block Type
var (
	PKCS1PrivateType = "RSA PRIVATE KEY" // PKCS#1
	PKCS1PublicType  = "RSA PUBLIC KEY"  // PKCS#1

	PKCS8PrivateType = "PRIVATE KEY" // PKCS#8
	PKCS8PublicType  = "PUBLIC KEY"  // PKCS#8
)

// DecodePublicKeyBase64 解码pem格式公钥
func DecodePublicKeyBase64(pubKey string) (string, error) {
	b, err := decodePublicKey([]byte(pubKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// DecodePrivateKeyBase64 解码pem格式私钥
func DecodePrivateKeyBase64(priKey string) (string, error) {
	b, err := decodePrivateKey([]byte(priKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// GetPublicKey 设置公钥
func GetPublicKey(pubKey []byte) (*rsa.PublicKey, error) {
	// decode public key
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, RSAPublicKeyPemDecodeErr
	}
	// x509 parse public key
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		pub, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	return pub.(*rsa.PublicKey), err
}

// GetPrivateKey 设置私钥
func GetPrivateKey(priKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, RSAPrivateKeyPemDecodeErr
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pri, nil
	}
	pri2, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pri2.(*rsa.PrivateKey), nil
}

// WritePrivateKeyToPem 私钥保存到pem文件中
func WritePrivateKeyToPem(priKey []byte, pemName string) error {
	privateFile, err := fileutil.CreateFile(pemName)
	if err != nil {
		return err
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: PKCS8PrivateType, Bytes: priKey}
	return pem.Encode(privateFile, &privateBlock)
}

// WritePublicKeyToPem 公钥保存到pem文件中
func WritePublicKeyToPem(pubKey []byte, pemName string) error {
	publicFile, err := fileutil.CreateFile(pemName)
	if err != nil {
		return err
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: PKCS8PublicType, Bytes: pubKey}
	return pem.Encode(publicFile, &publicBlock)
}

// GenerateRsaKeyPemFile 生成RSA私钥和公钥，保存到文件中
func GenerateRsaKeyPemFile(bits int, dirPath string) error {
	priKey, pubKey, err := GenerateRsaKey(bits)
	if err != nil {
		return err
	}
	if dirPath != "" {
		dirPath = strings.TrimSuffix(dirPath, "/") + "/"
	}
	privateFile, err := fileutil.CreateFile(dirPath + "private_key.pem")
	if err != nil {
		return err
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: PKCS8PrivateType, Bytes: priKey}
	if err = pem.Encode(privateFile, &privateBlock); err != nil {
		return err
	}
	publicFile, err := fileutil.CreateFile(dirPath + "public_key.pem")
	if err != nil {
		return err
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: PKCS8PublicType, Bytes: pubKey}
	return pem.Encode(publicFile, &publicBlock)
}

func decodePublicKey(pubKey []byte) ([]byte, error) {
	publicBlock, _ := pem.Decode(pubKey)
	if publicBlock == nil {
		return nil, RSAPublicKeyPemDecodeErr
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return x509.MarshalPKIXPublicKey(publicKey)
}

func decodePrivateKey(priKey []byte) ([]byte, error) {
	privateBlock, _ := pem.Decode(priKey)
	if privateBlock == nil {
		return nil, RSAPrivateKeyPemDecodeErr
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return x509.MarshalPKCS8PrivateKey(privateKey)
}

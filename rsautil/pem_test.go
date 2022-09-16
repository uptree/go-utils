package rsautil

import (
	"testing"

	"github.com/uptree/go-utils/assert"
	"github.com/uptree/go-utils/fileutil"
)

func TestGenerateRsaKeyPemFile(t *testing.T) {
	err := GenerateRsaKeyPemFile(1024, "testdata/")
	assert.Nil(t, err)
}

func TestWriteWritePrivateKeyToPem(t *testing.T) {
	prikey, _, _ := GenerateRsaKey(2048)
	err := WritePrivateKeyToPem(prikey, "testdata/Private.pem")
	assert.Nil(t, err)
}

func TestWritePublicKeyToPem(t *testing.T) {
	_, pubKey, _ := GenerateRsaKey(2048)
	err := WritePublicKeyToPem(pubKey, "testdata/Public.pem")
	assert.Nil(t, err)
}

func TestGetPrivateKey(t *testing.T) {
	data, _ := fileutil.ReadFileToBytes("testdata/Private.pem")
	_, err := GetPrivateKey(data)
	assert.Nil(t, err)
}

func TestGetPublicKey(t *testing.T) {
	data, _ := fileutil.ReadFileToBytes("testdata/Public.pem")
	_, err := GetPublicKey(data)
	assert.Nil(t, err)
}

func TestDecodePrivateKeyBase64(t *testing.T) {
	data, _ := fileutil.ReadFileToBytes("testdata/Private.pem")
	priKey, err := DecodePrivateKeyBase64(string(data))
	assert.NotEmpty(t, priKey)
	assert.Nil(t, err)
}

func TestDecodePublicKeyBase64(t *testing.T) {
	data, _ := fileutil.ReadFileToBytes("testdata/Public.pem")
	pubKey, err := DecodePublicKeyBase64(string(data))
	assert.NotEmpty(t, pubKey)
	assert.Nil(t, err)
}

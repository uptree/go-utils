package rsautil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignSha256Base64(t *testing.T) {
	rsaKey, _ := GenerateRsaKeyBase64(1024)
	sign, err := SignSha256Base64("text", rsaKey.PrivateKey)
	assert.NotEmpty(t, sign)
	assert.Nil(t, err)
}

func TestVerifySignSha256Base64(t *testing.T) {
	var publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDeShcKB80gH7NIwpGYHFsQoya3
IDunLUi3wflDLx/VY4g55TqbguogU9lXbr93X+deEp35kfF6E6SRmbEF/zeKe2VF
SW+w2JP1EybxHl00GigAsHr+QbC2UOKRe8GdSQK/83Yei+QiSUT0ronWbj/NOgUD
6TDN26aiaDCHj5nh8wIDAQAB
-----END PUBLIC KEY-----
`

	var privateKey = `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAN5KFwoHzSAfs0jC
kZgcWxCjJrcgO6ctSLfB+UMvH9VjiDnlOpuC6iBT2Vduv3df514SnfmR8XoTpJGZ
sQX/N4p7ZUVJb7DYk/UTJvEeXTQaKACwev5BsLZQ4pF7wZ1JAr/zdh6L5CJJRPSu
idZuP806BQPpMM3bpqJoMIePmeHzAgMBAAECgYB+v65daM9c5tZ6wX71Rg3i8bSs
rLN3Aso5BWDVHzS+ny9ZG80MSVyorEb8pMiGD+hEascrPD19x1+KKiGXcsw+mRvn
t9tt3DwCkqucDkUPMrqZUFeV6YctIVVt4nR0MhK5Z+qnU2AK3gRvSIuErBZAiC5O
utVmW/f6NtihsLdXIQJBAOXTn+KlURGLHyKaLIlTeJU+uwz8n6V4VKqFY7666mW3
DF6nhTqCkXK1lj2DoyKGujudxwRUvkADln/prTkuu8MCQQD3mrltbYCkYNN7g1Rk
CO9DYTnH2ba1STV3hAikqAFfHu32678cyVuGv52Y6OYc4QsOzEAj5B2u9Ao1p8O9
7k4RAkADGLXXxOhxtxElUzR0aTZ/tEeq5iB0h0mEvqVYZYZQ9oVpJRKrdiTlicLL
c2GYTn5l3TtYOJgTnrjBYnGeCp4nAkEA44kC94VWXzg/f9Rq+9aeEPxKvdXbKfMR
uDfUPYPKZuAka/GuWUTM4CI8MCVuOYGwyLN4CcN6Z6kJrM7zeRmCEQJAeLC+0Kny
0mZsvYvpOqIJ2vKb5zrrDBkFicrcCfDxkkBgPZNsVARNgtJHCFIGkHi+ye6hRo7c
C5sYrsqfOHa1OA==
-----END PRIVATE KEY-----
`
	msg := "text"
	priKey, _ := DecodePrivateKeyBase64(privateKey)
	pubKey, _ := DecodePublicKeyBase64(publicKey)
	sign, _ := SignSha256Base64(msg, priKey)
	err := VerifySignSha256Base64(msg, sign, pubKey)
	assert.Nil(t, err)
}

func TestSignSha256Hex(t *testing.T) {
	rsaKey, _ := GenerateRsaKeyHex(1024)
	sign, err := SignSha256Hex("text", rsaKey.PrivateKey)
	assert.NotEmpty(t, sign)
	assert.Nil(t, err)
}

func TestVerifySignSha256Hex(t *testing.T) {
	msg := "text"
	rsaKey, _ := GenerateRsaKeyHex(2048)
	sign, _ := SignSha256Hex(msg, rsaKey.PrivateKey)
	err := VerifySignSha256Hex(msg, sign, rsaKey.PublicKey)
	assert.Nil(t, err)
}

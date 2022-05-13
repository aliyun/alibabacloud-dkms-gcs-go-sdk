package auth

type AlibabaCloudCredentials interface {
	GetAccessKeyId() string
	GetAccessKeySecret() string
}

type RsaKeyPairCredentials struct {
	keyId            string
	privateKeySecret string
}

func NewRsaKeyPairCredentials(keyId, privateKeySecret string) *RsaKeyPairCredentials {
	return &RsaKeyPairCredentials{
		keyId:            keyId,
		privateKeySecret: privateKeySecret,
	}
}

func (rsa *RsaKeyPairCredentials) GetAccessKeyId() string {
	return rsa.keyId
}

func (rsa *RsaKeyPairCredentials) GetAccessKeySecret() string {
	return rsa.privateKeySecret
}

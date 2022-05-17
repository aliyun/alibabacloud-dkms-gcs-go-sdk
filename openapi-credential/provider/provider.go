package provider

import (
	"github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-credential/auth"
)

type AlibabaCloudCredentialsProvider interface {
	GetCredentials() auth.AlibabaCloudCredentials
}

type RsaKeyPairCredentialProvider struct {
	credentials auth.AlibabaCloudCredentials
}

func NewRsaKeyPairCredentialProvider(keyId, privateKeySecret string) *RsaKeyPairCredentialProvider {
	return &RsaKeyPairCredentialProvider{
		credentials: auth.NewRsaKeyPairCredentials(keyId, privateKeySecret),
	}
}

func (p *RsaKeyPairCredentialProvider) GetCredentials() auth.AlibabaCloudCredentials {
	return p.credentials
}

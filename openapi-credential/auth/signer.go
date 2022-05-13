package auth

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
)

type Signer interface {
	SignString(stringToSign, accessKeySecret string) (*string, error)
	GetSignerName() *string
	GetSignerVersion() *string
	GetSignerType() *string
}

type SHA256withRSASigner struct {
}

func NewSHA256withRSASigner() *SHA256withRSASigner {
	return &SHA256withRSASigner{}
}

func (s *SHA256withRSASigner) SignString(stringToSign, accessKeySecret string) (*string, error) {
	block, _ := pem.Decode([]byte(accessKeySecret))
	pkcs1Priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	hashed := sha256.Sum256([]byte(stringToSign))
	sig, err := rsa.SignPKCS1v15(rand.Reader, pkcs1Priv, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return tea.String(fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString(sig))), nil
}

func (s *SHA256withRSASigner) GetSignerName() *string {
	return tea.String("RSA_PKCS1_SHA_256")
}

func (s *SHA256withRSASigner) GetSignerVersion() *string {
	return tea.String("1.0")
}

func (s *SHA256withRSASigner) GetSignerType() *string {
	return tea.String("rsa_key_pair")
}

func GetSigner(credentials AlibabaCloudCredentials) (Signer, error) {
	switch credentials.(type) {
	case *RsaKeyPairCredentials:
		return NewSHA256withRSASigner(), nil
	default:
		return nil, errors.New("only support rsa key pair credential now")
	}
}

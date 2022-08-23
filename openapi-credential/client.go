package openapi_credential

import (
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-credential/auth"
	"github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-credential/provider"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

type Config struct {
	Type             *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	AccessKeyId      *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	PrivateKey       *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
	ClientKeyFile    *string `json:"clientKeyFile,omitempty" xml:"clientKeyFile,omitempty"`
	ClientKeyContent *string `json:"clientKeyContent,omitempty" xml:"clientKeyContent,omitempty"`
	Password         *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetType(v string) *Config {
	s.Type = &v
	return s
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetPrivateKey(v string) *Config {
	s.PrivateKey = &v
	return s
}

func (s *Config) SetClientKeyFile(v string) *Config {
	s.ClientKeyFile = &v
	return s
}

func (s *Config) SetClientKeyContent(v string) *Config {
	s.ClientKeyContent = &v
	return s
}

func (s *Config) SetPassword(v string) *Config {
	s.Password = &v
	return s
}

type ClientKey struct {
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	PrivateKeyData *string `json:"PrivateKeyData,omitempty" xml:"PrivateKeyData,omitempty"`
}

func (s ClientKey) String() string {
	return tea.Prettify(s)
}

func (s ClientKey) GoString() string {
	return s.String()
}

func (s *ClientKey) SetKeyId(v string) *ClientKey {
	s.KeyId = &v
	return s
}

func (s *ClientKey) SetPrivateKeyData(v string) *ClientKey {
	s.PrivateKeyData = &v
	return s
}

type Client struct {
	CredentialsProvider provider.AlibabaCloudCredentialsProvider
	ClientKeyCertPem    string
}

func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.EqualString(config.Type, tea.String("rsa_key_pair"))) {
		if !tea.BoolValue(util.Empty(config.ClientKeyContent)) {
			clientKey := &ClientKey{}
			certPem, privateKeyPem, err := client.parseClientKeyContent(tea.StringValue(config.ClientKeyContent), tea.StringValue(config.Password), clientKey)
			if err != nil {
				return err
			}
			client.ClientKeyCertPem = certPem
			client.CredentialsProvider = provider.NewRsaKeyPairCredentialProvider(tea.StringValue(clientKey.KeyId), privateKeyPem)
		} else if !tea.BoolValue(util.Empty(config.ClientKeyFile)) {
			clientKey := &ClientKey{}
			content, err := ioutil.ReadFile(tea.StringValue(config.ClientKeyFile))
			if err != nil {
				return err
			}
			certPem, privateKeyPem, err := client.parseClientKeyContent(string(content), tea.StringValue(config.Password), clientKey)
			if err != nil {
				return err
			}
			client.ClientKeyCertPem = certPem
			client.CredentialsProvider = provider.NewRsaKeyPairCredentialProvider(tea.StringValue(clientKey.KeyId), privateKeyPem)
		} else {
			client.CredentialsProvider = provider.NewRsaKeyPairCredentialProvider(tea.StringValue(config.AccessKeyId), tea.StringValue(config.PrivateKey))
		}
	} else {
		return errors.New("only support rsa key pair credential provider now")
	}
	return nil
}

func (client *Client) GetAccessKeyId() (_result *string) {
	_result = tea.String(client.CredentialsProvider.GetCredentials().GetAccessKeyId())
	return
}

func (client *Client) GetAccessKeySecret() (_result *string) {
	_result = tea.String(client.CredentialsProvider.GetCredentials().GetAccessKeySecret())
	return
}

func (client *Client) GetClientKeyCertPem() string {
	return client.ClientKeyCertPem
}

func (client *Client) GetSignature(strToSign *string) (_result *string, _err error) {
	credentials := client.CredentialsProvider.GetCredentials()
	signer, _err := auth.GetSigner(credentials)
	if _err != nil {
		return
	}
	_result, _err = signer.SignString(tea.StringValue(strToSign), credentials.GetAccessKeySecret())
	return
}

func (client *Client) parseClientKeyContent(content, password string, clientKey *ClientKey) (string, string, error) {
	err := json.Unmarshal([]byte(content), clientKey)
	if err != nil {
		return "", "", err
	}
	privateKeyData, err := base64.StdEncoding.DecodeString(tea.StringValue(clientKey.PrivateKeyData))
	if err != nil {
		return "", "", err
	}
	blocks, err := pkcs12.ToPEM(privateKeyData, password)
	if err != nil {
		return "", "", err
	}
	return string(pem.EncodeToMemory(blocks[0])), string(pem.EncodeToMemory(blocks[1])), nil
}

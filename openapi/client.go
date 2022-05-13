// This file is auto-generated, don't edit it. Thanks.
package openapi

import (
	dedicatedkmsopenapicredential "github.com/alibabacloud-dkms-gcs-go-sdk/openapi-credential"
	dedicatedkmsopenapiutil "github.com/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

var (
	defaultCert = `-----BEGIN CERTIFICATE-----
MIIDETCCAfkCFEoG8qIWsGo3DUYJztr2DDduzt0sMA0GCSqGSIb3DQEBCwUAMDsx
CzAJBgNVBAYTAmNuMREwDwYDVQQIDAh6aGVqaWFuZzELMAkGA1UEBwwCaHoxDDAK
BgNVBAoMA2ttczAgFw0yMjA0MTQwMzE5NTlaGA8yMTIyMDMyMTAzMTk1OVowTTEL
MAkGA1UEBhMCY24xEzARBgNVBAgMCnpoZW5namlhbmcxCzAJBgNVBAcMAmh6MRww
GgYDVQQKDBNEZWZhdWx0IENvbXBhbnkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAru5l20y1+qDKkLnbb7jqeZWzH8IrqPpunU9JpRhJteEcNh4/
mjYWRRsZrZRLDz4jzUxvQZE6wZT6c1S/A4PAp1T9hZFyk8M/Vcr/dlZJWW4BQx1J
/x9xjPbSlEFfNiHGHAwDo4DCbTplTCm2lrGj4YQokn3ddT+9AOXeD3jahzt9tth5
Wzf3Vo1ajQWhtcKtwJjTfY+RcsagoeJvmJt4DkhkwVTx/DH9Z05ZWIIwKDt2liJ+
bmwbO99JC1pwLRhHs77kmjMWPFF/HLSO11bxeU0sCILz1VcugDO6BAAuBlipwvTk
djLP/nrSxCqCB8pctslHSlStIFNAK8TvXeYILQIDAQABMA0GCSqGSIb3DQEBCwUA
A4IBAQC3ZKcBlrBo8YE3sO6ML9x0qBWa1f5YwzHoXGvO5nO3X2J89MWHBfJlE1zj
ayAHpvE5RxCCP6WpzSwsfuu/XQUTcXrh+KzRiU0zXzkumPX7lqNe2+Cscp1e2QwL
NMCy/6PyfM0Bu6mEGZZhMmwT+prZs5sLk3qa+o90sZGXUAghBllHkcmrCQ2Ja6n7
vewYgO9g7CsmmymQccAdtj3fQoF4NBD0HS+dtSYbLQjMebBx1L6zCryoNXKna0iM
97uvmjfji2K9V1wdSh106YC3MDJZRIwopSOTN/QezehzT+3dbq5axukGa2BWRjVA
V5bon9PzAcW2VDpkcHByzUgrpreV
-----END CERTIFICATE-----`
	defaultKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAru5l20y1+qDKkLnbb7jqeZWzH8IrqPpunU9JpRhJteEcNh4/
mjYWRRsZrZRLDz4jzUxvQZE6wZT6c1S/A4PAp1T9hZFyk8M/Vcr/dlZJWW4BQx1J
/x9xjPbSlEFfNiHGHAwDo4DCbTplTCm2lrGj4YQokn3ddT+9AOXeD3jahzt9tth5
Wzf3Vo1ajQWhtcKtwJjTfY+RcsagoeJvmJt4DkhkwVTx/DH9Z05ZWIIwKDt2liJ+
bmwbO99JC1pwLRhHs77kmjMWPFF/HLSO11bxeU0sCILz1VcugDO6BAAuBlipwvTk
djLP/nrSxCqCB8pctslHSlStIFNAK8TvXeYILQIDAQABAoIBAQCP+KFYFhCID6Hn
7y9NChHgs3ZTIwmv7zzut3zTJeQ0jusbF1YdY3p8scs2cyA7kVcB9nBytEZ5YFWo
ag0zpVH/5hEpxwVX24/a2lHaYTXJJqJfTFHEaWqZK+MwJYVpFe7o0nKENMrP/wmz
33jpGZbZWl3SA0kq4KR7grR6Y2rndJAjhal09ivyTHMOKYVXiMUZznuZufCc40m4
i5FlrDvIagFcB852tlHzKqOICG6fLYZocJJB5dZZRDoQy/CYNqYt1z5mf2/jGIfE
1o+KUOx1KEs5A8/pJ5fzjovcAwtn7uSTOHNXMBLGIQEG/l7foVqlNCkuFotYYMhc
Ikc91Sa5AoGBAN/1XFS+AUca0bR/ffyvfaT4hQJUyUPoG01O8JnmsYtnRCPg3NmA
zvkaR+uqmLkz6ctenV9md5+24TH9sHegbWp4OzNwLv3c2/hmtcl2ZI7rZmrIezuc
3+qQodrVUGs/uwE3b2f0cDbsn69VhwHOE6uIe3KskNBQ5LfYUxN33KJXAoGBAMf1
YXb8bIjMvJnXvkj9qNf6flCaXNTaBDptIMZb4Gf5bjLuruSOYwXtL+9v3rTpuY4u
Cr38/akT41pqevyLiC1a9yWPmEAdbZiRr26RFJnxJAHlgYMPA9efTzM5qjrDrmj2
6dFyR+ZMfdKkfeb2IuT6Z8V5iLMAX7YfUa7KTr8bAoGAJtIBDy9gSVL8xqzSydoh
jVcGt0C7Imo5UU09p2+1ltRvm0h+/7S0K7nCfbc0geXdwMLFxKOeWLD6mA/BZvsH
/Eya9m+btaVy+7vSBvnetGdZeWkG825d4erwI16WcgC6CZvyU7KrQNoDVAEfRdDS
2AjwXRBMic971eJJhDZiqEECgYA3gI3oKBK8OhATp8XCnt4R7Q7IssAXuNg5m3bg
eDomwh5rdsMG39rulzKaBHRNxR5jL7Z9m86Q8ttFio7OsLBA3qIBRAVfrotwMOgA
CbctL7jt6dV0GmPXovmp31S38P7lVbv7DfyVGSBHqq7jv+hZpIPPVXezL+RhOfpD
SRZohwKBgQCmQE3mLHe16EtzkJIBRv4m0+7CUvBqwr2WkndSpwlDZrriy5kdWra/
nKB/ilMWXQjgs6bvwGBRir0B9dYeOSoqcL5av6HJsXyiyye4llLS+tleFoMVYVjf
OBekUMUC54pZ9vIaF+sZsg9hJyTXHceVxrRsBe/Y258xCgh1m6z8Jg==
-----END RSA PRIVATE KEY-----`
)

type Config struct {
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// pkcs1 or pkcs8 PEM format private key
	PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
	// crypto service address
	Endpoint       *string                               `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Protocol       *string                               `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RegionId       *string                               `json:"regionId,omitempty" xml:"regionId,omitempty" pattern:"[a-zA-Z0-9-_]+"`
	ReadTimeout    *int                                  `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	ConnectTimeout *int                                  `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	HttpProxy      *string                               `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	HttpsProxy     *string                               `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	Socks5Proxy    *string                               `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	Socks5NetWork  *string                               `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	NoProxy        *string                               `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	MaxIdleConns   *int                                  `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	UserAgent      *string                               `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	Type           *string                               `json:"type,omitempty" xml:"type,omitempty"`
	Credential     *dedicatedkmsopenapicredential.Client `json:"credential,omitempty" xml:"credential,omitempty"`
	ClientKeyFile  *string                               `json:"clientKeyFile,omitempty" xml:"clientKeyFile,omitempty"`
	// client key content
	ClientKeyContent *string `json:"clientKeyContent,omitempty" xml:"clientKeyContent,omitempty"`
	Password         *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetPrivateKey(v string) *Config {
	s.PrivateKey = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetRegionId(v string) *Config {
	s.RegionId = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetType(v string) *Config {
	s.Type = &v
	return s
}

func (s *Config) SetCredential(v *dedicatedkmsopenapicredential.Client) *Config {
	s.Credential = v
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

type Client struct {
	Endpoint       *string
	RegionId       *string
	Protocol       *string
	ReadTimeout    *int
	ConnectTimeout *int
	HttpProxy      *string
	HttpsProxy     *string
	NoProxy        *string
	UserAgent      *string
	Socks5Proxy    *string
	Socks5NetWork  *string
	MaxIdleConns   *int
	Credential     *dedicatedkmsopenapicredential.Client
}

func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(tea.ToMap(config))) {
		_err = tea.NewSDKError(map[string]interface{}{
			"name":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	if tea.BoolValue(util.Empty(config.Endpoint)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config.endpoint' can not be empty",
		})
		return _err
	}

	if !tea.BoolValue(util.Empty(config.ClientKeyContent)) {
		config.Type = tea.String("rsa_key_pair")
		contentConfig := &dedicatedkmsopenapicredential.Config{
			Type:             config.Type,
			ClientKeyContent: config.ClientKeyContent,
			Password:         config.Password,
		}
		client.Credential, _err = dedicatedkmsopenapicredential.NewClient(contentConfig)
		if _err != nil {
			return _err
		}

	} else if !tea.BoolValue(util.Empty(config.ClientKeyFile)) {
		config.Type = tea.String("rsa_key_pair")
		clientKeyConfig := &dedicatedkmsopenapicredential.Config{
			Type:          config.Type,
			ClientKeyFile: config.ClientKeyFile,
			Password:      config.Password,
		}
		client.Credential, _err = dedicatedkmsopenapicredential.NewClient(clientKeyConfig)
		if _err != nil {
			return _err
		}

	} else if !tea.BoolValue(util.Empty(config.AccessKeyId)) && !tea.BoolValue(util.Empty(config.PrivateKey)) {
		config.Type = tea.String("rsa_key_pair")
		credentialConfig := &dedicatedkmsopenapicredential.Config{
			Type:        config.Type,
			AccessKeyId: config.AccessKeyId,
			PrivateKey:  config.PrivateKey,
		}
		client.Credential, _err = dedicatedkmsopenapicredential.NewClient(credentialConfig)
		if _err != nil {
			return _err
		}

	} else if !tea.BoolValue(util.IsUnset(config.Credential)) {
		client.Credential = config.Credential
	}

	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.RegionId = config.RegionId
	client.UserAgent = config.UserAgent
	client.ReadTimeout = config.ReadTimeout
	client.ConnectTimeout = config.ConnectTimeout
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = config.MaxIdleConns
	return nil
}

func (client *Client) DoRequest(apiName *string, apiVersion *string, protocol *string, method *string, signatureMethod *string, reqBodyBytes []byte, headers map[string]*string, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	var cert interface{}
	var key interface{}
	if runtime.Verify != nil {
		cert = defaultCert
		key = defaultKey
	}
	_runtime := map[string]interface{}{
		"timeouted":      "retry",
		"readTimeout":    tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout": tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":      tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":     tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":        tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"socks5Proxy":    tea.StringValue(util.DefaultString(runtime.Socks5Proxy, client.Socks5Proxy)),
		"socks5NetWork":  tea.StringValue(util.DefaultString(runtime.Socks5NetWork, client.Socks5NetWork)),
		"maxIdleConns":   tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
		"cert":      cert,
		"key":       key,
		"ca":        tea.StringValue(runtime.Verify),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = tea.String("/")
			request_.Headers = tea.Merge(map[string]*string{
				"accept":                tea.String("application/x-protobuf"),
				"host":                  dedicatedkmsopenapiutil.GetHost(client.RegionId, client.Endpoint),
				"date":                  util.GetDateUTCString(),
				"user-agent":            util.GetUserAgent(client.UserAgent),
				"x-kms-apiversion":      apiVersion,
				"x-kms-apiname":         apiName,
				"x-kms-signaturemethod": signatureMethod,
				"x-kms-acccesskeyid":    client.Credential.GetAccessKeyId(),
			}, headers)
			request_.Headers["content-type"] = tea.String("application/x-protobuf")
			request_.Headers["content-length"] = dedicatedkmsopenapiutil.GetContentLength(reqBodyBytes)
			request_.Headers["content-sha256"] = dedicatedkmsopenapiutil.GetContentSHA256(reqBodyBytes)
			request_.Body = tea.ToReader(reqBodyBytes)
			strToSign := dedicatedkmsopenapiutil.GetStringToSign(request_)
			sign, _err := client.Credential.GetSignature(strToSign)
			if _err != nil {
				return _result, _err
			}
			request_.Headers["authorization"] = sign
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			var bodyBytes []byte
			if tea.BoolValue(util.Is4xx(response_.StatusCode)) || tea.BoolValue(util.Is5xx(response_.StatusCode)) {
				bodyBytes, _err = util.ReadAsBytes(response_.Body)
				if _err != nil {
					return _result, _err
				}

				respMap, _err := dedicatedkmsopenapiutil.GetErrMessage(bodyBytes)
				if _err != nil {
					return _result, _err
				}
				_err = tea.NewSDKError(map[string]interface{}{
					"code":    respMap["Code"],
					"message": respMap["Message"],
					"data": map[string]interface{}{
						"httpCode":  tea.IntValue(response_.StatusCode),
						"requestId": respMap["RequestId"],
						"hostId":    respMap["HostId"],
					},
				})
				return _result, _err
			}

			bodyBytes, _err = util.ReadAsBytes(response_.Body)
			if _err != nil {
				return _result, _err
			}

			_result = map[string]interface{}{
				"headers": dedicatedkmsopenapiutil.FilterHeaders(response_.Headers, runtime.Headers),
				"body":    bodyBytes,
			}
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

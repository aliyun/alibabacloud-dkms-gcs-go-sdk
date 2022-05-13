package main

import (
	"fmt"
	dedicatedkmsopenapi "github.com/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	dedicatedkmssdk "github.com/alibabacloud-dkms-gcs-go-sdk/sdk"
	"github.com/alibabacloud-go/tea/tea"
	"io/ioutil"
)

// The aes encrypt context may be stored
type AesEncryptContext struct {
	KeyId          string
	Iv             []byte
	CiphertextBlob []byte
	// Use default algorithm value, if the value is not set
	Algorithm string
}

func main() {
	// 待加密明文
	plaintext := "yourPlaintext"
	// 专属KMS实例对称密钥的ID或别名（Alias）
	keyId := "<your symmetric key id>"

	// 创建DKMS Client对象
	client := getDkmsClientByClientKeyContent()
	//client := getDkmsClientByClientKeyFile()

	// 对称密钥加解密示例
	cipherCtx := encryptSample(client, []byte(plaintext), keyId)
	decryptResult := decryptSample(client, cipherCtx)
	fmt.Println(string(decryptResult))
}

// 对称加密示例
func encryptSample(client *dedicatedkmssdk.Client, plaintext []byte, keyId string) *AesEncryptContext {
	encryptRequest := &dedicatedkmssdk.EncryptRequest{
		KeyId:     tea.String(keyId),
		Plaintext: plaintext,
	}
	// 验证服务端证书
	ca, err := ioutil.ReadFile("path/to/caCert.pem")
	if err != nil {
		panic(err)
	}
	runtimeOptions := &dedicatedkmsopenapiutil.RuntimeOptions{
		Verify: tea.String(string(ca)),
	}
	// 或，忽略证书
	//runtimeOptions := &dedicatedkmsopenapiutil.RuntimeOptions{
	//	IgnoreSSL: tea.Bool(true),
	//}
	// 调用加密接口进行加密
	encryptResponse, err := client.EncryptWithOptions(encryptRequest, runtimeOptions)
	if err != nil {
		panic(err)
	}

	// 密钥ID
	_keyId := tea.StringValue(encryptResponse.KeyId)
	// 主密钥是对称密钥时，decrypt接口需要加密返回的Iv
	_iv := encryptResponse.Iv
	// 数据密文
	_cipher := encryptResponse.CiphertextBlob
	// 加密算法
	_algorithm := tea.StringValue(encryptResponse.Algorithm)

	fmt.Println("KeyId:", _keyId)
	fmt.Println("CiphertextBlob:", _cipher)
	fmt.Println("Iv:", _iv)
	fmt.Println("Algorithm:", _algorithm)
	fmt.Println("RequestId:", tea.StringValue(encryptResponse.RequestId))

	return &AesEncryptContext{
		KeyId:          _keyId,
		Iv:             _iv,
		CiphertextBlob: _cipher,
		Algorithm:      _algorithm,
	}
}

// 对称解密示例
func decryptSample(client *dedicatedkmssdk.Client, ctx *AesEncryptContext) []byte {
	decryptRequest := &dedicatedkmssdk.DecryptRequest{
		KeyId:          tea.String(ctx.KeyId),
		CiphertextBlob: ctx.CiphertextBlob, // 数据密文
		Iv:             ctx.Iv,             // 加密返回的Iv
		Algorithm:      tea.String(ctx.Algorithm),
	}
	// 验证服务端证书
	ca, err := ioutil.ReadFile("path/to/caCert.pem")
	if err != nil {
		panic(err)
	}
	runtimeOptions := &dedicatedkmsopenapiutil.RuntimeOptions{
		Verify: tea.String(string(ca)),
	}
	// 或，忽略证书
	//runtimeOptions := &dedicatedkmsopenapiutil.RuntimeOptions{
	//	IgnoreSSL: tea.Bool(true),
	//}
	// 调用解密接口进行解密
	decryptResponse, err := client.DecryptWithOptions(decryptRequest, runtimeOptions)
	if err != nil {
		panic(err)
	}

	// 数据明文
	_plaintext := decryptResponse.Plaintext

	fmt.Println("KeyId:", tea.StringValue(decryptResponse.KeyId))
	fmt.Println("Plaintext:", string(_plaintext))
	fmt.Println("RequestId:", tea.StringValue(decryptResponse.RequestId))

	return decryptResponse.Plaintext
}

// 使用ClientKey内容创建DKMS Client对象
func getDkmsClientByClientKeyContent() *dedicatedkmssdk.Client {
	// 创建DKMS Client配置
	config := &dedicatedkmsopenapi.Config{
		Protocol: tea.String("https"),
		// 请替换为您在KMS应用管理获取的ClientKey文件内容
		ClientKeyContent: tea.String("yourClientKeyContent"),
		// 请替换为您在KMS应用管理创建ClientKey时输入的加密口令
		Password: tea.String("yourClientKeyPassword"),
		// 请替换为您实际的专属KMS实例服务地址(不包括协议头https://)
		Endpoint: tea.String("yourEndpoint"),
	}
	// 创建DKMS Client对象
	client, err := dedicatedkmssdk.NewClient(config)
	if err != nil {
		// 异常处理
		panic(err)
	}
	return client
}

// 使用ClientKey文件路径创建DKMS Client对象
func getDkmsClientByClientKeyFile() *dedicatedkmssdk.Client {
	// 创建DKMS Client配置
	config := &dedicatedkmsopenapi.Config{
		Protocol: tea.String("https"),
		// 请替换为您在KMS应用管理获取的ClientKey文件的路径
		ClientKeyFile: tea.String("yourClientKeyFile"),
		// 请替换为您在KMS应用管理创建ClientKey时输入的加密口令
		Password: tea.String("yourClientKeyPassword"),
		// 请替换为您实际的专属KMS实例服务地址(不包括协议头https://)
		Endpoint: tea.String("yourEndpoint"),
	}
	// 创建DKMS Client对象
	client, err := dedicatedkmssdk.NewClient(config)
	if err != nil {
		// 异常处理
		panic(err)
	}
	return client
}

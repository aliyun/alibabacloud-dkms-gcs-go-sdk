package main

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	dedicatedkmsopenapi "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	dedicatedkmssdk "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/sdk"
	"io/ioutil"
)

// The encrypt context may be stored
type EncryptContext struct {
	KeyId          string
	Iv             []byte
	CiphertextBlob []byte
	// Use default algorithm value, if the value is not set
	Algorithm string
}

func main() {
	// 待加密明文
	plaintext := "yourPlaintext"
	// 专属KMS实例密钥的ID或别名（Alias）
	keyId := "<your key id>"

	// 创建DKMS Client对象
	client, err := getDkmsClientByClientKeyContent()
	//client := getDkmsClientByClientKeyFile()
	if err != nil {
		panic(err)
	}
	// 密钥加解密示例
	cipherCtx, err := encrypt(client, []byte(plaintext), keyId)
	if err != nil {
		panic(err)
	}
	decryptResult, err := advanceDecrypt(client, cipherCtx)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decryptResult))
}

// 加密示例
func encrypt(client *dedicatedkmssdk.Client, plaintext []byte, keyId string) (*EncryptContext, error) {
	encryptRequest := &dedicatedkmssdk.EncryptRequest{
		KeyId:     tea.String(keyId),
		Plaintext: plaintext,
	}
	// 验证服务端证书
	ca, err := ioutil.ReadFile("path/to/caCert.pem")
	if err != nil {
		return nil, err
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
		return nil, err
	}

	// 密钥ID
	_keyId := tea.StringValue(encryptResponse.KeyId)
	// Iv
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

	return &EncryptContext{
		KeyId:          _keyId,
		Iv:             _iv,
		CiphertextBlob: _cipher,
		Algorithm:      _algorithm,
	}, nil
}

// 高级解密示例
func advanceDecrypt(client *dedicatedkmssdk.Client, ctx *EncryptContext) ([]byte, error) {
	request := &dedicatedkmssdk.AdvanceDecryptRequest{
		KeyId:          tea.String(ctx.KeyId),
		CiphertextBlob: ctx.CiphertextBlob, // 数据密文
		Iv:             ctx.Iv,             // 加密返回的Iv
		Algorithm:      tea.String(ctx.Algorithm),
	}
	// 验证服务端证书
	ca, err := ioutil.ReadFile("path/to/caCert.pem")
	if err != nil {
		return nil, err
	}
	runtimeOptions := &dedicatedkmsopenapiutil.RuntimeOptions{
		Verify: tea.String(string(ca)),
	}
	// 或，忽略证书
	//runtimeOptions := &dedicatedkmsopenapiutil.RuntimeOptions{
	//	IgnoreSSL: tea.Bool(true),
	//}
	// 调用解密接口进行解密
	response, err := client.AdvanceDecryptWithOptions(request, runtimeOptions)
	if err != nil {
		return nil, err
	}
	// 数据明文
	_plaintext := response.Plaintext

	fmt.Println("KeyId:", tea.StringValue(response.KeyId))
	fmt.Println("Plaintext:", string(_plaintext))
	fmt.Println("RequestId:", tea.StringValue(response.RequestId))

	return response.Plaintext, nil
}

// 使用ClientKey内容创建DKMS Client对象
func getDkmsClientByClientKeyContent() (*dedicatedkmssdk.Client, error) {
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
	return dedicatedkmssdk.NewClient(config)
}

// 使用ClientKey文件路径创建DKMS Client对象
func getDkmsClientByClientKeyFile() (*dedicatedkmssdk.Client, error) {
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
	return dedicatedkmssdk.NewClient(config)
}

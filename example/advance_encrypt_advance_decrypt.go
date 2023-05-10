package main

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	dedicatedkmsopenapi "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	dedicatedkmssdk "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/sdk"
	"io/ioutil"
)

// The advance encrypt context may be stored
type AdvanceEncryptContext struct {
	CiphertextBlob []byte
}

func main() {
	// 待加密明文
	plaintext := "your plaintext"
	// 专属KMS实例密钥的ID或别名（Alias）
	keyId := "<your key id>"

	// 创建DKMS Client对象
	client, err := getDkmsClientByClientKeyContent()
	//client := getDkmsClientByClientKeyFile()
	if err != nil {
		panic(err)
	}
	// 高级密钥加解密示例
	cipherCtx, err := advanceEncryptSample(client, []byte(plaintext), keyId)
	if err != nil {
		panic(err)
	}
	result, err := advanceDecryptSample(client, cipherCtx)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

// 高级加密示例
func advanceEncryptSample(client *dedicatedkmssdk.Client, plaintext []byte, keyId string) (*AdvanceEncryptContext, error) {
	request := &dedicatedkmssdk.AdvanceEncryptRequest{
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
	response, err := client.AdvanceEncryptWithOptions(request, runtimeOptions)
	if err != nil {
		return nil, err
	}

	// 密钥ID
	_keyId := tea.StringValue(response.KeyId)
	// Iv
	_iv := response.Iv
	// 数据密文
	_cipher := response.CiphertextBlob
	// 加密算法
	_algorithm := tea.StringValue(response.Algorithm)

	fmt.Println("KeyId:", _keyId)
	fmt.Println("CiphertextBlob:", _cipher)
	fmt.Println("Iv:", _iv)
	fmt.Println("Algorithm:", _algorithm)
	fmt.Println("RequestId:", tea.StringValue(response.RequestId))

	return &AdvanceEncryptContext{
		CiphertextBlob: _cipher,
	}, nil
}

// 高级解密示例
func advanceDecryptSample(client *dedicatedkmssdk.Client, ctx *AdvanceEncryptContext) ([]byte, error) {
	request := &dedicatedkmssdk.AdvanceDecryptRequest{
		CiphertextBlob: ctx.CiphertextBlob, // 数据密文
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

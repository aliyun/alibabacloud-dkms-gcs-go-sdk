package main

import (
	"fmt"
	dedicatedkmsopenapi "github.com/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	dedicatedkmssdk "github.com/alibabacloud-dkms-gcs-go-sdk/sdk"
	"github.com/alibabacloud-go/tea/tea"
	"io/ioutil"
)

func main() {
	// 专属KMS实例对称密钥的ID或别名（Alias）
	keyId := "yourSymmetricKeyId"

	// 创建DKMS Client对象
	client := getDkmsClientByClientKeyContent()
	//client := getDkmsClientByClientKeyFile()

	generateDataKeyRequest := &dedicatedkmssdk.GenerateDataKeyRequest{
		KeyId:         tea.String(keyId),
		NumberOfBytes: tea.Int32(32), // 生成的数据密钥长度
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

	// 调用生成数据密钥接口
	response, err := client.GenerateDataKeyWithOptions(generateDataKeyRequest, runtimeOptions)
	if err != nil {
		panic(err)
	}

	// 密钥ID
	_keyId := tea.StringValue(response.KeyId)
	// 明文数据密钥
	_plaintext := response.Plaintext
	// 密文数据密钥
	_ciphertextBlob := response.CiphertextBlob
	// 由专属KMS生成的加密初始向量，解密数据密钥时需要传入
	_iv := response.Iv

	fmt.Println("KeyId:", _keyId)
	fmt.Println("Plaintext:", _plaintext)
	fmt.Println("CiphertextBlob:", _ciphertextBlob)
	fmt.Println("Iv:", _iv)
	fmt.Println("RequestId:", tea.StringValue(response.RequestId))
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

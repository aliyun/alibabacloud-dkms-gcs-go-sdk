package main

import (
	"fmt"
	dedicatedkmsopenapi "github.com/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	dedicatedkmssdk "github.com/alibabacloud-dkms-gcs-go-sdk/sdk"
	"github.com/alibabacloud-go/tea/tea"
	"io/ioutil"
)

// The signature context may be stored
type SignatureContext struct {
	KeyId     string
	Signature []byte
	// Use default algorithm value,if the value is not set
	Algorithm   string
	MessageType string
}

func main() {
	// 专属KMS实例签名密钥的ID或别名（Alias）
	keyId := "yourKeyId"
	// 待签名数据摘要或预处理数据
	//digest := sha256.Sum256([]byte("message"))
	message := "message"
	// 签名数据类型，RAW-待签名原始数据 DIGEST-签名数据的摘要
	//messageType := "DIGEST"
	messageType := "RAW"

	// 创建DKMS Client对象
	client := getDkmsClientByClientKeyContent()
	//client := getDkmsClientByClientKeyFile()

	signatureCtx := signSample(client, keyId, []byte(message), messageType)
	verifyResult := verifySample(client, []byte(message), signatureCtx)
	fmt.Println(verifyResult)
}

// 签名示例
func signSample(client *dedicatedkmssdk.Client, keyId string, message []byte, messageType string) *SignatureContext {
	signRequest := &dedicatedkmssdk.SignRequest{
		KeyId:       tea.String(keyId),
		Message:     message,
		MessageType: tea.String(messageType),
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
	// 调用签名接口进行签名
	signResponse, err := client.SignWithOptions(signRequest, runtimeOptions)
	if err != nil {
		panic(err)
	}

	// 密钥ID
	_keyId := tea.StringValue(signResponse.KeyId)
	// 签名值
	_signature := signResponse.Signature
	// 消息类型
	_messageType := tea.StringValue(signResponse.MessageType)
	// 签名算法
	_algorithm := tea.StringValue(signResponse.Algorithm)

	fmt.Println("KeyId:", _keyId)
	fmt.Println("Signature:", _signature)
	fmt.Println("MessageType:", _messageType)
	fmt.Println("RequestId:", tea.StringValue(signResponse.RequestId))

	return &SignatureContext{
		KeyId:       _keyId,
		Signature:   _signature,
		MessageType: messageType,
		Algorithm:   _algorithm,
	}
}

// 验签示例
func verifySample(client *dedicatedkmssdk.Client, message []byte, ctx *SignatureContext) (_value bool) {
	verifyRequest := &dedicatedkmssdk.VerifyRequest{
		KeyId:       tea.String(ctx.KeyId),
		Message:     message,
		MessageType: tea.String(ctx.MessageType),
		Signature:   ctx.Signature,
		Algorithm:   tea.String(ctx.Algorithm),
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
	// 调用验签接口进行验签
	verifyResponse, err := client.VerifyWithOptions(verifyRequest, runtimeOptions)
	if err != nil {
		panic(err)
	}

	// 验签结果
	_value = tea.BoolValue(verifyResponse.Value)
	// 消息类型
	_messageType := verifyResponse.MessageType

	fmt.Println("KeyId:", tea.StringValue(verifyResponse.KeyId))
	fmt.Println("Value:", _value)
	fmt.Println("MessageType:", tea.StringValue(_messageType))
	fmt.Println("RequestId:", tea.StringValue(verifyResponse.RequestId))

	return tea.BoolValue(verifyResponse.Value)
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

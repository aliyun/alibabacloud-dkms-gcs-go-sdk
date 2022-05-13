package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	dedicatedkmsopenapi "github.com/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	dedicatedkmssdk "github.com/alibabacloud-dkms-gcs-go-sdk/sdk"
	"github.com/alibabacloud-go/tea/tea"
	"io/ioutil"
	"math/rand"
)

const (
	GcmIvLength = 12
)

// The envelope cipher text may be stored
type EnvelopeCipherText struct {
	DataKeyIV        []byte
	EncryptedDataKey []byte
	Iv               []byte
	CipherText       []byte
}

func main() {
	// 本地待加密数据
	data := []byte("<your plaintext data>")
	// 专属KMS实例对称密钥的ID或别名（Alias）
	keyId := "<your symmetric key id or alias>"

	// 创建DKMS Client对象并初始化
	client := getDkmsClientByClientKeyContent()
	//client := getDkmsClientByClientKeyFile()

	// 信封加解密示例
	envelopeCipherText := envelopeEncryptSample(client, keyId, data)
	decryptedData := envelopeDecryptSample(client, keyId, envelopeCipherText)
	fmt.Println(string(decryptedData))
}

// 信封加密示例
func envelopeEncryptSample(client *dedicatedkmssdk.Client, keyId string, data []byte) *EnvelopeCipherText {
	// 获取数据密钥，下面以Aliyun_AES_256密钥为例进行说明，数据密钥长度32字节
	generateDataKeyRequest := &dedicatedkmssdk.GenerateDataKeyRequest{
		KeyId:         tea.String(keyId),
		NumberOfBytes: tea.Int32(32),
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
	generateDataKeyResponse, err := client.GenerateDataKeyWithOptions(generateDataKeyRequest, runtimeOptions)
	if err != nil {
		panic(err)
	}

	// 专属KMS返回的数据密钥明文, 加密本地数据使用
	plainDataKey := generateDataKeyResponse.Plaintext
	// 专属KMS返回的数据密钥密文，解密本地数据密文时，先将数据密钥密文解密后使用
	encryptedDataKey := generateDataKeyResponse.CiphertextBlob
	// 由专属KMS生成的加密初始向量，解密数据密钥时需要传入
	dataKeyIv := generateDataKeyResponse.Iv

	// 使用专属KMS返回的数据密钥明文在本地对数据进行加密，下面以AES-256 GCM模式为例
	iv := make([]byte, GcmIvLength) // 加密初始向量，解密时需要传入
	rand.Read(iv)
	block, err := aes.NewCipher(plainDataKey)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	cipherText := gcm.Seal(nil, iv, data, nil)

	// 输出密文，密文输出或持久化由用户根据需要进行处理，下面示例仅展示将密文输出到一个对象的情况
	// 假如envelopeCipherText是需要输出的密文对象，至少需要包括以下四个内容:
	// (1) dataKeyIV: 由专属KMS生成的加密初始向量，解密数据密钥密文时需要传入
	// (2) encryptedDataKey: 专属KMS返回的数据密钥密文
	// (3) iv: 加密初始向量
	// (4) cipherText: 密文数据
	envelopeCipherText := &EnvelopeCipherText{
		DataKeyIV:        dataKeyIv,
		EncryptedDataKey: encryptedDataKey,
		Iv:               iv,
		CipherText:       cipherText,
	}
	return envelopeCipherText
}

// 信封解密示例
func envelopeDecryptSample(client *dedicatedkmssdk.Client, keyId string, envelopeCipherText *EnvelopeCipherText) []byte {
	// 由专属KMS生成的加密初始向量，解密数据密钥需要传入
	dataKeyIv := envelopeCipherText.DataKeyIV
	// 待解密数据密钥密文，由专属KMS生成
	encryptedDataKey := envelopeCipherText.EncryptedDataKey

	// 解密数据密钥密文，得到数据密钥明文
	decryptRequest := &dedicatedkmssdk.DecryptRequest{
		KeyId:          tea.String(keyId),
		CiphertextBlob: encryptedDataKey,
		Iv:             dataKeyIv,
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

	// 数据密钥明文
	plainDataKey := decryptResponse.Plaintext

	// 使用数据密钥明文在本地进行解密, 下面是以AES-256 GCM模式为例
	iv := envelopeCipherText.Iv                 // 本地加密时使用的初始向量, 解密数据需要传入
	cipherText := envelopeCipherText.CipherText // 待解密密文
	block, err := aes.NewCipher(plainDataKey)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	decryptedData, err := gcm.Open(nil, iv, cipherText, nil)
	if err != nil {
		panic(err)
	}
	return decryptedData
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

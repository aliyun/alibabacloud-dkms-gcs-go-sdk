package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	dedicatedkmsopenapi "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmssdk "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/sdk"
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

	// 创建DKMS Client对象并初始化
	client, err := initClient()
	if err != nil {
		panic(err)
	}
	// 高级接口信封加密示例
	err = envelopeAdvanceEncryptSample(client)
	if err != nil {
		panic(err)
	}
	// 高级接口信封解密示例
	decryptedData, err := envelopeAdvanceDecryptSample(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decryptedData))
}

// 高级接口信封加密示例
func envelopeAdvanceEncryptSample(client *dedicatedkmssdk.Client) error {
	// 本地待加密数据
	data := []byte("<your plaintext data>")
	// KMS实例对称密钥的ID或别名（Alias）
	keyId := "<your symmetric key id or alias>"
	// 获取数据密钥，下面以Aliyun_AES_256密钥为例进行说明，数据密钥长度32字节
	generateDataKeyRequest := &dedicatedkmssdk.AdvanceGenerateDataKeyRequest{
		KeyId:         tea.String(keyId),
		NumberOfBytes: tea.Int32(32),
	}

	// 调用生成数据密钥接口
	generateDataKeyResponse, err := client.AdvanceGenerateDataKey(generateDataKeyRequest)
	if err != nil {
		return err
	}

	// KMS实例返回的数据密钥明文, 加密本地数据使用
	plainDataKey := generateDataKeyResponse.Plaintext
	// KMS实例返回的数据密钥密文，解密本地数据密文时，先将数据密钥密文解密后使用
	encryptedDataKey := generateDataKeyResponse.CiphertextBlob
	// 由KMS实例生成的加密初始向量，解密数据密钥时需要传入
	dataKeyIv := generateDataKeyResponse.Iv

	// 使用KMS实例返回的数据密钥明文在本地对数据进行加密，下面以AES-256 GCM模式为例
	iv := make([]byte, GcmIvLength) // 加密初始向量，解密时需要传入
	rand.Read(iv)
	block, err := aes.NewCipher(plainDataKey)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	cipherText := gcm.Seal(nil, iv, data, nil)

	// 输出密文，密文输出或持久化由用户根据需要进行处理，下面示例仅展示将密文输出到一个对象的情况
	// 假如envelopeCipherText是需要输出的密文对象，至少需要包括以下四个内容:
	envelopeCipherText := &EnvelopeCipherText{
		// (1) dataKeyIV: 由KMS实例生成的加密初始向量，解密数据密钥密文时需要传入
		DataKeyIV: dataKeyIv,
		// (2) encryptedDataKey: KMS实例返回的数据密钥密文
		EncryptedDataKey: encryptedDataKey,
		// (3) iv: 加密初始向量
		Iv: iv,
		// (4) cipherText: 密文数据
		CipherText: cipherText,
	}
	// 保存信封密文持久化对象
	saveEnvelopeCipherPersistObject(envelopeCipherText)
	return nil
}

func saveEnvelopeCipherPersistObject(envelopeCipherText *EnvelopeCipherText) {
	// TODO 用户自行保存输出的密文对象
}

// 高级接口信封解密示例
func envelopeAdvanceDecryptSample(client *dedicatedkmssdk.Client) ([]byte, error) {
	// 从存储中读取封信加密持久化对象
	envelopeCipherText := getEnvelopeCipherPersistObject()
	// 由KMS实例生成的加密初始向量，解密数据密钥需要传入
	dataKeyIv := envelopeCipherText.DataKeyIV
	// 待解密数据密钥密文，由KMS实例生成
	encryptedDataKey := envelopeCipherText.EncryptedDataKey

	// 解密数据密钥密文，得到数据密钥明文
	decryptRequest := &dedicatedkmssdk.AdvanceDecryptRequest{
		CiphertextBlob: encryptedDataKey,
		Iv:             dataKeyIv,
	}

	// 调用解密接口进行解密
	decryptResponse, err := client.AdvanceDecrypt(decryptRequest)
	if err != nil {
		return nil, err
	}

	// 数据密钥明文
	plainDataKey := decryptResponse.Plaintext

	// 使用数据密钥明文在本地进行解密, 下面是以AES-256 GCM模式为例
	iv := envelopeCipherText.Iv                 // 本地加密时使用的初始向量, 解密数据需要传入
	cipherText := envelopeCipherText.CipherText // 待解密密文
	block, err := aes.NewCipher(plainDataKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	decryptedData, err := gcm.Open(nil, iv, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

func getEnvelopeCipherPersistObject() *EnvelopeCipherText {
	// TODO 用户需要在此处代码进行替换，从存储中读取封信加密持久化对象
	return &EnvelopeCipherText{}
}

func initClient() (*dedicatedkmssdk.Client, error) {
	// 构建KMS实例Client配置
	config := &dedicatedkmsopenapi.Config{
		Protocol: tea.String("https"),
		// 请替换为您在KMS应用管理获取的ClientKey文件的路径
		ClientKeyFile: tea.String("yourClientKeyFile"),
		// 请替换为您在KMS应用管理创建ClientKey时输入的加密口令
		Password: tea.String("yourClientKeyPassword"),
		// 请替换为您实际的KMS实例服务地址(不包括协议头https://)
		Endpoint: tea.String("yourEndpoint"),
		// 验证服务端证书，这里需要设置为您的服务端证书路径
		CaFilePath: tea.String("yourCaFilePath"),
	}
	// 创建Client对象
	client, err := dedicatedkmssdk.NewClient(config)
	if err != nil {
		// 异常处理
		return nil, err
	}
	return client, nil
}

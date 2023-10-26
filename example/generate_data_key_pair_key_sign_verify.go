package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	dedicatedkmsopenapi "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmssdk "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/sdk"
)

const (
	KEY_FORMAT_DER = "DER"
	KEY_FORMAT_PEM = "PEM"
)

func main() {
	// 创建DKMS Client对象
	client, err := initClient()
	if err != nil {
		// 异常处理
		panic(err)
	}
	//调用GenerateDataKeyPair接口生成公私钥
	keyFormat := tea.String(KEY_FORMAT_DER)
	generateDataKeyPairResponse, err := generateDataKeyPair(client, keyFormat)
	if err != nil {
		// 异常处理
		panic(err)
	}
	//保存签名以及密钥等信息
	keyPairInfo := saveKeyPairInfo(generateDataKeyPairResponse, keyFormat)
	//签名
	privateKeyCiphertextBlob, err := base64.StdEncoding.DecodeString(tea.StringValue(keyPairInfo.PrivateKeyCiphertextBlob))
	if err != nil {
		// 异常处理
		panic(err)
	}
	//解密私钥密文
	iv := keyPairInfo.Iv
	privateKeyPlaintext, err := decrypt(client, keyPairInfo.KeyId, privateKeyCiphertextBlob, iv)
	if err != nil {
		// 异常处理
		panic(err)
	}
	message := []byte("your message")
	privateKey := privateKeyPlaintext
	if tea.StringValue(keyPairInfo.KeyFormat) == KEY_FORMAT_PEM {
		privateKey, err = decodePem(privateKeyPlaintext)
		if err != nil {
			// 异常处理
			panic(err)
		}
	}
	signature, err := sign(message, privateKey)
	if err != nil {
		// 异常处理
		panic(err)
	}

	//验签
	publicKey, err := base64.StdEncoding.DecodeString(tea.StringValue(keyPairInfo.PublicKey))
	if err != nil {
		// 异常处理
		panic(err)
	}
	if tea.StringValue(keyPairInfo.KeyFormat) == KEY_FORMAT_PEM {
		publicKey, err = decodePem(publicKey)
		if err != nil {
			// 异常处理
			panic(err)
		}
	}
	err = verify(message, signature, publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(true)
}

func saveKeyPairInfo(response *dedicatedkmssdk.GenerateDataKeyPairResponse, keyFormat *string) *KeyPairInfo {
	keyPairInfo := &KeyPairInfo{
		KeyId:                    response.KeyId,
		KeyFormat:                keyFormat,
		KeyPairSpec:              response.KeyPairSpec,
		PrivateKeyCiphertextBlob: tea.String(base64.StdEncoding.EncodeToString(response.PrivateKeyCiphertextBlob)),
		PublicKey:                tea.String(base64.StdEncoding.EncodeToString(response.PublicKey)),
		Iv:                       response.Iv,
		Algorithm:                response.Algorithm,
	}
	// TODO 此处可以持久化公钥及私钥密文等信息
	return keyPairInfo
}

// 创建DKMS Client对象
func initClient() (*dedicatedkmssdk.Client, error) {
	// 创建DKMS Client配置
	config := &dedicatedkmsopenapi.Config{
		Protocol: tea.String("https"),
		// 请替换为您在KMS应用管理获取的ClientKey文件的路径
		ClientKeyFile: tea.String("your clientKeyFile"),
		// 请替换为您在KMS应用管理创建ClientKey时输入的加密口令
		Password: tea.String("your clientKeyPassword"),
		// 请替换为您实际的KMS实例服务地址(不包括协议头https://)
		Endpoint: tea.String("your endpoint"),
		// CA证书文件路径
		CaFilePath: tea.String("your caFilePath"),
	}
	// 创建DKMS Client对象
	client, err := dedicatedkmssdk.NewClient(config)
	return client, err
}

// generateDataKeyPair 生成密钥对
func generateDataKeyPair(client *dedicatedkmssdk.Client, keyFormat *string) (_result *dedicatedkmssdk.GenerateDataKeyPairResponse, _err error) {
	keyId := tea.String("your keyId")
	keyPairSpec := tea.String("your keyPairSpec")
	request := &dedicatedkmssdk.GenerateDataKeyPairRequest{
		KeyFormat:   keyFormat,
		KeyId:       keyId,
		KeyPairSpec: keyPairSpec,
	}
	_result = &dedicatedkmssdk.GenerateDataKeyPairResponse{}
	return client.GenerateDataKeyPair(request)
}

// decrypt 解密
func decrypt(client *dedicatedkmssdk.Client, keyId *string, ciphertextBlob []byte, iv []byte) ([]byte, error) {
	decryptRequest := &dedicatedkmssdk.DecryptRequest{
		KeyId:          keyId,
		CiphertextBlob: ciphertextBlob, // 数据密文
		Iv:             iv,             // 加密返回的Iv
	}
	// 调用解密接口进行解密
	decryptResponse, err := client.Decrypt(decryptRequest)
	return decryptResponse.Plaintext, err
}

func decodePem(keyBytes []byte) ([]byte, error) {
	// 解码
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("解码pem错误！")
	}
	return block.Bytes, nil
}

// sign 私钥签名
func sign(data []byte, privateKeyBytes []byte) (string, error) {
	// 1、选择hash算法，对需要签名的数据进行hash运算
	algorithm := crypto.SHA256
	hashed := hash(data, algorithm)
	// 2.生成私钥对象
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	if err != nil {
		return "", err
	}
	// 3、RSA数字签名（参数是随机数、私钥对象、哈希类型、签名文件的哈希串，生成bash64编码）
	bytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), algorithm, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// verify 公钥验证
func verify(data []byte, base64Sig string, publicKeyBytes []byte) error {
	// 1、对base64编码的签名内容进行解码，返回签名字节
	bytes, err := base64.StdEncoding.DecodeString(base64Sig)
	if err != nil {
		return err
	}
	// 2、选择hash算法，对需要签名的数据进行hash运算
	algorithm := crypto.SHA256
	hashed := hash(data, algorithm)

	// 3、解析DER编码的公钥，生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return err
	}
	// 4、公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	if err != nil {
		return err
	}
	// 5、RSA验证数字签名（参数是公钥对象、哈希类型、签名文件的哈希串、签名后的字节）
	return rsa.VerifyPKCS1v15(publicKey, algorithm, hashed, bytes)
}

func hash(data []byte, algorithm crypto.Hash) []byte {
	hashInstance := algorithm.New()
	hashInstance.Write(data)
	return hashInstance.Sum(nil)
}

type KeyPairInfo struct {
	KeyId                    *string
	Iv                       []byte
	KeyPairSpec              *string
	PrivateKeyCiphertextBlob *string
	PublicKey                *string
	Algorithm                *string
	KeyFormat                *string
}

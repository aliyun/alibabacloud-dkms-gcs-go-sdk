[English](README.md) | 简体中文

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

## AlibabaCloud DKMS-GCS SDK for Go

- [概述](https://help.aliyun.com/document_detail/311016.html)
- [入门指南](https://help.aliyun.com/document_detail/311368.html)
- [代码示例](/example)

## 依赖

- Go 1.12及以上。

## 安装

您可以使用`go mod`管理您的依赖：

```
require (
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.1.0
	github.com/alibabacloud-go/tea v1.1.17
)
```

或者，通过`go get`命令获取远程代码包：

```
$ go get -u github.com/aliyun/alibabacloud-dkms-gcs-go-sdk
```

## 快速使用

```go
package example

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	dedicatedkmsopenapi "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-util"
	dedicatedkmssdk "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/sdk"
)

func main() {
	config := &dedicatedkmsopenapi.Config{
		Protocol:         tea.String("https"),
		ClientKeyContent: tea.String("<your client key content>"),
		Password:         tea.String("<your client key password>"),
		Endpoint:         tea.String("<your dkms instance service endpoint>"),
	}
	client, err := dedicatedkmssdk.NewClient(config)
	if err != nil {
		panic(err)
	}
	runtimeOptions := &dedicatedkmsopenapiutil.RuntimeOptions{
		IgnoreSSL: tea.Bool(true),
	}
	encryptRequest := &dedicatedkmssdk.EncryptRequest{
		KeyId:     tea.String("<your key id>"),
		Plaintext: []byte("plaintext"),
	}
	encryptResponse, err := client.EncryptWithOptions(encryptRequest, runtimeOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println(encryptResponse)
}
```

## 许可证

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

版权所有 2009-present, 阿里巴巴集团.

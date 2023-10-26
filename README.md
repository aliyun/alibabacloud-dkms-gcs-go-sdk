English | [简体中文](README-CN.md)

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

## AlibabaCloud DKMS-GCS SDK for Go

- [Overview](https://www.alibabacloud.com/help/doc-detail/311016.htm)
- [Quick Start](https://www.alibabacloud.com/help/doc-detail/311368.htm)
- [Sample Code](/example)

## Requirements

- Go 1.13 or later.

## Installation

If you use `go mod` to manage your dependence, You can declare the dependency on AlibabaCloud DKMS SDK for Go in the
go.mod file:

```text
require (
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.5.0
	github.com/alibabacloud-go/tea v1.1.17
)
```

Or, Run the following command to get the remote code package:

```shell
$ go get -u github.com/aliyun/alibabacloud-dkms-gcs-go-sdk
```

## Quick Examples

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

## License

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.

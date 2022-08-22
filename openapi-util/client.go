package openapi_util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-util/protobuf/api"
	"github.com/golang/protobuf/proto"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type RuntimeOptions struct {
	Autoretry      *bool     `json:"autoretry,omitempty" xml:"autoretry,omitempty"`
	IgnoreSSL      *bool     `json:"ignoreSSL,omitempty" xml:"ignoreSSL,omitempty"`
	MaxAttempts    *int      `json:"maxAttempts,omitempty" xml:"maxAttempts,omitempty"`
	BackoffPolicy  *string   `json:"backoffPolicy,omitempty" xml:"backoffPolicy,omitempty"`
	BackoffPeriod  *int      `json:"backoffPeriod,omitempty" xml:"backoffPeriod,omitempty"`
	ReadTimeout    *int      `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	ConnectTimeout *int      `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	HttpProxy      *string   `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	HttpsProxy     *string   `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	NoProxy        *string   `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	MaxIdleConns   *int      `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	Socks5Proxy    *string   `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	Socks5NetWork  *string   `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	Verify         *string   `json:"verify,omitempty" xml:"verify,omitempty"`
	Headers        []*string `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
}

func (s RuntimeOptions) String() string {
	return tea.Prettify(s)
}

func (s RuntimeOptions) GoString() string {
	return s.String()
}

func (s *RuntimeOptions) SetAutoretry(v bool) *RuntimeOptions {
	s.Autoretry = &v
	return s
}

func (s *RuntimeOptions) SetIgnoreSSL(v bool) *RuntimeOptions {
	s.IgnoreSSL = &v
	return s
}

func (s *RuntimeOptions) SetMaxAttempts(v int) *RuntimeOptions {
	s.MaxAttempts = &v
	return s
}

func (s *RuntimeOptions) SetBackoffPolicy(v string) *RuntimeOptions {
	s.BackoffPolicy = &v
	return s
}

func (s *RuntimeOptions) SetBackoffPeriod(v int) *RuntimeOptions {
	s.BackoffPeriod = &v
	return s
}

func (s *RuntimeOptions) SetReadTimeout(v int) *RuntimeOptions {
	s.ReadTimeout = &v
	return s
}

func (s *RuntimeOptions) SetConnectTimeout(v int) *RuntimeOptions {
	s.ConnectTimeout = &v
	return s
}

func (s *RuntimeOptions) SetHttpProxy(v string) *RuntimeOptions {
	s.HttpProxy = &v
	return s
}

func (s *RuntimeOptions) SetHttpsProxy(v string) *RuntimeOptions {
	s.HttpsProxy = &v
	return s
}

func (s *RuntimeOptions) SetNoProxy(v string) *RuntimeOptions {
	s.NoProxy = &v
	return s
}

func (s *RuntimeOptions) SetMaxIdleConns(v int) *RuntimeOptions {
	s.MaxIdleConns = &v
	return s
}

func (s *RuntimeOptions) SetSocks5Proxy(v string) *RuntimeOptions {
	s.Socks5Proxy = &v
	return s
}

func (s *RuntimeOptions) SetSocks5NetWork(v string) *RuntimeOptions {
	s.Socks5NetWork = &v
	return s
}

func (s *RuntimeOptions) SetVerify(v string) *RuntimeOptions {
	s.Verify = &v
	return s
}

func (s *RuntimeOptions) SetHeaders(v []*string) *RuntimeOptions {
	s.Headers = v
	return s
}

type ErrorResponse struct {
	StatusCode   *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty" require:"true"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s ErrorResponse) GoString() string {
	return s.String()
}

func (s *ErrorResponse) SetStatusCode(v string) *ErrorResponse {
	s.StatusCode = &v
	return s
}

func (s *ErrorResponse) SetErrorCode(v string) *ErrorResponse {
	s.ErrorCode = &v
	return s
}

func (s *ErrorResponse) SetErrorMessage(v string) *ErrorResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ErrorResponse) SetRequestId(v string) *ErrorResponse {
	s.RequestId = &v
	return s
}

func FilterHeaders(headers map[string]*string, constraint []*string) (_result map[string]*string) {
	if constraint == nil {
		return
	}
	_result = make(map[string]*string)
	for _, key := range constraint {
		if value, ok := headers[tea.StringValue(key)]; ok {
			_result[tea.StringValue(key)] = value
		}
	}
	return
}

func GetHost(regionId *string, endpoint *string) (_result *string) {
	_result = endpoint
	return
}

func GetErrMessage(msg []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.Error{}
	err = proto.Unmarshal(msg, response)
	if err != nil {
		err = errors.New(fmt.Sprintf("proto.Unmarshal(%s), err:%v", string(msg), err))
		return
	}
	_result["Code"] = response.ErrorCode
	_result["Message"] = response.ErrorMessage
	_result["RequestId"] = response.RequestId
	return
}

func ConvertToMap(body interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	val := reflect.ValueOf(body).Elem()
	dataType := val.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, _ := field.Tag.Lookup("json")
		name = strings.Split(name, ",omitempty")[0]
		_, ok := val.Field(i).Interface().(io.Reader)
		if !ok {
			res[name] = val.Field(i).Interface()
		}
	}
	return res
}

func GetStringToSign(request *tea.Request) (_result *string) {
	return tea.String(getStringToSign(request))
}

func getStringToSign(req *tea.Request) string {
	var canoHeaders string
	var dkmsHeaderKeys sort.StringSlice

	contentSHA256 := tea.StringValue(req.Headers["content-sha256"])
	contentType := tea.StringValue(req.Headers["content-type"])
	date := tea.StringValue(req.Headers["date"])

	// Calc CanonicalizedDKMSHeaders
	dkmsHeaders := make(map[string]string, len(req.Headers))
	for k, _ := range req.Headers {
		l := strings.TrimSpace(strings.ToLower(k))
		if strings.HasPrefix(l, "x-kms-") {
			dkmsHeaders[l] = strings.TrimSpace(tea.StringValue(req.Headers[k]))
			dkmsHeaderKeys = append(dkmsHeaderKeys, l)
		}
	}

	sort.Sort(dkmsHeaderKeys)
	for i, k := range dkmsHeaderKeys {
		canoHeaders += k + ":" + dkmsHeaders[k]
		if i+1 < len(dkmsHeaderKeys) {
			canoHeaders += "\n"
		}
	}

	// Calc CanonicalizedResource
	canoResource := tea.StringValue(req.Pathname)
	if len(req.Query) > 0 {
		queryParams := req.Query
		var keys sort.StringSlice
		for key := range queryParams {
			keys = append(keys, key)
		}
		sort.Sort(keys)
		canoResource += "?"
		for i, k := range keys {
			if i > 0 {
				canoResource += "&"
			}
			v := tea.StringValue(queryParams[k])
			if v != "" {
				canoResource += k + "=" + v
			} else {
				canoResource += k
			}
		}
	}

	return tea.StringValue(req.Method) + "\n" +
		contentSHA256 + "\n" +
		contentType + "\n" +
		date + "\n" +
		canoHeaders + "\n" +
		canoResource
}

func GetContentLength(reqBody []byte) (_result *string) {
	return tea.String(strconv.Itoa(len(reqBody)))
}

func GetContentSHA256(reqBody []byte) (_result *string) {
	bodyHashed := sha256.Sum256(reqBody)
	return tea.String(strings.ToUpper(hex.EncodeToString(bodyHashed[:])))
}

func ToHexString(byteArray []byte) (_result *string) {
	return tea.String(hex.EncodeToString(byteArray))
}

func GetSerializedEncryptRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.EncryptRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Plaintext"]; ok {
		request.Plaintext = v.([]byte)
	}
	if v, ok := reqBody["Algorithm"]; ok {
		request.Algorithm = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Iv"]; ok {
		request.Iv = v.([]byte)
	}
	if v, ok := reqBody["Aad"]; ok {
		request.Aad = v.([]byte)
	}
	if v, ok := reqBody["PaddingMode"]; ok {
		request.PaddingMode = tea.StringValue(v.(*string))
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseEncryptResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.EncryptResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["CiphertextBlob"] = response.CiphertextBlob
	_result["Iv"] = response.Iv
	_result["Algorithm"] = tea.String(response.Algorithm)
	_result["PaddingMode"] = tea.String(response.PaddingMode)
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedDecryptRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.DecryptRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["CiphertextBlob"]; ok {
		request.CiphertextBlob = v.([]byte)
	}
	if v, ok := reqBody["Algorithm"]; ok {
		request.Algorithm = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Aad"]; ok {
		request.Aad = v.([]byte)
	}
	if v, ok := reqBody["Iv"]; ok {
		request.Iv = v.([]byte)
	}
	if v, ok := reqBody["PaddingMode"]; ok {
		request.PaddingMode = tea.StringValue(v.(*string))
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseDecryptResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.DecryptResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["Plaintext"] = response.Plaintext
	_result["Algorithm"] = tea.String(response.Algorithm)
	_result["PaddingMode"] = tea.String(response.PaddingMode)
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedHmacRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.HmacRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Message"]; ok {
		request.Message = v.([]byte)
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseHmacResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.HmacResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["Signature"] = response.Signature
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedSignRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.SignRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Algorithm"]; ok {
		request.Algorithm = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Digest"]; ok {
		request.Digest = v.([]byte)
	}
	if v, ok := reqBody["Message"]; ok {
		request.Message = v.([]byte)
	}
	if v, ok := reqBody["MessageType"]; ok {
		request.MessageType = tea.StringValue(v.(*string))
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseSignResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.SignResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["Signature"] = response.Signature
	_result["Algorithm"] = tea.String(response.Algorithm)
	_result["MessageType"] = tea.String(response.MessageType)
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedVerifyRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.VerifyRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Algorithm"]; ok {
		request.Algorithm = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Signature"]; ok {
		request.Signature = v.([]byte)
	}
	if v, ok := reqBody["Digest"]; ok {
		request.Digest = v.([]byte)
	}
	if v, ok := reqBody["Message"]; ok {
		request.Message = v.([]byte)
	}
	if v, ok := reqBody["MessageType"]; ok {
		request.MessageType = tea.StringValue(v.(*string))
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseVerifyResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.VerifyResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["Value"] = tea.Bool(response.Value)
	_result["Algorithm"] = tea.String(response.Algorithm)
	_result["MessageType"] = tea.String(response.MessageType)
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedGenerateRandomRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.GenerateRandomRequest{}
	if v, ok := reqBody["Length"]; ok {
		request.Length = tea.Int32Value(v.(*int32))
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseGenerateRandomResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.GenerateRandomResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["Random"] = response.Random
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedGenerateDataKeyRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.GenerateDataKeyRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Algorithm"]; ok {
		request.Algorithm = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["NumberOfBytes"]; ok {
		request.NumberOfBytes = tea.Int32Value(v.(*int32))
	}
	if v, ok := reqBody["Aad"]; ok {
		request.Aad = v.([]byte)
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseGenerateDataKeyResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.GenerateDataKeyResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["Iv"] = response.Iv
	_result["Plaintext"] = response.Plaintext
	_result["CiphertextBlob"] = response.CiphertextBlob
	_result["Algorithm"] = tea.String(response.Algorithm)
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedGetPublicKeyRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.GetPublicKeyRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseGetPublicKeyResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.GetPublicKeyResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["PublicKey"] = tea.String(response.PublicKey)
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedHashRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.HashRequest{}
	if v, ok := reqBody["Algorithm"]; ok {
		request.Algorithm = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Message"]; ok {
		request.Message = v.([]byte)
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseHashResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.HashResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["Digest"] = response.Digest
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedKmsEncryptRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.KmsEncryptRequest{}
	if v, ok := reqBody["KeyId"]; ok {
		request.KeyId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["Plaintext"]; ok {
		request.Plaintext = v.([]byte)
	}
	if v, ok := reqBody["Aad"]; ok {
		request.Aad = v.([]byte)
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseKmsEncryptResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.KmsEncryptResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["CiphertextBlob"] = response.CiphertextBlob
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedKmsDecryptRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.KmsDecryptRequest{}
	if v, ok := reqBody["CiphertextBlob"]; ok {
		request.CiphertextBlob = v.([]byte)
	}
	if v, ok := reqBody["Aad"]; ok {
		request.Aad = v.([]byte)
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseKmsDecryptResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.KmsDecryptResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	_result["KeyId"] = tea.String(response.KeyId)
	_result["Plaintext"] = response.Plaintext
	_result["RequestId"] = tea.String(response.RequestId)
	return
}

func GetSerializedGetSecretValueRequest(reqBody map[string]interface{}) (_result []byte, err error) {
	request := &api.GetSecretValueRequest{}
	if v, ok := reqBody["SecretName"]; ok {
		request.SecretName = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["VersionStage"]; ok {
		request.VersionStage = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["VersionId"]; ok {
		request.VersionId = tea.StringValue(v.(*string))
	}
	if v, ok := reqBody["FetchExtendedConfig"]; ok {
		request.FetchExtendedConfig = tea.BoolValue(v.(*bool))
	}
	_result, err = proto.Marshal(request)
	return
}

func ParseGetSecretValueResponse(resBody []byte) (_result map[string]interface{}, err error) {
	_result = make(map[string]interface{})
	response := &api.GetSecretValueResponse{}
	err = proto.Unmarshal(resBody, response)
	if err != nil {
		return
	}
	var versionStages []*string
	for _, stage := range response.VersionStages {
		versionStages = append(versionStages, tea.String(stage))
	}
	_result["SecretName"] = tea.String(response.SecretName)
	_result["SecretType"] = tea.String(response.SecretType)
	_result["SecretData"] = tea.String(response.SecretData)
	_result["SecretDataType"] = tea.String(response.SecretDataType)
	_result["VersionStages"] = versionStages
	_result["VersionId"] = tea.String(response.VersionId)
	_result["CreateTime"] = tea.String(response.CreateTime)
	_result["RequestId"] = tea.String(response.RequestId)
	_result["LastRotationDate"] = tea.String(response.LastRotationDate)
	_result["NextRotationDate"] = tea.String(response.NextRotationDate)
	_result["ExtendedConfig"] = tea.String(response.ExtendedConfig)
	_result["AutomaticRotation"] = tea.String(response.AutomaticRotation)
	_result["RotationInterval"] = tea.String(response.RotationInterval)
	return
}

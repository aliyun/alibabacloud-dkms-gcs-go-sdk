// This file is auto-generated, don't edit it. Thanks.
package sdk

import (
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	dedicatedkmsopenapi "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi"
	dedicatedkmsopenapiutil "github.com/aliyun/alibabacloud-dkms-gcs-go-sdk/openapi-util"
)

type EncryptRequest struct {
	Headers     map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId       *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Plaintext   []byte             `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	Algorithm   *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Aad         []byte             `json:"Aad,omitempty" xml:"Aad,omitempty"`
	Iv          []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	PaddingMode *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
}

func (s EncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptRequest) GoString() string {
	return s.String()
}

func (s *EncryptRequest) SetHeaders(v map[string]*string) *EncryptRequest {
	s.Headers = v
	return s
}

func (s *EncryptRequest) SetKeyId(v string) *EncryptRequest {
	s.KeyId = &v
	return s
}

func (s *EncryptRequest) SetPlaintext(v []byte) *EncryptRequest {
	s.Plaintext = v
	return s
}

func (s *EncryptRequest) SetAlgorithm(v string) *EncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *EncryptRequest) SetAad(v []byte) *EncryptRequest {
	s.Aad = v
	return s
}

func (s *EncryptRequest) SetIv(v []byte) *EncryptRequest {
	s.Iv = v
	return s
}

func (s *EncryptRequest) SetPaddingMode(v string) *EncryptRequest {
	s.PaddingMode = &v
	return s
}

type EncryptResponse struct {
	Headers        map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId          *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	CiphertextBlob []byte             `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	Iv             []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	Algorithm      *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	PaddingMode    *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
	RequestId      *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s EncryptResponse) GoString() string {
	return s.String()
}

func (s *EncryptResponse) SetHeaders(v map[string]*string) *EncryptResponse {
	s.Headers = v
	return s
}

func (s *EncryptResponse) SetKeyId(v string) *EncryptResponse {
	s.KeyId = &v
	return s
}

func (s *EncryptResponse) SetCiphertextBlob(v []byte) *EncryptResponse {
	s.CiphertextBlob = v
	return s
}

func (s *EncryptResponse) SetIv(v []byte) *EncryptResponse {
	s.Iv = v
	return s
}

func (s *EncryptResponse) SetAlgorithm(v string) *EncryptResponse {
	s.Algorithm = &v
	return s
}

func (s *EncryptResponse) SetPaddingMode(v string) *EncryptResponse {
	s.PaddingMode = &v
	return s
}

func (s *EncryptResponse) SetRequestId(v string) *EncryptResponse {
	s.RequestId = &v
	return s
}

type DecryptRequest struct {
	Headers        map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	CiphertextBlob []byte             `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	KeyId          *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Algorithm      *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Aad            []byte             `json:"Aad,omitempty" xml:"Aad,omitempty"`
	Iv             []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	PaddingMode    *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
}

func (s DecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s DecryptRequest) GoString() string {
	return s.String()
}

func (s *DecryptRequest) SetHeaders(v map[string]*string) *DecryptRequest {
	s.Headers = v
	return s
}

func (s *DecryptRequest) SetCiphertextBlob(v []byte) *DecryptRequest {
	s.CiphertextBlob = v
	return s
}

func (s *DecryptRequest) SetKeyId(v string) *DecryptRequest {
	s.KeyId = &v
	return s
}

func (s *DecryptRequest) SetAlgorithm(v string) *DecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *DecryptRequest) SetAad(v []byte) *DecryptRequest {
	s.Aad = v
	return s
}

func (s *DecryptRequest) SetIv(v []byte) *DecryptRequest {
	s.Iv = v
	return s
}

func (s *DecryptRequest) SetPaddingMode(v string) *DecryptRequest {
	s.PaddingMode = &v
	return s
}

type DecryptResponse struct {
	Headers     map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId       *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Plaintext   []byte             `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	Algorithm   *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	PaddingMode *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
	RequestId   *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s DecryptResponse) GoString() string {
	return s.String()
}

func (s *DecryptResponse) SetHeaders(v map[string]*string) *DecryptResponse {
	s.Headers = v
	return s
}

func (s *DecryptResponse) SetKeyId(v string) *DecryptResponse {
	s.KeyId = &v
	return s
}

func (s *DecryptResponse) SetPlaintext(v []byte) *DecryptResponse {
	s.Plaintext = v
	return s
}

func (s *DecryptResponse) SetAlgorithm(v string) *DecryptResponse {
	s.Algorithm = &v
	return s
}

func (s *DecryptResponse) SetPaddingMode(v string) *DecryptResponse {
	s.PaddingMode = &v
	return s
}

func (s *DecryptResponse) SetRequestId(v string) *DecryptResponse {
	s.RequestId = &v
	return s
}

type HmacRequest struct {
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId   *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Message []byte             `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s HmacRequest) String() string {
	return tea.Prettify(s)
}

func (s HmacRequest) GoString() string {
	return s.String()
}

func (s *HmacRequest) SetHeaders(v map[string]*string) *HmacRequest {
	s.Headers = v
	return s
}

func (s *HmacRequest) SetKeyId(v string) *HmacRequest {
	s.KeyId = &v
	return s
}

func (s *HmacRequest) SetMessage(v []byte) *HmacRequest {
	s.Message = v
	return s
}

type HmacResponse struct {
	Headers   map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId     *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Signature []byte             `json:"Signature,omitempty" xml:"Signature,omitempty"`
	RequestId *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HmacResponse) String() string {
	return tea.Prettify(s)
}

func (s HmacResponse) GoString() string {
	return s.String()
}

func (s *HmacResponse) SetHeaders(v map[string]*string) *HmacResponse {
	s.Headers = v
	return s
}

func (s *HmacResponse) SetKeyId(v string) *HmacResponse {
	s.KeyId = &v
	return s
}

func (s *HmacResponse) SetSignature(v []byte) *HmacResponse {
	s.Signature = v
	return s
}

func (s *HmacResponse) SetRequestId(v string) *HmacResponse {
	s.RequestId = &v
	return s
}

type SignRequest struct {
	Headers     map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId       *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Algorithm   *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Digest      []byte             `json:"Digest,omitempty" xml:"Digest,omitempty"`
	Message     []byte             `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageType *string            `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s SignRequest) String() string {
	return tea.Prettify(s)
}

func (s SignRequest) GoString() string {
	return s.String()
}

func (s *SignRequest) SetHeaders(v map[string]*string) *SignRequest {
	s.Headers = v
	return s
}

func (s *SignRequest) SetKeyId(v string) *SignRequest {
	s.KeyId = &v
	return s
}

func (s *SignRequest) SetAlgorithm(v string) *SignRequest {
	s.Algorithm = &v
	return s
}

func (s *SignRequest) SetDigest(v []byte) *SignRequest {
	s.Digest = v
	return s
}

func (s *SignRequest) SetMessage(v []byte) *SignRequest {
	s.Message = v
	return s
}

func (s *SignRequest) SetMessageType(v string) *SignRequest {
	s.MessageType = &v
	return s
}

type SignResponse struct {
	Headers     map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId       *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Signature   []byte             `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Algorithm   *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	MessageType *string            `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	RequestId   *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SignResponse) String() string {
	return tea.Prettify(s)
}

func (s SignResponse) GoString() string {
	return s.String()
}

func (s *SignResponse) SetHeaders(v map[string]*string) *SignResponse {
	s.Headers = v
	return s
}

func (s *SignResponse) SetKeyId(v string) *SignResponse {
	s.KeyId = &v
	return s
}

func (s *SignResponse) SetSignature(v []byte) *SignResponse {
	s.Signature = v
	return s
}

func (s *SignResponse) SetAlgorithm(v string) *SignResponse {
	s.Algorithm = &v
	return s
}

func (s *SignResponse) SetMessageType(v string) *SignResponse {
	s.MessageType = &v
	return s
}

func (s *SignResponse) SetRequestId(v string) *SignResponse {
	s.RequestId = &v
	return s
}

type VerifyRequest struct {
	Headers     map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId       *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Signature   []byte             `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Algorithm   *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Digest      []byte             `json:"Digest,omitempty" xml:"Digest,omitempty"`
	Message     []byte             `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageType *string            `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s VerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyRequest) GoString() string {
	return s.String()
}

func (s *VerifyRequest) SetHeaders(v map[string]*string) *VerifyRequest {
	s.Headers = v
	return s
}

func (s *VerifyRequest) SetKeyId(v string) *VerifyRequest {
	s.KeyId = &v
	return s
}

func (s *VerifyRequest) SetSignature(v []byte) *VerifyRequest {
	s.Signature = v
	return s
}

func (s *VerifyRequest) SetAlgorithm(v string) *VerifyRequest {
	s.Algorithm = &v
	return s
}

func (s *VerifyRequest) SetDigest(v []byte) *VerifyRequest {
	s.Digest = v
	return s
}

func (s *VerifyRequest) SetMessage(v []byte) *VerifyRequest {
	s.Message = v
	return s
}

func (s *VerifyRequest) SetMessageType(v string) *VerifyRequest {
	s.MessageType = &v
	return s
}

type VerifyResponse struct {
	Headers     map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId       *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Value       *bool              `json:"Value,omitempty" xml:"Value,omitempty"`
	Algorithm   *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	MessageType *string            `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	RequestId   *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyResponse) GoString() string {
	return s.String()
}

func (s *VerifyResponse) SetHeaders(v map[string]*string) *VerifyResponse {
	s.Headers = v
	return s
}

func (s *VerifyResponse) SetKeyId(v string) *VerifyResponse {
	s.KeyId = &v
	return s
}

func (s *VerifyResponse) SetValue(v bool) *VerifyResponse {
	s.Value = &v
	return s
}

func (s *VerifyResponse) SetAlgorithm(v string) *VerifyResponse {
	s.Algorithm = &v
	return s
}

func (s *VerifyResponse) SetMessageType(v string) *VerifyResponse {
	s.MessageType = &v
	return s
}

func (s *VerifyResponse) SetRequestId(v string) *VerifyResponse {
	s.RequestId = &v
	return s
}

type GenerateRandomRequest struct {
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	Length  *int32             `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s GenerateRandomRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateRandomRequest) GoString() string {
	return s.String()
}

func (s *GenerateRandomRequest) SetHeaders(v map[string]*string) *GenerateRandomRequest {
	s.Headers = v
	return s
}

func (s *GenerateRandomRequest) SetLength(v int32) *GenerateRandomRequest {
	s.Length = &v
	return s
}

type GenerateRandomResponse struct {
	Headers   map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	Random    []byte             `json:"Random,omitempty" xml:"Random,omitempty"`
	RequestId *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateRandomResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateRandomResponse) GoString() string {
	return s.String()
}

func (s *GenerateRandomResponse) SetHeaders(v map[string]*string) *GenerateRandomResponse {
	s.Headers = v
	return s
}

func (s *GenerateRandomResponse) SetRandom(v []byte) *GenerateRandomResponse {
	s.Random = v
	return s
}

func (s *GenerateRandomResponse) SetRequestId(v string) *GenerateRandomResponse {
	s.RequestId = &v
	return s
}

type GenerateDataKeyRequest struct {
	Headers       map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId         *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Algorithm     *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	NumberOfBytes *int32             `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	Aad           []byte             `json:"Aad,omitempty" xml:"Aad,omitempty"`
}

func (s GenerateDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyRequest) SetHeaders(v map[string]*string) *GenerateDataKeyRequest {
	s.Headers = v
	return s
}

func (s *GenerateDataKeyRequest) SetKeyId(v string) *GenerateDataKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyRequest) SetAlgorithm(v string) *GenerateDataKeyRequest {
	s.Algorithm = &v
	return s
}

func (s *GenerateDataKeyRequest) SetNumberOfBytes(v int32) *GenerateDataKeyRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateDataKeyRequest) SetAad(v []byte) *GenerateDataKeyRequest {
	s.Aad = v
	return s
}

type GenerateDataKeyResponse struct {
	Headers        map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId          *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Iv             []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	Plaintext      []byte             `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	CiphertextBlob []byte             `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	Algorithm      *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	RequestId      *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDataKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyResponse) SetHeaders(v map[string]*string) *GenerateDataKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateDataKeyResponse) SetKeyId(v string) *GenerateDataKeyResponse {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyResponse) SetIv(v []byte) *GenerateDataKeyResponse {
	s.Iv = v
	return s
}

func (s *GenerateDataKeyResponse) SetPlaintext(v []byte) *GenerateDataKeyResponse {
	s.Plaintext = v
	return s
}

func (s *GenerateDataKeyResponse) SetCiphertextBlob(v []byte) *GenerateDataKeyResponse {
	s.CiphertextBlob = v
	return s
}

func (s *GenerateDataKeyResponse) SetAlgorithm(v string) *GenerateDataKeyResponse {
	s.Algorithm = &v
	return s
}

func (s *GenerateDataKeyResponse) SetRequestId(v string) *GenerateDataKeyResponse {
	s.RequestId = &v
	return s
}

type GetPublicKeyRequest struct {
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId   *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s GetPublicKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetPublicKeyRequest) SetHeaders(v map[string]*string) *GetPublicKeyRequest {
	s.Headers = v
	return s
}

func (s *GetPublicKeyRequest) SetKeyId(v string) *GetPublicKeyRequest {
	s.KeyId = &v
	return s
}

type GetPublicKeyResponse struct {
	Headers   map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	KeyId     *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	PublicKey *string            `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	RequestId *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublicKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponse) SetHeaders(v map[string]*string) *GetPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *GetPublicKeyResponse) SetKeyId(v string) *GetPublicKeyResponse {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyResponse) SetPublicKey(v string) *GetPublicKeyResponse {
	s.PublicKey = &v
	return s
}

func (s *GetPublicKeyResponse) SetRequestId(v string) *GetPublicKeyResponse {
	s.RequestId = &v
	return s
}

type GetSecretValueRequest struct {
	Headers             map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	SecretName          *string            `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	VersionStage        *string            `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
	VersionId           *string            `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	FetchExtendedConfig *bool              `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
}

func (s GetSecretValueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueRequest) GoString() string {
	return s.String()
}

func (s *GetSecretValueRequest) SetHeaders(v map[string]*string) *GetSecretValueRequest {
	s.Headers = v
	return s
}

func (s *GetSecretValueRequest) SetSecretName(v string) *GetSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionStage(v string) *GetSecretValueRequest {
	s.VersionStage = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionId(v string) *GetSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueRequest) SetFetchExtendedConfig(v bool) *GetSecretValueRequest {
	s.FetchExtendedConfig = &v
	return s
}

type GetSecretValueResponse struct {
	Headers           map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	SecretName        *string            `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	SecretType        *string            `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	SecretData        *string            `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	SecretDataType    *string            `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	VersionStages     []*string          `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Repeated"`
	VersionId         *string            `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	CreateTime        *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RequestId         *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LastRotationDate  *string            `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	NextRotationDate  *string            `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	ExtendedConfig    *string            `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	AutomaticRotation *string            `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	RotationInterval  *string            `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s GetSecretValueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueResponse) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponse) SetHeaders(v map[string]*string) *GetSecretValueResponse {
	s.Headers = v
	return s
}

func (s *GetSecretValueResponse) SetSecretName(v string) *GetSecretValueResponse {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueResponse) SetSecretType(v string) *GetSecretValueResponse {
	s.SecretType = &v
	return s
}

func (s *GetSecretValueResponse) SetSecretData(v string) *GetSecretValueResponse {
	s.SecretData = &v
	return s
}

func (s *GetSecretValueResponse) SetSecretDataType(v string) *GetSecretValueResponse {
	s.SecretDataType = &v
	return s
}

func (s *GetSecretValueResponse) SetVersionStages(v []*string) *GetSecretValueResponse {
	s.VersionStages = v
	return s
}

func (s *GetSecretValueResponse) SetVersionId(v string) *GetSecretValueResponse {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueResponse) SetCreateTime(v string) *GetSecretValueResponse {
	s.CreateTime = &v
	return s
}

func (s *GetSecretValueResponse) SetRequestId(v string) *GetSecretValueResponse {
	s.RequestId = &v
	return s
}

func (s *GetSecretValueResponse) SetLastRotationDate(v string) *GetSecretValueResponse {
	s.LastRotationDate = &v
	return s
}

func (s *GetSecretValueResponse) SetNextRotationDate(v string) *GetSecretValueResponse {
	s.NextRotationDate = &v
	return s
}

func (s *GetSecretValueResponse) SetExtendedConfig(v string) *GetSecretValueResponse {
	s.ExtendedConfig = &v
	return s
}

func (s *GetSecretValueResponse) SetAutomaticRotation(v string) *GetSecretValueResponse {
	s.AutomaticRotation = &v
	return s
}

func (s *GetSecretValueResponse) SetRotationInterval(v string) *GetSecretValueResponse {
	s.RotationInterval = &v
	return s
}

type AdvanceEncryptRequest struct {
	KeyId          *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Plaintext      []byte             `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	Algorithm      *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Aad            []byte             `json:"Aad,omitempty" xml:"Aad,omitempty"`
	Iv             []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	PaddingMode    *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
	Headers map[string]*string 		  `json:"Headers,omitempty" xml:"Headers,omitempty"`
}

func (s AdvanceEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AdvanceEncryptRequest) GoString() string {
	return s.String()
}

func (s *AdvanceEncryptRequest) SetKeyId(v string) *AdvanceEncryptRequest {
	s.KeyId = &v
	return s
}

func (s *AdvanceEncryptRequest) SetPlaintext(v []byte) *AdvanceEncryptRequest {
	s.Plaintext = v
	return s
}

func (s *AdvanceEncryptRequest) SetAlgorithm(v string) *AdvanceEncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *AdvanceEncryptRequest) SetAad(v []byte) *AdvanceEncryptRequest {
	s.Aad = v
	return s
}

func (s *AdvanceEncryptRequest) SetIv(v []byte) *AdvanceEncryptRequest {
	s.Iv = v
	return s
}

func (s *AdvanceEncryptRequest) SetPaddingMode(v string) *AdvanceEncryptRequest {
	s.PaddingMode = &v
	return s
}

func (s *AdvanceEncryptRequest) SetHeaders(v map[string]*string) *AdvanceEncryptRequest {
	s.Headers = v
	return s
}

type AdvanceEncryptResponse struct {
	KeyId           *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	CiphertextBlob  []byte             `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	Iv              []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	RequestId       *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Algorithm       *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	PaddingMode     *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
	KeyVersionId    *string            `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	Headers map[string]*string 		   `json:"Headers,omitempty" xml:"Headers,omitempty"`
}

func (s AdvanceEncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s AdvanceEncryptResponse) GoString() string {
	return s.String()
}

func (s *AdvanceEncryptResponse) SetKeyId(v string) *AdvanceEncryptResponse {
	s.KeyId = &v
	return s
}

func (s *AdvanceEncryptResponse) SetCiphertextBlob(v []byte) *AdvanceEncryptResponse {
	s.CiphertextBlob = v
	return s
}

func (s *AdvanceEncryptResponse) SetIv(v []byte) *AdvanceEncryptResponse {
	s.Iv = v
	return s
}

func (s *AdvanceEncryptResponse) SetRequestId(v string) *AdvanceEncryptResponse {
	s.RequestId = &v
	return s
}

func (s *AdvanceEncryptResponse) SetAlgorithm(v string) *AdvanceEncryptResponse {
	s.Algorithm = &v
	return s
}

func (s *AdvanceEncryptResponse) SetPaddingMode(v string) *AdvanceEncryptResponse {
	s.PaddingMode = &v
	return s
}

func (s *AdvanceEncryptResponse) SetKeyVersionId(v string) *AdvanceEncryptResponse {
	s.KeyVersionId = &v
	return s
}

func (s *AdvanceEncryptResponse) SetHeaders(v map[string]*string) *AdvanceEncryptResponse {
	s.Headers = v
	return s
}

type AdvanceDecryptRequest struct {
	CiphertextBlob []byte             `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	KeyId          *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Algorithm      *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Aad            []byte             `json:"Aad,omitempty" xml:"Aad,omitempty"`
	Iv             []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	PaddingMode    *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
	Headers map[string]*string 		  `json:"Headers,omitempty" xml:"Headers,omitempty"`
}

func (s AdvanceDecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AdvanceDecryptRequest) GoString() string {
	return s.String()
}

func (s *AdvanceDecryptRequest) SetCiphertextBlob(v []byte) *AdvanceDecryptRequest {
	s.CiphertextBlob = v
	return s
}

func (s *AdvanceDecryptRequest) SetKeyId(v string) *AdvanceDecryptRequest {
	s.KeyId = &v
	return s
}

func (s *AdvanceDecryptRequest) SetAlgorithm(v string) *AdvanceDecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *AdvanceDecryptRequest) SetAad(v []byte) *AdvanceDecryptRequest {
	s.Aad = v
	return s
}

func (s *AdvanceDecryptRequest) SetIv(v []byte) *AdvanceDecryptRequest {
	s.Iv = v
	return s
}

func (s *AdvanceDecryptRequest) SetPaddingMode(v string) *AdvanceDecryptRequest {
	s.PaddingMode = &v
	return s
}

func (s *AdvanceDecryptRequest) SetHeaders(v map[string]*string) *AdvanceDecryptRequest {
	s.Headers = v
	return s
}

type AdvanceDecryptResponse struct {
	KeyId           *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Plaintext       []byte             `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	RequestId       *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Algorithm       *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	PaddingMode     *string            `json:"PaddingMode,omitempty" xml:"PaddingMode,omitempty"`
	KeyVersionId    *string            `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	Headers map[string]*string 		   `json:"Headers,omitempty" xml:"Headers,omitempty"`
}

func (s AdvanceDecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s AdvanceDecryptResponse) GoString() string {
	return s.String()
}

func (s *AdvanceDecryptResponse) SetKeyId(v string) *AdvanceDecryptResponse {
	s.KeyId = &v
	return s
}

func (s *AdvanceDecryptResponse) SetPlaintext(v []byte) *AdvanceDecryptResponse {
	s.Plaintext = v
	return s
}

func (s *AdvanceDecryptResponse) SetRequestId(v string) *AdvanceDecryptResponse {
	s.RequestId = &v
	return s
}

func (s *AdvanceDecryptResponse) SetAlgorithm(v string) *AdvanceDecryptResponse {
	s.Algorithm = &v
	return s
}

func (s *AdvanceDecryptResponse) SetPaddingMode(v string) *AdvanceDecryptResponse {
	s.PaddingMode = &v
	return s
}

func (s *AdvanceDecryptResponse) SetKeyVersionId(v string) *AdvanceDecryptResponse {
	s.KeyVersionId = &v
	return s
}

func (s *AdvanceDecryptResponse) SetHeaders(v map[string]*string) *AdvanceDecryptResponse {
	s.Headers = v
	return s
}

type AdvanceGenerateDataKeyRequest struct {
	KeyId          *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	NumberOfBytes  *int32             `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	Aad            []byte             `json:"Aad,omitempty" xml:"Aad,omitempty"`
	Headers map[string]*string 		  `json:"Headers,omitempty" xml:"Headers,omitempty"`
}

func (s AdvanceGenerateDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s AdvanceGenerateDataKeyRequest) GoString() string {
	return s.String()
}

func (s *AdvanceGenerateDataKeyRequest) SetKeyId(v string) *AdvanceGenerateDataKeyRequest {
	s.KeyId = &v
	return s
}

func (s *AdvanceGenerateDataKeyRequest) SetNumberOfBytes(v int32) *AdvanceGenerateDataKeyRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *AdvanceGenerateDataKeyRequest) SetAad(v []byte) *AdvanceGenerateDataKeyRequest {
	s.Aad = v
	return s
}

func (s *AdvanceGenerateDataKeyRequest) SetHeaders(v map[string]*string) *AdvanceGenerateDataKeyRequest {
	s.Headers = v
	return s
}

type AdvanceGenerateDataKeyResponse struct {
	KeyId           *string            `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Iv              []byte             `json:"Iv,omitempty" xml:"Iv,omitempty"`
	Plaintext       []byte             `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	CiphertextBlob  []byte             `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	RequestId       *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Algorithm       *string            `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	KeyVersionId    *string            `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	Headers map[string]*string 		   `json:"Headers,omitempty" xml:"Headers,omitempty"`
}

func (s AdvanceGenerateDataKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s AdvanceGenerateDataKeyResponse) GoString() string {
	return s.String()
}

func (s *AdvanceGenerateDataKeyResponse) SetKeyId(v string) *AdvanceGenerateDataKeyResponse {
	s.KeyId = &v
	return s
}

func (s *AdvanceGenerateDataKeyResponse) SetIv(v []byte) *AdvanceGenerateDataKeyResponse {
	s.Iv = v
	return s
}

func (s *AdvanceGenerateDataKeyResponse) SetPlaintext(v []byte) *AdvanceGenerateDataKeyResponse {
	s.Plaintext = v
	return s
}

func (s *AdvanceGenerateDataKeyResponse) SetCiphertextBlob(v []byte) *AdvanceGenerateDataKeyResponse {
	s.CiphertextBlob = v
	return s
}

func (s *AdvanceGenerateDataKeyResponse) SetRequestId(v string) *AdvanceGenerateDataKeyResponse {
	s.RequestId = &v
	return s
}

func (s *AdvanceGenerateDataKeyResponse) SetAlgorithm(v string) *AdvanceGenerateDataKeyResponse {
	s.Algorithm = &v
	return s
}

func (s *AdvanceGenerateDataKeyResponse) SetKeyVersionId(v string) *AdvanceGenerateDataKeyResponse {
	s.KeyVersionId = &v
	return s
}

func (s *AdvanceGenerateDataKeyResponse) SetHeaders(v map[string]*string) *AdvanceGenerateDataKeyResponse {
	s.Headers = v
	return s
}

type Client struct {
	dedicatedkmsopenapi.Client
}

func NewClient(config *dedicatedkmsopenapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *dedicatedkmsopenapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	return nil
}

func (client *Client) EncryptWithOptions(request *EncryptRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *EncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedEncryptRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("Encrypt"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseEncryptResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &EncryptResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":        response["headers"],
		"RequestId":      respMap["RequestId"],
		"KeyId":          respMap["KeyId"],
		"CiphertextBlob": respMap["CiphertextBlob"],
		"Iv":             respMap["Iv"],
		"Algorithm":      respMap["Algorithm"],
		"PaddingMode":    respMap["PaddingMode"],
	}, &_result)
	return _result, _err
}

func (client *Client) Encrypt(request *EncryptRequest) (_result *EncryptResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &EncryptResponse{}
	_body, _err := client.EncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DecryptWithOptions(request *DecryptRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *DecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedDecryptRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("Decrypt"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseDecryptResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &DecryptResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":     response["headers"],
		"RequestId":   respMap["RequestId"],
		"KeyId":       respMap["KeyId"],
		"Plaintext":   respMap["Plaintext"],
		"Algorithm":   respMap["Algorithm"],
		"PaddingMode": respMap["PaddingMode"],
	}, &_result)
	return _result, _err
}

func (client *Client) Decrypt(request *DecryptRequest) (_result *DecryptResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &DecryptResponse{}
	_body, _err := client.DecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HmacWithOptions(request *HmacRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *HmacResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedHmacRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("Hmac"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseHmacResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &HmacResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":   response["headers"],
		"RequestId": respMap["RequestId"],
		"KeyId":     respMap["KeyId"],
		"Signature": respMap["Signature"],
	}, &_result)
	return _result, _err
}

func (client *Client) Hmac(request *HmacRequest) (_result *HmacResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &HmacResponse{}
	_body, _err := client.HmacWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignWithOptions(request *SignRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *SignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedSignRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("Sign"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseSignResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &SignResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":     response["headers"],
		"RequestId":   respMap["RequestId"],
		"KeyId":       respMap["KeyId"],
		"Signature":   respMap["Signature"],
		"Algorithm":   respMap["Algorithm"],
		"MessageType": respMap["MessageType"],
	}, &_result)
	return _result, _err
}

func (client *Client) Sign(request *SignRequest) (_result *SignResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &SignResponse{}
	_body, _err := client.SignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyWithOptions(request *VerifyRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *VerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedVerifyRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("Verify"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseVerifyResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":     response["headers"],
		"RequestId":   respMap["RequestId"],
		"KeyId":       respMap["KeyId"],
		"Value":       respMap["Value"],
		"Algorithm":   respMap["Algorithm"],
		"MessageType": respMap["MessageType"],
	}, &_result)
	return _result, _err
}

func (client *Client) Verify(request *VerifyRequest) (_result *VerifyResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &VerifyResponse{}
	_body, _err := client.VerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateRandomWithOptions(request *GenerateRandomRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *GenerateRandomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedGenerateRandomRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("GenerateRandom"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseGenerateRandomResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &GenerateRandomResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":   response["headers"],
		"RequestId": respMap["RequestId"],
		"Random":    respMap["Random"],
	}, &_result)
	return _result, _err
}

func (client *Client) GenerateRandom(request *GenerateRandomRequest) (_result *GenerateRandomResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &GenerateRandomResponse{}
	_body, _err := client.GenerateRandomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDataKeyWithOptions(request *GenerateDataKeyRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *GenerateDataKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedGenerateDataKeyRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("GenerateDataKey"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseGenerateDataKeyResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &GenerateDataKeyResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":        response["headers"],
		"RequestId":      respMap["RequestId"],
		"KeyId":          respMap["KeyId"],
		"Iv":             respMap["Iv"],
		"Plaintext":      respMap["Plaintext"],
		"CiphertextBlob": respMap["CiphertextBlob"],
		"Algorithm":      respMap["Algorithm"],
	}, &_result)
	return _result, _err
}

func (client *Client) GenerateDataKey(request *GenerateDataKeyRequest) (_result *GenerateDataKeyResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &GenerateDataKeyResponse{}
	_body, _err := client.GenerateDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublicKeyWithOptions(request *GetPublicKeyRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *GetPublicKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedGetPublicKeyRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("GetPublicKey"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseGetPublicKeyResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPublicKeyResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":   response["headers"],
		"RequestId": respMap["RequestId"],
		"KeyId":     respMap["KeyId"],
		"PublicKey": respMap["PublicKey"],
	}, &_result)
	return _result, _err
}

func (client *Client) GetPublicKey(request *GetPublicKeyRequest) (_result *GetPublicKeyResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &GetPublicKeyResponse{}
	_body, _err := client.GetPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecretValueWithOptions(request *GetSecretValueRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, _err := dedicatedkmsopenapiutil.GetSerializedGetSecretValueRequest(reqBody)
	if _err != nil {
		return _result, _err
	}
	response, _err := client.DoRequest(tea.String("GetSecretValue"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if _err != nil {
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, _err := dedicatedkmsopenapiutil.ParseGetSecretValueResponse(respBody)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSecretValueResponse{}
	_err = tea.Convert(map[string]interface{}{
		"Headers":           response["headers"],
		"RequestId":         respMap["RequestId"],
		"SecretName":        respMap["SecretName"],
		"SecretType":        respMap["SecretType"],
		"SecretData":        respMap["SecretData"],
		"SecretDataType":    respMap["SecretDataType"],
		"VersionStages":     respMap["VersionStages"],
		"VersionId":         respMap["VersionId"],
		"CreateTime":        respMap["CreateTime"],
		"LastRotationDate":  respMap["LastRotationDate"],
		"NextRotationDate":  respMap["NextRotationDate"],
		"ExtendedConfig":    respMap["ExtendedConfig"],
		"AutomaticRotation": respMap["AutomaticRotation"],
		"RotationInterval":  respMap["RotationInterval"],
	}, &_result)
	return _result, _err
}

func (client *Client) GetSecretValue(request *GetSecretValueRequest) (_result *GetSecretValueResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &GetSecretValueResponse{}
	_body, _err := client.GetSecretValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AdvanceEncrypt(request *AdvanceEncryptRequest) (_result *AdvanceEncryptResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &AdvanceEncryptResponse{}
	_body, _err := client.AdvanceEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AdvanceEncryptWithOptions(request *AdvanceEncryptRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *AdvanceEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, err := dedicatedkmsopenapiutil.GetSerializedAdvanceEncryptRequest(reqBody)
	response, err := client.DoRequest(tea.String("AdvanceEncrypt"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, err := dedicatedkmsopenapiutil.ParseAdvanceEncryptResponse(respBody)
	if err != nil {
		_err = err
		return _result, _err
	}
	_result = &AdvanceEncryptResponse{}
	_err = tea.Convert(map[string]interface{}{
		"KeyId":           respMap["KeyId"],
		"CiphertextBlob":  respMap["CiphertextBlob"],
		"Iv":              respMap["Iv"],
		"RequestId":       respMap["RequestId"],
		"Algorithm":       respMap["Algorithm"],
		"PaddingMode":     respMap["PaddingMode"],
		"KeyVersionId":    respMap["KeyVersionId"],
		"Headers": 		   response["headers"],
	}, &_result)
	return _result, _err
}

func (client *Client) AdvanceDecrypt(request *AdvanceDecryptRequest) (_result *AdvanceDecryptResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &AdvanceDecryptResponse{}
	_body, _err := client.AdvanceDecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AdvanceDecryptWithOptions(request *AdvanceDecryptRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *AdvanceDecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, err := dedicatedkmsopenapiutil.GetSerializedAdvanceDecryptRequest(reqBody)
	response, err := client.DoRequest(tea.String("AdvanceDecrypt"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, err := dedicatedkmsopenapiutil.ParseAdvanceDecryptResponse(respBody)
	if err != nil {
		_err = err
		return _result, _err
	}
	_result = &AdvanceDecryptResponse{}
	_err = tea.Convert(map[string]interface{}{
		"KeyId":           respMap["KeyId"],
		"Plaintext":       respMap["Plaintext"],
		"RequestId":       respMap["RequestId"],
		"Algorithm":       respMap["Algorithm"],
		"PaddingMode":     respMap["PaddingMode"],
		"KeyVersionId":    respMap["KeyVersionId"],
		"Headers": response["headers"],
	}, &_result)
	return _result, _err
}

func (client *Client) AdvanceGenerateDataKey(request *AdvanceGenerateDataKeyRequest) (_result *AdvanceGenerateDataKeyResponse, _err error) {
	runtime := &dedicatedkmsopenapiutil.RuntimeOptions{}
	_result = &AdvanceGenerateDataKeyResponse{}
	_body, _err := client.AdvanceGenerateDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AdvanceGenerateDataKeyWithOptions(request *AdvanceGenerateDataKeyRequest, runtime *dedicatedkmsopenapiutil.RuntimeOptions) (_result *AdvanceGenerateDataKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	
	reqBody := dedicatedkmsopenapiutil.ConvertToMap(request)
	reqBodyBytes, err := dedicatedkmsopenapiutil.GetSerializedAdvanceGenerateDataKeyRequest(reqBody)
	response, err := client.DoRequest(tea.String("AdvanceGenerateDataKey"), tea.String("dkms-gcs-0.2"), tea.String("https"), tea.String("POST"), tea.String("RSA_PKCS1_SHA_256"), reqBodyBytes, request.Headers, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}

	respBody := util.AssertAsBytes(response["body"])
	respMap, err := dedicatedkmsopenapiutil.ParseAdvanceGenerateDataKeyResponse(respBody)
	if err != nil {
		_err = err
		return _result, _err
	}
	_result = &AdvanceGenerateDataKeyResponse{}
	_err = tea.Convert(map[string]interface{}{
		"KeyId":           respMap["KeyId"],
		"Iv":              respMap["Iv"],
		"Plaintext":       respMap["Plaintext"],
		"CiphertextBlob":  respMap["CiphertextBlob"],
		"RequestId":       respMap["RequestId"],
		"Algorithm":       respMap["Algorithm"],
		"KeyVersionId":    respMap["KeyVersionId"],
		"Headers": response["headers"],
	}, &_result)
	return _result, _err
}

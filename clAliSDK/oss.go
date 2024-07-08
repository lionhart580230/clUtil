package clAliSDK

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
)

type OssConfig struct {
	OssId         uint32
	OssEndPoint   string
	OssAccessKey  string
	OssSecretKey  string
	OssBucketName string
	OssDomain     string // oss请求地址

	client *oss.Client
}

// 创建一个新的oss配置对象
func New(_bucketName, _accessKey, _secretKey, _endPoint, _domain string) *OssConfig {
	return &OssConfig{
		OssEndPoint:   _endPoint,
		OssAccessKey:  _accessKey,
		OssSecretKey:  _secretKey,
		OssBucketName: _bucketName,
		OssDomain:     _domain,
		client:        nil,
	}
}

// 获取客户端
func (this *OssConfig) GetClient() (error, *oss.Client) {
	if this.client == nil {
		var err error
		this.client, err = oss.New(this.OssEndPoint, this.OssAccessKey, this.OssSecretKey)
		if err != nil {
			return err, nil
		}
	}
	return nil, this.client
}

// 上传文件到oss
func (this *OssConfig) UploadDataTo(_sourceData []byte, _toPath string, _tryCount uint32) error {

	err, client := this.GetClient()
	if err != nil {
		return err
	}

	ossBucket, err := client.Bucket(this.OssBucketName)
	if err != nil {
		return err
	}

	// 标准存储
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	// 公共读
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	err = ossBucket.PutObject(_toPath, bytes.NewReader(_sourceData), storageType, objectAcl)
	if err != nil {
		if strings.Contains(err.Error(), "UserDisable") {
			return err
		}
		if _tryCount > 0 {
			return this.UploadDataTo(_sourceData, _toPath, _tryCount-1)
		}
	}
	return err
}

// 上传文件到oss
func (this *OssConfig) UploadFileTo(_path string, _toPath string, _tryCount uint32) error {

	err, client := this.GetClient()
	if err != nil {
		return err
	}
	ossBucket, err := client.Bucket(this.OssBucketName)
	if err != nil {
		return err
	}

	// 标准存储
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	// 公共读
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	err = ossBucket.PutObjectFromFile(_toPath, _path, storageType, objectAcl)
	if err != nil {
		if strings.Contains(err.Error(), "UserDisable") {
			return err
		}
		if _tryCount > 0 {
			return this.UploadFileTo(_path, _toPath, _tryCount-1)
		}
	}
	return err
}

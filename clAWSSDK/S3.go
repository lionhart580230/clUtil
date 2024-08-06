package clAWSSDK

import (
	"bytes"
	"github.com/minio/minio-go"
	"io/ioutil"
)

type OssConf struct {
	EndPoint  string
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
}

var mOssConf OssConf
var mClient *minio.Client

// 配置OSS
func InitOssConf(_conf OssConf) {
	mOssConf = _conf
}

// 获取Minion客户端
func GetClient() (error, *minio.Client) {
	if mClient != nil {
		return nil, mClient
	}
	minioClient, err := minio.NewWithRegion(mOssConf.EndPoint, mOssConf.AccessKey, mOssConf.SecretKey, true, mOssConf.Region)
	if err != nil {
		return err, nil
	}
	return nil, minioClient
}

// 上传文件
func UploadFile(_fromPath, _toPath string, _tryCount int32) error {
	fileBuffer, err := ioutil.ReadFile(_fromPath)
	if err != nil {
		return err
	}

	err, client := GetClient()
	if err != nil {
		return err
	}

	_, err = client.PutObject(mOssConf.Bucket, _toPath, bytes.NewBuffer(fileBuffer), int64(len(fileBuffer)), minio.PutObjectOptions{
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	})
	if err != nil {
		if _tryCount > 0 {
			return UploadFile(_fromPath, _toPath, _tryCount-1)
		}
		return err
	}

	return nil
}

// 上传文件内容
func UploadFileBySource(_source []byte, _toPath string, _tryCount int32) error {
	err, client := GetClient()
	if err != nil {
		return err
	}

	_, err = client.PutObject(mOssConf.Bucket, _toPath, bytes.NewBuffer(_source), int64(len(_source)), minio.PutObjectOptions{
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	})
	if err != nil {
		if _tryCount > 0 {
			return UploadFileBySource(_source, _toPath, _tryCount-1)
		}
		return err
	}

	return nil
}

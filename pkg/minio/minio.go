package minio

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type IMinio interface {
	BucketPublicName() string

	BucketPrivateName() string

	Stat(bucketName string, objectName string) (*FileStatInfo, error)

	Write(bucketName string, objectName string, stream []byte) error

	Copy(bucketName string, srcObjectName, objectName string) error

	CopyObject(srcBucketName string, srcObjectName, dstBucketName string, dstObjectName string) error

	Delete(bucketName string, objectName string) error

	GetObject(bucketName string, objectName string) ([]byte, error)

	PublicUrl(bucketName, objectName string) string

	PrivateUrl(bucketName, objectName string, filename string, expire time.Duration) string

	InitiateMultipartUpload(bucketName, objectName string) (string, error)

	PutObjectPart(bucketName, objectName string, uploadID string, index int, data io.Reader, size int64) (ObjectPart, error)

	CompleteMultipartUpload(bucketName, objectName, uploadID string, parts []ObjectPart) error

	AbortMultipartUpload(bucketName, objectName, uploadID string) error
}

var _ IMinio = (*Minio)(nil)

type Config struct {
	Endpoint      string
	SSL           bool
	SecretId      string
	SecretKey     string
	BucketPublic  string
	BucketPrivate string
}

type Minio struct {
	Core   *minio.Core
	Config Config
}

func NewMinio(config Config) *Minio {
	client, err := minio.NewCore(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.SecretId, config.SecretKey, ""),
		Secure: config.SSL,
	})

	if err != nil {
		panic(fmt.Sprintf("Не удалось инициализировать minio-клиент, %s", err))
	}

	return &Minio{
		Core:   client,
		Config: config,
	}
}

func (m Minio) BucketPublicName() string {
	return m.Config.BucketPublic
}

func (m Minio) BucketPrivateName() string {
	return m.Config.BucketPrivate
}

type FileStatInfo struct {
	Name        string
	Size        int64
	Ext         string
	MimeType    string
	LastModTime time.Time
}

func (m Minio) Stat(bucketName string, objectName string) (*FileStatInfo, error) {
	objInfo, err := m.Core.Client.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &FileStatInfo{
		LastModTime: objInfo.LastModified,
		MimeType:    objInfo.ContentType,
		Name:        objInfo.Key,
		Size:        objInfo.Size,
		Ext:         path.Ext(objectName),
	}, nil
}

func (m Minio) Write(bucketName string, objectName string, stream []byte) error {
	_, err := m.Core.Client.PutObject(context.Background(), bucketName, objectName, strings.NewReader(string(stream)), int64(len(stream)), minio.PutObjectOptions{})
	return err
}

func (m Minio) Copy(bucketName string, srcObjectName, objectName string) error {
	return m.CopyObject(bucketName, srcObjectName, bucketName, objectName)
}

func (m Minio) CopyObject(srcBucketName string, srcObjectName, dstBucketName string, dstObjectName string) error {
	srcOpts := minio.CopySrcOptions{
		Bucket: srcBucketName,
		Object: srcObjectName,
	}

	dstOpts := minio.CopyDestOptions{
		Bucket: dstBucketName,
		Object: dstObjectName,
	}

	_, err := m.Core.Client.CopyObject(context.Background(), dstOpts, srcOpts)
	return err
}

func (m Minio) Delete(bucketName string, objectName string) error {
	return m.Core.Client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
}

func (m Minio) GetObject(bucketName string, objectName string) ([]byte, error) {
	object, err := m.Core.Client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer object.Close()

	return io.ReadAll(object)
}

func (m Minio) PublicUrl(bucketName, objectName string) string {
	uri, err := m.Core.Client.PresignedGetObject(context.Background(), bucketName, objectName, 30*time.Minute, nil)
	if err != nil {
		panic(err)
	}

	if m.BucketPublicName() == bucketName {
		uri.RawQuery = ""
	}

	return uri.String()
}

func (m Minio) PrivateUrl(bucketName, objectName string, filename string, expire time.Duration) string {
	// set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	uri, err := m.Core.Client.PresignedGetObject(context.Background(), bucketName, objectName, expire, reqParams)
	if err != nil {
		panic(err)
	}

	return uri.String()
}

func (m Minio) InitiateMultipartUpload(bucketName, objectName string) (string, error) {
	return m.Core.NewMultipartUpload(context.Background(), bucketName, objectName, minio.PutObjectOptions{})
}

type ObjectPart struct {
	PartNumber     int
	ETag           string
	PartObjectName string
}

func (m Minio) PutObjectPart(bucketName, objectName string, uploadID string, index int, data io.Reader, size int64) (ObjectPart, error) {
	part, err := m.Core.PutObjectPart(context.Background(), bucketName, objectName, uploadID, index, data, size, minio.PutObjectPartOptions{})
	if err != nil {
		return ObjectPart{}, err
	}

	return ObjectPart{
		PartNumber: part.PartNumber,
		ETag:       part.ETag,
	}, nil
}

func (m Minio) CompleteMultipartUpload(bucketName, objectName, uploadID string, parts []ObjectPart) error {
	completeParts := make([]minio.CompletePart, 0)

	for _, part := range parts {
		completeParts = append(completeParts, minio.CompletePart{
			PartNumber: part.PartNumber,
			ETag:       part.ETag,
		})
	}

	_, err := m.Core.CompleteMultipartUpload(context.Background(), bucketName, objectName, uploadID, completeParts, minio.PutObjectOptions{})
	return err
}

func (m Minio) AbortMultipartUpload(bucketName, objectName, uploadID string) error {
	return m.Core.AbortMultipartUpload(context.Background(), bucketName, objectName, uploadID)
}

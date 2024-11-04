package provider

import (
	"voo.su/internal/config"
	"voo.su/pkg/minio"
)

func NewMinioClient(conf *config.Config) minio.IMinio {
	return minio.NewMinio(minio.Config{
		Endpoint:      conf.Minio.Endpoint,
		SSL:           conf.Minio.SSL,
		SecretId:      conf.Minio.SecretId,
		SecretKey:     conf.Minio.SecretKey,
		BucketPublic:  conf.Minio.BucketPublic,
		BucketPrivate: conf.Minio.BucketPrivate,
	})
}

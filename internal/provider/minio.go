package provider

import (
	"voo.su/internal/config"
	"voo.su/pkg/minio"
)

func NewMinioClient(conf *config.Config) minio.IMinio {
	return minio.NewMinio(minio.Config{
		Endpoint:  conf.Minio.Host,
		SSL:       conf.Minio.SSL,
		SecretId:  conf.Minio.SecretId,
		SecretKey: conf.Minio.SecretKey,
	})
}

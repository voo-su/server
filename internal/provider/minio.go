// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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

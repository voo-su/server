package usecase

import (
	"voo.su/pkg/locale"
	"voo.su/pkg/minio"
)

type StorageUseCase struct {
	Locale locale.ILocale
	Minio  minio.IMinio
}

func NewStorageUseCase(
	locale locale.ILocale,
	minio minio.IMinio,
) *StorageUseCase {
	return &StorageUseCase{
		Locale: locale,
		Minio:  minio,
	}
}

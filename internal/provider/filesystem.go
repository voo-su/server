package provider

import (
	"voo.su/internal/config"
	"voo.su/pkg/filesystem"
)

func NewFilesystem(conf *config.Config) *filesystem.Filesystem {
	s := &filesystem.Filesystem{}
	s.SetDriver(conf.File.Default)
	s.Local = filesystem.NewLocalFilesystem(conf)
	s.Default = s.Local

	return s
}

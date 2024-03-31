package filesystem

import (
	"time"
	"voo.su/internal/config"
)

type IAdapter interface {
	Write(data []byte, filePath string) error
	WriteLocal(localFile string, filePath string) error
	Copy(srcPath, filePath string) error
	Delete(filePath string) error
	DeleteDir(path string) error
	CreateDir(path string) error
	Stat(filePath string) (*FileStat, error)
	PublicUrl(filePath string) string
	PrivateUrl(filePath string, timeout time.Duration) string
	ReadStream(filePath string) ([]byte, error)
	InitiateMultipartUpload(filePath string, fileName string) (string, error)
}

type FileStat struct {
	Name        string
	Size        int64
	Ext         string
	LastModTime time.Time
	MimeType    string
}

type Filesystem struct {
	driver  string
	Default IAdapter
	Local   *LocalFilesystem
}

func NewFilesystem(conf *config.Config) *Filesystem {
	s := &Filesystem{}
	s.driver = conf.File.Default
	s.Local = NewLocalFilesystem(conf)
	s.Default = s.Local

	return s
}

func (f *Filesystem) Driver() string {
	return f.driver
}

func (f *Filesystem) SetDriver(value string) {
	f.driver = value
}

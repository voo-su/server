package filesystem

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	"voo.su/internal/config"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/strutil"
)

type LocalFilesystem struct {
	conf *config.Config
}

func NewLocalFilesystem(conf *config.Config) *LocalFilesystem {
	return &LocalFilesystem{
		conf: conf,
	}
}

func isDirExist(fileAddr string) bool {
	s, err := os.Stat(fileAddr)

	return err == nil && s.IsDir()
}

func (s *LocalFilesystem) Path(path string) string {
	return fmt.Sprintf("%s/%s",
		strings.TrimRight(s.conf.File.Local.Root, "/"),
		strings.TrimLeft(path, "/"),
	)
}

func (s *LocalFilesystem) Write(data []byte, filePath string) error {
	filePath = s.Path(filePath)
	dir := path.Dir(filePath)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	_, err = f.Write(data)

	return err
}

func (s *LocalFilesystem) WriteLocal(localFile string, filePath string) error {
	filePath = s.Path(filePath)
	srcFile, err := os.Open(localFile)
	if err != nil {
		return err
	}

	defer srcFile.Close()
	dir := path.Dir(filePath)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)

	return err
}

func (s *LocalFilesystem) AppendWrite(data []byte, filePath string) error {
	filePath = s.Path(filePath)
	dir := path.Dir(filePath)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0766); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0766)
	if err != nil {
		return err
	}

	_, err = f.Write(data)

	return err
}

func (s *LocalFilesystem) Copy(srcPath, filePath string) error {
	return s.WriteLocal(s.Path(srcPath), filePath)
}

func (s *LocalFilesystem) Delete(filePath string) error {
	return os.Remove(s.Path(filePath))
}

func (s *LocalFilesystem) CreateDir(dir string) error {
	return os.MkdirAll(s.Path(dir), 0766)
}

func (s *LocalFilesystem) DeleteDir(dir string) error {
	return os.RemoveAll(s.Path(dir))
}

func (s *LocalFilesystem) Stat(filePath string) (*FileStat, error) {
	info, err := os.Stat(s.Path(filePath))
	if err != nil {
		return nil, err
	}

	return &FileStat{
		Name:        filepath.Base(filePath),
		Size:        info.Size(),
		Ext:         filepath.Ext(filePath),
		MimeType:    "",
		LastModTime: info.ModTime(),
	}, nil
}

func (s *LocalFilesystem) PublicUrl(filePath string) string {
	return fmt.Sprintf(
		"%s/%s",
		strings.TrimRight(s.conf.File.Local.Domain, "/"),
		strings.Trim(filePath, "/"),
	)
}

func (s *LocalFilesystem) PrivateUrl(filePath string, timeout time.Duration) string {
	return ""
}

func (s *LocalFilesystem) ReadStream(filePath string) ([]byte, error) {
	return os.ReadFile(s.Path(filePath))
}

func (s *LocalFilesystem) InitiateMultipartUpload(_ string, _ string) (string, error) {
	str := fmt.Sprintf("%d%s", time.Now().Unix(), encrypt.Md5(strutil.Random(20)))

	return str, nil
}

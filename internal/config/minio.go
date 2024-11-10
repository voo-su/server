package config

type Minio struct {
	Host          string `yaml:"host"`
	SSL           bool   `yaml:"ssl"`
	SecretId      string `yaml:"secret_id"`
	SecretKey     string `yaml:"secret_key"`
	BucketPublic  string `yaml:"bucket_public"`
	BucketPrivate string `yaml:"bucket_private"`
}

func (m *Minio) GetHost() string {
	return m.Host
}

func (m *Minio) GetSSL() bool {
	return m.SSL
}

func (m *Minio) GetSecretId() string {
	return m.SecretId
}

func (m *Minio) GetSecretKey() string {
	return m.SecretKey
}

func (m *Minio) GetBucketPublic() string {
	return m.BucketPublic
}

func (m *Minio) GetBucketPrivate() string {
	return m.BucketPrivate
}

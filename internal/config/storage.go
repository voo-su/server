package config

type Minio struct {
	Endpoint      string `yaml:"endpoint"`
	SSL           bool   `yaml:"ssl"`
	SecretId      string `yaml:"secret_id"`
	SecretKey     string `yaml:"secret_key"`
	BucketPublic  string `yaml:"bucket_public"`
	BucketPrivate string `yaml:"bucket_private"`
}

package config

type Jwt struct {
	Secret      string `yaml:"secret"`
	ExpiresTime int64  `yaml:"expires_time"`
	BufferTime  int64  `yaml:"buffer_time"`
}

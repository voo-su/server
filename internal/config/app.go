package config

import "fmt"

type App struct {
	Env         string `yaml:"env"`
	DefaultLang string `yaml:"default_lang"`
	Jwt         *Jwt   `yaml:"jwt"`
	Cors        *Cors  `yaml:"cors"`
}

func (a App) LogPath(filename string) string {
	// TODO
	return fmt.Sprintf("/var/log/voo-su/%s", filename)
}

func (a *App) GetEnv() string {
	return a.Env
}
func (a *App) GetDefaultLang() string {
	return a.DefaultLang
}

type Cors struct {
	Origin      string `yaml:"origin"`
	Headers     string `yaml:"headers"`
	Methods     string `yaml:"methods"`
	Credentials string `yaml:"credentials"`
	MaxAge      string `yaml:"max_age"`
}

func (c *Cors) GetOrigin() string {
	return c.Origin
}

func (c *Cors) GetHeaders() string {
	return c.Headers
}

func (c *Cors) GetMethods() string {
	return c.Methods
}

func (c *Cors) GetCredentials() string {
	return c.Credentials
}

func (c *Cors) GetMaxAge() string {
	return c.MaxAge
}

type Jwt struct {
	Secret      string `yaml:"secret"`
	ExpiresTime int64  `yaml:"expires_time"`
	BufferTime  int64  `yaml:"buffer_time"`
}

func (j *Jwt) GetSecret() string {
	return j.Secret
}

func (j *Jwt) GetExpiresTime() int64 {
	return j.ExpiresTime
}

func (j *Jwt) GetBufferTime() int64 {
	return j.BufferTime
}

package config

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

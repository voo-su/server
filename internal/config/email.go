package config

type Email struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	Name     string `yaml:"name"`
	Report   string `yaml:"report"`
}

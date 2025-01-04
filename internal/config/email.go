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

func (e *Email) GetHost() string {
	return e.Host
}

func (e *Email) GetPort() int {
	return e.Port
}

func (e *Email) GetUserName() string {
	return e.UserName
}

func (e *Email) GetPassword() string {
	return e.Password
}

func (e *Email) GetFrom() string {
	return e.From
}

func (e *Email) GetName() string {
	return e.Name
}

func (e *Email) GetReport() string {
	return e.Report
}

package config

type LocalSystem struct {
	Root   string `yaml:"root"`
	Domain string `yaml:"domain"`
}

type File struct {
	Default string      `yaml:"default"`
	Local   LocalSystem `yaml:"local"`
}

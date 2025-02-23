package config

type Push struct {
	WebPush  WebPush  `yaml:"web_push"`
	Firebase Firebase `yaml:"firebase"`
}

type WebPush struct {
	PrivateKey string `yaml:"private_key"`
	PublicKey  string `yaml:"public_key"`
}

func (w *WebPush) GetPrivateKey() string {
	return w.PrivateKey
}

func (w *WebPush) GetPublicKey() string {
	return w.PublicKey
}

type Firebase struct {
	ProjectId string `yaml:"project_id"`
	JsonPath  string `yaml:"json_path"`
}

func (f *Firebase) GetProjectId() string {
	return f.ProjectId
}

func (f *Firebase) GetJsonPath() string {
	return f.JsonPath
}

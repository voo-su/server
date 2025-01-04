package config

type Push struct {
	WebPush WebPush `yaml:"web_push"`
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

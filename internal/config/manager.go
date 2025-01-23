package config

type Manager struct {
	Ips []string `yaml:"ips"`
}

func (m *Manager) GetIps() []string {
	return m.Ips
}

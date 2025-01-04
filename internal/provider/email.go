package provider

import (
	"voo.su/internal/config"
	"voo.su/pkg/email"
)

func NewEmailClient(conf *config.Config) *email.Email {
	return email.NewEmail(&email.Config{
		Host:     conf.Email.Host,
		Port:     conf.Email.Port,
		UserName: conf.Email.UserName,
		Password: conf.Email.Password,
		From:     conf.Email.From,
		Name:     conf.Email.Name,
	})
}

// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Conf *Config
}

type Config struct {
	Host     string
	Port     int
	UserName string
	Password string
	From     string
	Name     string
}

func NewEmail(config *Config) *Email {
	return &Email{
		Conf: config,
	}
}

type Option struct {
	To      string
	Subject string
	Body    string
}

type OptionFunc func(msg *gomail.Message)

func (e *Email) do(msg *gomail.Message) error {
	dialer := gomail.NewDialer(e.Conf.Host, e.Conf.Port, e.Conf.UserName, e.Conf.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer.DialAndSend(msg)
}

func (e *Email) SendMail(email *Option, opt ...OptionFunc) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(e.Conf.From, e.Conf.Name))

	if len(email.To) > 0 {
		m.SetHeader("To", email.To)
	}

	if len(email.Subject) > 0 {
		m.SetHeader("Subject", email.Subject)
	}

	if len(email.Body) > 0 {
		m.SetBody("text/html", email.Body)
	}

	for _, o := range opt {
		o(m)
	}

	return e.do(m)
}

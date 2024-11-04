package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Client struct {
	config *Config
}

type Config struct {
	Host     string
	Port     int
	UserName string
	Password string
	From     string
	Name     string
}

func NewEmail(config *Config) *Client {
	return &Client{
		config: config,
	}
}

type Option struct {
	To      string
	Subject string
	Body    string
}

type OptionFunc func(msg *gomail.Message)

func (c *Client) do(msg *gomail.Message) error {
	dialer := gomail.NewDialer(c.config.Host, c.config.Port, c.config.UserName, c.config.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer.DialAndSend(msg)
}

func (c *Client) SendMail(email *Option, opt ...OptionFunc) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(c.config.From, c.config.Name))

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

	return c.do(m)
}

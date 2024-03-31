package provider

import "voo.su/pkg/email"

type Providers struct {
	EmailClient *email.Client
}

func NewProviders(emailClient *email.Client) *Providers {
	return &Providers{EmailClient: emailClient}
}

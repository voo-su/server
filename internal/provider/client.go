package provider

import (
	"net/http"
	"time"
	"voo.su/pkg/client"
)

func NewHttpClient() *client.HttpClient {
	return client.NewHttpClient(&http.Client{
		Timeout: 15 * time.Second,
	})
}

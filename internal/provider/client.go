// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package provider

import (
	"net/http"
	"time"
	"voo.su/pkg/client"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}

func NewRequestClient(c *http.Client) *client.RequestClient {
	return client.NewRequestClient(c)
}

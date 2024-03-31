package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestClient struct {
	debug  bool
	client *http.Client
}

type FileData struct {
	Field    string
	FileName string
	Content  []byte
}

func NewRequestClient(client *http.Client) *RequestClient {
	return &RequestClient{client: client}
}

func (c *RequestClient) SetDebug() {
	c.debug = true
}

func (c *RequestClient) Get(url string, params *url.Values) ([]byte, error) {
	if params != nil {
		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&%s", url, params.Encode())
		} else {
			url = fmt.Sprintf("%s?%s", url, params.Encode())
		}
	}

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("\n[GET] HTTP Запрос\n")
		fmt.Printf("URL запроса: %s\n", url)
		fmt.Printf("Статусный код ответа: %d\n", resp.StatusCode)
		fmt.Printf("Данные ответа: %s\n\n", string(res))
	}

	return res, nil
}

func (c *RequestClient) Post(url string, params *url.Values) ([]byte, error) {
	resp, err := c.client.PostForm(url, *params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("\n[POST] HTTP Запрос\n")
		fmt.Printf("URL запроса: %s\n", url)
		fmt.Printf("Данные запроса: %s\n", params.Encode())
		fmt.Printf("Статусный код ответа: %d\n", resp.StatusCode)
		fmt.Printf("Данные ответа: %s\n\n", string(res))
	}

	return res, nil
}

func (c *RequestClient) PostJson(url string, params any) ([]byte, error) {
	text, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(text)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("\n[POST] HTTP Запрос\n")
		fmt.Printf("URL запроса: %s\n", url)
		fmt.Printf("Данные запроса: %s\n", string(text))
		fmt.Printf("Статусный код ответа: %d\n", resp.StatusCode)
		fmt.Printf("Данные ответа: %s\n\n", string(res))
	}

	return res, nil
}

func (c *RequestClient) PostFrom(url string, params *url.Values, files []*FileData) {

}

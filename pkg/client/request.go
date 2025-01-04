package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type RequestClient struct {
	Debug  bool
	Client *http.Client
}

type FileData struct {
	Field    string
	FileName string
	Content  []byte
}

func NewRequestClient(client *http.Client) *RequestClient {
	return &RequestClient{Client: client}
}

func (r *RequestClient) SetDebug() {
	r.Debug = true
}

func (r *RequestClient) Get(url string, params *url.Values) ([]byte, error) {
	if params != nil {
		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&%s", url, params.Encode())
		} else {
			url = fmt.Sprintf("%s?%s", url, params.Encode())
		}
	}

	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if r.Debug {
		log.Printf("\n[GET] HTTP Запрос\n")
		log.Printf("URL запроса: %s\n", url)
		log.Printf("Статусный код ответа: %d\n", resp.StatusCode)
		log.Printf("Данные ответа: %s\n\n", string(res))
	}

	return res, nil
}

func (r *RequestClient) Post(url string, params *url.Values) ([]byte, error) {
	resp, err := r.Client.PostForm(url, *params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if r.Debug {
		log.Printf("\n[POST] HTTP Запрос\n")
		log.Printf("URL запроса: %s\n", url)
		log.Printf("Данные запроса: %s\n", params.Encode())
		log.Printf("Статусный код ответа: %d\n", resp.StatusCode)
		log.Printf("Данные ответа: %s\n\n", string(res))
	}

	return res, nil
}

func (r *RequestClient) PostJson(url string, params any) ([]byte, error) {
	text, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(text)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if r.Debug {
		log.Printf("\n[POST] HTTP Запрос\n")
		log.Printf("URL запроса: %s\n", url)
		log.Printf("Данные запроса: %s\n", string(text))
		log.Printf("Статусный код ответа: %d\n", resp.StatusCode)
		log.Printf("Данные ответа: %s\n\n", string(res))
	}

	return res, nil
}

func (r *RequestClient) PostFrom(url string, params *url.Values, files []*FileData) {

}

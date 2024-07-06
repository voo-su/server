package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"voo.su/internal/config"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/client"
	"voo.su/pkg/sliceutil"
)

type IpAddressService struct {
	*repo.Source
	Config     *config.Config
	HttpClient *client.RequestClient
}

func NewIpAddressService(
	source *repo.Source,
	conf *config.Config,
	httpClient *client.RequestClient,
) *IpAddressService {
	return &IpAddressService{
		Source:     source,
		Config:     conf,
		HttpClient: httpClient,
	}
}

type IpAddressResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
	Country  string `json:"country,omitempty"`
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
	Isp      string `json:"isp,omitempty"`
}

func (i *IpAddressService) FindAddress(ip string) (string, error) {
	if val, err := i.getCache(ip); err == nil {
		return val, nil
	}

	_url := fmt.Sprintf("http://ip-api.com/json/%v", ip)
	resp, err := i.HttpClient.Get(_url, nil)
	if err != nil {
		return "", err
	}

	data := &IpAddressResponse{}
	if err := json.Unmarshal(resp, data); err != nil {
		return "", err
	}
	fmt.Println(data)
	if data.Status != "success" {
		return "", errors.New(data.Message)
	}

	arr := []string{data.Country, data.Province, data.City, data.Isp}
	val := strings.Join(sliceutil.Unique(arr), " ")
	val = strings.TrimSpace(val)
	_ = i.setCache(ip, val)

	return val, nil
}

func (i *IpAddressService) getCache(ip string) (string, error) {
	return i.Source.Redis().HGet(context.TODO(), "rds:hash:ip-address", ip).Result()
}

func (i *IpAddressService) setCache(ip string, value string) error {
	return i.Source.Redis().HSet(context.TODO(), "rds:hash:ip-address", ip, value).Err()
}

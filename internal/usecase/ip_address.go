package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"voo.su/internal/config"
	"voo.su/internal/repository"
	"voo.su/pkg/client"
	"voo.su/pkg/sliceutil"
)

type IpAddressUseCase struct {
	*repository.Source
	Config     *config.Config
	HttpClient *client.RequestClient
}

func NewIpAddressUseCase(
	source *repository.Source,
	conf *config.Config,
	httpClient *client.RequestClient,
) *IpAddressUseCase {
	return &IpAddressUseCase{
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

func (i *IpAddressUseCase) FindAddress(ip string) (string, error) {
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

	if data.Status != "success" {
		return "", errors.New(data.Message)
	}

	arr := []string{data.Country, data.Province, data.City, data.Isp}
	val := strings.Join(sliceutil.Unique(arr), " ")
	val = strings.TrimSpace(val)
	if err := i.setCache(ip, val); err != nil {
		fmt.Println(err)
	}

	return val, nil
}

func (i *IpAddressUseCase) getCache(ip string) (string, error) {
	return i.Source.Redis().HGet(context.TODO(), "rds:hash:ip-address", ip).Result()
}

func (i *IpAddressUseCase) setCache(ip string, value string) error {
	return i.Source.Redis().HSet(context.TODO(), "rds:hash:ip-address", ip, value).Err()
}

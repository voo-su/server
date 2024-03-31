package provider

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"voo.su/internal/config"
)

func NewNsqProducer(conf *config.Config) *nsq.Producer {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(conf.Nsq.Address, nsqConfig)
	if err != nil {
		panic(fmt.Errorf("не удалось создать, ошибка:%v\n", err))
	}

	return producer
}

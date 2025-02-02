package queue

import (
	"context"
	"encoding/json"
	"fmt"
	_nats "github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/domain/entity"
	"voo.su/pkg/logger"
	"voo.su/pkg/nats"
	webPush "voo.su/pkg/push/web_push"
)

type PushHandle struct {
	Conf *config.Config
	Nats nats.INatsClient
}

func (p *PushHandle) Handle(ctx context.Context) error {
	_, err := p.Nats.Subscribe("web-push", func(msg *_nats.Msg) {
		p.WebPush(string(msg.Data))
	})
	if err != nil {
		log.Println(err)
	}

	select {}
	return nil
}

func (p *PushHandle) WebPush(message string) {
	var in entity.WebPush
	if err := json.Unmarshal([]byte(message), &in); err != nil {
		fmt.Println(err)
		logger.Errorf("WebPush json decode err: %s", err.Error())
		return
	}

	subscription := &webPush.Subscription{
		Endpoint: in.Endpoint,
		Keys: webPush.Keys{
			P256dh: in.Keys.P256dh,
			Auth:   in.Keys.Auth,
		},
	}

	options := &webPush.Options{
		Topic:           "",
		Subscriber:      "",
		VAPIDPublicKey:  p.Conf.Push.WebPush.PublicKey,
		VAPIDPrivateKey: p.Conf.Push.WebPush.PrivateKey,
		TTL:             30,
	}

	res, err := webPush.SendNotification([]byte(in.Message), subscription, options)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	defer res.Body.Close()
}

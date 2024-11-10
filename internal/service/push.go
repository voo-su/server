package service

import (
	"fmt"
	_nats "github.com/nats-io/nats.go"
	"io/ioutil"
	"voo.su/pkg/nats"
	webPush "voo.su/pkg/push/web_push"
)

type PushService struct {
	Nats nats.INatsClient
}

func NewPushService() *PushService {
	return &PushService{}
}

func (p *PushService) Push() {
	if err := p.Nats.Publish("web-push", "Test"); err != nil {
		fmt.Println(err)
	}
}

func (p *PushService) Web() {
	_, err := p.Nats.Subscribe("web-push", func(msg *_nats.Msg) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		fmt.Println(err)
	}

	subscription := &webPush.Subscription{
		Endpoint: "",
		Keys: webPush.Keys{
			P256dh: "",
			Auth:   "",
		},
	}

	resp, err := webPush.SendNotification([]byte("Test"), subscription, &webPush.Options{
		Topic:           "",
		Subscriber:      "",
		VAPIDPublicKey:  "",
		VAPIDPrivateKey: "",
		TTL:             30,
	})
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	defer resp.Body.Close()
}

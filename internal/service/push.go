package service

import (
	"fmt"
	"io/ioutil"
	webPush "voo.su/pkg/push/web_push"
)

type PushService struct {
}

func NewPushService() *PushService {
	return &PushService{}
}

func (c *PushService) Generate() {
	privateKey, publicKey, err := webPush.GenerateVAPIDKeys()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(privateKey)
	fmt.Println(publicKey)
}

func (c *PushService) Web() {
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

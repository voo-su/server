package queue

import (
	"context"
	"encoding/json"
	"fmt"
	_nats "github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/nats"
	"voo.su/pkg/push"
	webPush "voo.su/pkg/push/web_push"
)

type PushHandle struct {
	Conf   *config.Config
	Nats   nats.INatsClient
	Source *infrastructure.Source
}

func (p *PushHandle) Handle(ctx context.Context) error {
	_, err := p.Nats.Subscribe(constant.QueuePush, func(msg *_nats.Msg) {
		p.Router(string(msg.Data))
	})
	if err != nil {
		log.Println(err)
	}

	select {}
	return nil
}

func (p *PushHandle) Router(message string) {
	var in entity.PushPayload
	if err := json.Unmarshal([]byte(message), &in); err != nil {
		log.Fatalf("Push json decode err: %s", err)
		return
	}

	uids := make([]int, 0)
	for _, uid := range in.UserIds {
		uids = append(uids, uid)
	}

	pushTokens := make([]*postgresModel.PushToken, 0)
	if err := p.Source.Postgres().
		Table("push_tokens").
		Where("user_id IN ?", uids).
		Scan(&pushTokens).
		Error; err != nil {
		fmt.Println(err)
	}

	for _, item := range pushTokens {
		switch item.Platform {
		case constant.PushPlatformWeb:
			p.WebPush(&entity.WebPush{
				Endpoint: item.WebEndpoint,
				Keys: entity.WebPushKeys{
					P256dh: item.WebP256dh,
					Auth:   item.WebAuth,
				},
				Message: in.Message,
			})
		case constant.PushPlatformMobile:
			p.MobilePush(&entity.MobilePush{
				Token:   item.Token,
				Message: in.Message,
			})
		case constant.PushPlatformDesktop:
			// TODO
		}
	}
}

func (p *PushHandle) WebPush(in *entity.WebPush) {
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

	defer res.Body.Close()

	fmt.Println(string(body))
}

func (p *PushHandle) MobilePush(in *entity.MobilePush) {
	firebase := push.NewFIREBASEPush(p.Conf.Push.Firebase.JsonPath, "VooSu", p.Conf.Push.Firebase.ProjectId, "")

	payloadInfo := &push.PayloadInfo{
		Title:   "",
		Content: in.Message,
		Badge:   1,
	}

	_, err := firebase.Push(in.Token, push.NewFIREBASEPayload(payloadInfo, ""))
	if err != nil {
		fmt.Println(err)
	}
}

package socket

import (
	"context"
	"errors"
	"log"
	"time"
	"voo.su/pkg/timeutil"
)

var ack *AckBuffer

type AckBuffer struct {
	TimeWheel *timeutil.SimpleTimeWheel[*AckBufferContent]
}

type AckBufferContent struct {
	Cid      int64
	Uid      int64
	Channel  string
	Response *ClientResponse
}

func InitAck() {
	ack = &AckBuffer{}
	ack.TimeWheel = timeutil.NewSimpleTimeWheel[*AckBufferContent](1*time.Second, 30, ack.handle)
}

func (a *AckBuffer) Start(ctx context.Context) error {
	go a.TimeWheel.Start()
	<-ctx.Done()
	a.TimeWheel.Stop()
	return errors.New("сервис подтверждений остановлен")
}

func (a *AckBuffer) insert(ackKey string, value *AckBufferContent) {
	a.TimeWheel.Add(ackKey, value, time.Duration(5)*time.Second)
}

func (a *AckBuffer) delete(ackKey string) {
	a.TimeWheel.Remove(ackKey)
}

func (a *AckBuffer) handle(_ *timeutil.SimpleTimeWheel[*AckBufferContent], _ string, bufferContent *AckBufferContent) {
	ch, ok := Session.Channel(bufferContent.Channel)
	if !ok {
		return
	}

	client, ok := ch.Client(bufferContent.Cid)
	if !ok {
		return
	}

	if client.Closed() || int64(client.uid) != bufferContent.Uid {
		return
	}

	if err := client.Write(bufferContent.Response); err != nil {
		log.Println("ошибка: ", err)
	}
}

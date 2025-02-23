package push

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

// FIREBASEPush https://github.com/firebase/firebase-admin-go/blob/61c6c041bf807c045f6ff3fd0d02fc480f806c9a/snippets/messaging.go#L29-L55

type FIREBASEPush struct {
	jsonPath    string
	packageName string
	projectId   string
	channelID   string
	client      messaging.Client
}

func NewFIREBASEPush(jsonPath string, packageName string, projectID string, channelID string) *FIREBASEPush {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: projectID,
	}
	opt := option.WithCredentialsFile(jsonPath)
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Println("Не удалось инициализировать Firebase: ошибка при создании клиента через json")
		return nil
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Printf("Ошибка при создании message client через App client: %s", err)
		return nil
	}

	return &FIREBASEPush{
		jsonPath:    jsonPath,
		packageName: packageName,
		channelID:   channelID,
		client:      *client,
		projectId:   projectID,
	}
}

type FIREBASEPayload struct {
	Payload
	notifyID string
}

func NewFIREBASEPayload(payloadInfo *PayloadInfo, notifyID string) *FIREBASEPayload {
	return &FIREBASEPayload{
		Payload:  payloadInfo.toPayload(),
		notifyID: notifyID,
	}
}

func (f *FIREBASEPush) GetPayload() (Payload, error) {
	payloadInfo, err := ParsePushInfo()
	if err != nil {
		return nil, err
	}

	return NewFIREBASEPayload(payloadInfo, "11"), nil
}

func (f *FIREBASEPush) Push(deviceToken string, payload Payload) (*string, error) {
	ctx := context.Background()
	miPayload := payload.(*FIREBASEPayload)
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: miPayload.GetTitle(),
			Body:  miPayload.GetContent(),
		},
		Token: deviceToken,
	}

	res, err := f.client.Send(ctx, message)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

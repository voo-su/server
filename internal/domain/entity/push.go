package entity

type PushPayload struct {
	UserIds []int
	Message string
}

type WebPushKeys struct {
	P256dh string
	Auth   string
}

type WebPush struct {
	Endpoint string
	Keys     WebPushKeys
	Message  string
}

type MobilePush struct {
	Token   string
	Message string
}

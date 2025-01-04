package entity

type WebPushKeys struct {
	P256dh string
	Auth   string
}

type WebPush struct {
	Endpoint string
	Keys     WebPushKeys
	Message  string
}

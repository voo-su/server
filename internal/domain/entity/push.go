// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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

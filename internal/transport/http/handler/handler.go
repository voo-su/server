package handler

import "voo.su/internal/transport/http/handler/v1"

type V1 struct {
	Auth             *v1.Auth
	Account          *v1.Account
	Contact          *v1.Contact
	ContactRequest   *v1.ContactRequest
	Chat             *v1.Chat
	Message          *v1.Message
	MessagePublish   *v1.Publish
	Upload           *v1.Upload
	GroupChat        *v1.GroupChat
	GroupChatRequest *v1.GroupChatRequest
	Sticker          *v1.Sticker
	ContactFolder    *v1.ContactFolder
	GroupChatAds     *v1.GroupChatAds
	Search           *v1.Search
}

type Handler struct {
	V1 *V1
}

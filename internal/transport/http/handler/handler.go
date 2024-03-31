package handler

import "voo.su/internal/transport/http/handler/v1"

type V1 struct {
	Auth             *v1.Auth
	Account          *v1.Account
	User             *v1.User
	Contact          *v1.Contact
	ContactRequest   *v1.ContactRequest
	Dialog           *v1.Dialog
	Message          *v1.Message
	MessagePublish   *v1.Publish
	Upload           *v1.Upload
	GroupChat        *v1.GroupChat
	GroupChatRequest *v1.GroupChatRequest
}

type Handler struct {
	V1 *V1
}

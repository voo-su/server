package v1

type Handler struct {
	Auth               *Auth
	Account            *Account
	Contact            *Contact
	ContactRequest     *ContactRequest
	Chat               *Chat
	Message            *Message
	Upload             *Upload
	GroupChat          *GroupChat
	GroupChatRequest   *GroupChatRequest
	Sticker            *Sticker
	ContactFolder      *ContactFolder
	GroupChatAds       *GroupChatAds
	Search             *Search
	Bot                *Bot
	Project            *Project
	ProjectTask        *ProjectTask
	ProjectTaskComment *ProjectTaskComment
}

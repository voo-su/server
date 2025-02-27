package entity

type NotifySettings struct {
	ChatsMuteUntil    int32
	ChatsShowPreviews bool
	ChatsSilent       bool
	GroupMuteUntil    int32
	GroupShowPreviews bool
	GroupSilent       bool
}

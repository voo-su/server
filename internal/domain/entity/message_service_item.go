package entity

import (
	"time"
	"voo.su/internal/constant"
)

type ServiceType int

const (
	ServiceTypeMsgSysText                ServiceType = constant.ChatMsgSysText
	ServiceTypeChatMsgTypeLogin          ServiceType = constant.ChatMsgTypeLogin
	ServiceTypeChatMsgSysGroupCreate     ServiceType = constant.ChatMsgSysGroupCreate
	ServiceTypeChatMsgSysGroupUserInvite ServiceType = constant.ChatMsgSysGroupUserInvite
	ServiceTypeChatMsgSysGroupUserRemove ServiceType = constant.ChatMsgSysGroupUserRemove
)

type ServiceItem struct {
	Type      ServiceType        `json:"type"`
	Action    ServiceItemMessage `json:"action,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
}

type ServiceItemMessage interface {
	MessageType() ServiceType
}

type ServiceLogin struct {
	Ip      string  `json:"ip,omitempty"`
	Agent   string  `json:"agent,omitempty"`
	Address *string `json:"address,omitempty"`
}

func (s ServiceLogin) MessageType() ServiceType {
	return ServiceTypeChatMsgTypeLogin
}

type ServiceItemMember struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type ServiceItemGroupCancelMutedMessage struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

func (s ServiceItemGroupCancelMutedMessage) MessageType() ServiceType {
	return ServiceTypeMsgSysText
}

type ServiceItemGroupCreateMessage struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

func (s ServiceItemGroupCreateMessage) MessageType() ServiceType {
	return ServiceTypeChatMsgSysGroupCreate
}

type ServiceItemGroupJoinMessage struct {
	OwnerId   int               `json:"owner_id"`
	OwnerName string            `json:"owner_name"`
	Member    ServiceItemMember `json:"member"`
}

func (s ServiceItemGroupJoinMessage) MessageType() ServiceType {
	return ServiceTypeChatMsgSysGroupCreate
}

type ServiceItemGroupMemberCancelMutedMessage struct {
	OwnerId   int               `json:"owner_id"`
	OwnerName string            `json:"owner_name"`
	Member    ServiceItemMember `json:"member"`
}

func (s ServiceItemGroupMemberCancelMutedMessage) MessageType() ServiceType {
	return ServiceTypeMsgSysText
}

type ServiceItemGroupMemberKickedMessage struct {
	OwnerId   int               `json:"owner_id"`
	OwnerName string            `json:"owner_name"`
	Member    ServiceItemMember `json:"member"`
}

func (s ServiceItemGroupMemberKickedMessage) MessageType() ServiceType {
	return ServiceTypeChatMsgSysGroupUserRemove
}

type ServiceItemGroupMemberMutedMessage struct {
	OwnerId   int               `json:"owner_id"`
	OwnerName string            `json:"owner_name"`
	Member    ServiceItemMember `json:"member"`
}

func (s ServiceItemGroupMemberMutedMessage) MessageType() ServiceType {
	return ServiceTypeMsgSysText
}

type ServiceItemGroupMemberQuitMessage struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

func (s ServiceItemGroupMemberQuitMessage) MessageType() ServiceType {
	return ServiceTypeMsgSysText
}

type ServiceItemGroupMutedMessage struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

func (s ServiceItemGroupMutedMessage) MessageType() ServiceType {
	return ServiceTypeMsgSysText
}

type ServiceItemGroupTransferMessage struct {
	OldOwnerId   int    `json:"old_owner_id"`
	OldOwnerName string `json:"old_owner_name"`
	NewOwnerId   int    `json:"new_owner_id"`
	NewOwnerName string `json:"new_owner_name"`
}

func (s ServiceItemGroupTransferMessage) MessageType() ServiceType {
	return ServiceTypeMsgSysText
}

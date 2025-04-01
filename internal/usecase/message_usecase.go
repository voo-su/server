package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"html"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/domain/logic"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisModel "voo.su/internal/infrastructure/redis/model"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/minio"
	"voo.su/pkg/nats"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type IMessageUseCase interface {
	IsAccess(ctx context.Context, opt *entity.MessageAccess) error

	GetHistory(ctx context.Context, opt *entity.QueryGetHistoryOpt) ([]*entity.MessageItem, error)

	GetMessage(ctx context.Context, messageId int64) (*entity.MessageItem, error)

	GetMessageByMessageId(ctx context.Context, messageId int64) (*postgresModel.Message, error)

	SendText(ctx context.Context, uid int, req *entity.SendText) error

	SendSystemText(ctx context.Context, uid int, req *entity.TextMessageRequest) error

	SendImage(ctx context.Context, uid int, req *entity.SendImage) error

	SendVideo(ctx context.Context, uid int, req *entity.SendVideo) error

	SendAudio(ctx context.Context, uid int, req *entity.SendAudio) error

	SendFile(ctx context.Context, uid int, req *entity.SendFile) error

	SendBotFile(ctx context.Context, uid int, req *entity.SendBotFile) error

	SendVote(ctx context.Context, uid int, req *v1Pb.VoteMessageRequest) error

	SendForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) error

	SendSysOther(ctx context.Context, data *postgresModel.Message) error

	SendMixedMessage(ctx context.Context, uid int, req *v1Pb.MixedMessageRequest) error

	SendLogin(ctx context.Context, uid int, req *entity.SendLogin) error

	SendSticker(ctx context.Context, uid int, req *v1Pb.StickerMessageRequest) error

	SendCode(ctx context.Context, uid int, req *v1Pb.CodeMessageRequest) error

	SendLocation(ctx context.Context, uid int, req *v1Pb.LocationMessageRequest) error

	Vote(ctx context.Context, uid int, msgId int64, optionsValue string) (*postgresRepo.VoteStatistics, error)

	Revoke(ctx context.Context, uid int, msgId int64) error
}

var _ IMessageUseCase = (*MessageUseCase)(nil)

type MessageUseCase struct {
	Conf                 *config.Config
	Locale               locale.ILocale
	Source               *infrastructure.Source
	MessageForward       *logic.MessageForward
	Minio                minio.IMinio
	GroupChatMemberRepo  *postgresRepo.GroupChatMemberRepository
	FileSplitRepo        *postgresRepo.FileSplitRepository
	Sequence             *postgresRepo.SequenceRepository
	UserRepo             *postgresRepo.UserRepository
	MessageRepo          *postgresRepo.MessageRepository
	MessageVoteRepo      *postgresRepo.MessageVoteRepository
	MessageLoginRepo     *postgresRepo.MessageLoginRepository
	MessageCodeRepo      *postgresRepo.MessageCodeRepository
	MessageLocationRepo  *postgresRepo.MessageLocationRepository
	MessageForwardedRepo *postgresRepo.MessageForwardedRepository
	BotRepo              *postgresRepo.BotRepository
	GroupChatRepo        *postgresRepo.GroupChatRepository
	ContactRepo          *postgresRepo.ContactRepository
	UnreadCache          *redisRepo.UnreadCacheRepository
	MessageCache         *redisRepo.MessageCacheRepository
	ServerCache          *redisRepo.ServerCacheRepository
	ClientCache          *redisRepo.ClientCacheRepository
	Nats                 nats.INatsClient
}

func (m *MessageUseCase) IsAccess(ctx context.Context, opt *entity.MessageAccess) error {
	if opt.ChatType == constant.ChatPrivateMode {
		if m.ContactRepo.IsFriend(ctx, opt.UserId, opt.ReceiverId, false) {
			return nil
		}
		return errors.New(m.Locale.Localize("no_permission_to_send_messages"))
	}

	groupInfo, err := m.GroupChatRepo.FindById(ctx, opt.ReceiverId)
	if err != nil {
		return err
	}

	if groupInfo.IsDismiss == 1 {
		return errors.New(m.Locale.Localize("group_deleted"))
	}

	memberInfo, err := m.GroupChatMemberRepo.FindByUserId(ctx, opt.ReceiverId, opt.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(m.Locale.Localize("no_permission_to_send_messages"))
		}
		return errors.New(m.Locale.Localize("system_busy_try_later"))
	}

	if memberInfo.IsQuit == constant.GroupMemberQuitStatusYes {
		return errors.New(m.Locale.Localize("no_permission_to_send_messages"))
	}

	if memberInfo.IsMute == constant.GroupMemberMuteStatusYes {
		return errors.New(m.Locale.Localize("message_sending_prohibited_by_admin"))
	}

	if opt.IsVerifyGroupMute && groupInfo.IsMute == 1 && memberInfo.Leader == 0 {
		return errors.New(m.Locale.Localize("group_message_sending_disabled"))
	}

	return nil
}

func (m *MessageUseCase) GetHistory(ctx context.Context, opt *entity.QueryGetHistoryOpt) ([]*entity.MessageItem, error) {
	var (
		fields = []string{
			"messages.id",
			"messages.sequence",
			"messages.chat_type",
			"messages.msg_type",
			"messages.user_id",
			"messages.receiver_id",
			"messages.is_revoke",
			"messages.is_read",
			"messages.content",
			"messages.extra",
			"messages.quote_id",
			"messages.created_at",
		}
		items = make([]*entity.QueryMessageItem, 0, opt.Limit)
	)
	query := m.Source.Postgres().WithContext(ctx).Select(fields).
		Table("messages").
		Joins("LEFT JOIN message_delete ON messages.id = message_delete.message_id AND message_delete.user_id = ?", opt.UserId)
	if opt.MessageId > 0 {
		query.Where("messages.sequence < ?", opt.MessageId)
	}

	if opt.ChatType == constant.ChatPrivateMode {
		subQuery := m.Source.Postgres().
			Where("messages.user_id = ? AND messages.receiver_id = ?", opt.UserId, opt.ReceiverId).
			Or("messages.user_id = ? AND messages.receiver_id = ?", opt.ReceiverId, opt.UserId)
		query.Where(subQuery)
	} else {
		query.Where("messages.receiver_id = ?", opt.ReceiverId)
	}

	if opt.MsgType != nil && len(opt.MsgType) > 0 {
		query.Where("messages.msg_type in ?", opt.MsgType)
	}

	query.Where("messages.chat_type = ?", opt.ChatType).
		Where("COALESCE(message_delete.id,0) = 0").
		Order("messages.sequence desc").
		Limit(opt.Limit)

	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return make([]*entity.MessageItem, 0), nil
	}

	return m.HandleMessages(ctx, items)
}

func (m *MessageUseCase) GetMessage(ctx context.Context, messageId int64) (*entity.MessageItem, error) {
	var (
		fields = []string{
			"messages.id",
			"messages.sequence",
			"messages.chat_type",
			"messages.msg_type",
			"messages.user_id",
			"messages.receiver_id",
			"messages.is_revoke",
			"messages.content",
			"messages.extra",
			"messages.created_at",
		}
		err  error
		item *entity.QueryMessageItem
	)
	query := m.Source.Postgres().
		Table("messages").
		Where("messages.id = ?", messageId)
	if err = query.Select(fields).Take(&item).Error; err != nil {
		return nil, err
	}

	list, err := m.HandleMessages(ctx, []*entity.QueryMessageItem{item})
	if err != nil {
		return nil, err
	}

	return list[0], nil
}

func (m *MessageUseCase) GetForwardMessages(ctx context.Context, uid int, messageId int64) ([]*entity.MessageItem, error) {
	record, err := m.MessageRepo.FindById(ctx, int(messageId))
	if err != nil {
		return nil, err
	}

	//if record.ChatType == domain.ChatPrivateMode {
	//	if record.UserId != uid && record.ReceiverId != uid {
	//		return nil, domain.ErrPermissionDenied
	//	}
	//} else if record.ChatType == domain.ChatGroupMode {
	//	if !s.GroupMemberRepo.IsMember(ctx, record.ReceiverId, uid, true) {
	//		return nil, domain.ErrPermissionDenied
	//	}
	//} else {
	//	return nil, domain.ErrPermissionDenied
	//}

	var extra entity.MessageExtraForward
	if err := jsonutil.Decode(record.Extra, &extra); err != nil {
		return nil, err
	}
	var (
		items  = make([]*entity.QueryMessageItem, 0)
		fields = []string{
			"messages.id",
			"messages.sequence",
			"messages.chat_type",
			"messages.msg_type",
			"messages.user_id",
			"messages.receiver_id",
			"messages.is_revoke",
			"messages.content",
			"messages.extra",
			"messages.created_at",
		}
	)
	tx := m.Source.Postgres().Select(fields).
		Table("messages").
		Where("messages.id IN ?", extra.Ids).
		Order("messages.sequence ASC")
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return m.HandleMessages(ctx, items)
}

func (m *MessageUseCase) HandleMessages(ctx context.Context, items []*entity.QueryMessageItem) ([]*entity.MessageItem, error) {
	var (
		ids       []int
		uIds      []int
		mediaId   []int
		loginId   []int
		voteId    []int
		systemIds []int
	)

	for _, item := range items {
		ids = append(ids, item.Id)
		uIds = append(uIds, item.UserId)

		switch item.MsgType {
		case constant.ChatMsgTypeImage,
			constant.ChatMsgTypeVideo,
			constant.ChatMsgTypeAudio,
			constant.ChatMsgTypeFile:
			mediaId = append(mediaId, item.Id)

		case constant.ChatMsgTypeLogin:
			loginId = append(loginId, item.Id)

		case constant.ChatMsgTypeVote:
			voteId = append(voteId, item.Id)

		case constant.ChatMsgSysGroupCreate,
			constant.ChatMsgSysGroupMemberJoin,
			constant.ChatMsgSysGroupMemberQuit,
			constant.ChatMsgSysGroupMemberKicked,
			constant.ChatMsgSysGroupMessageRevoke,
			constant.ChatMsgSysGroupDismissed,
			constant.ChatMsgSysGroupMuted,
			constant.ChatMsgSysGroupCancelMuted,
			constant.ChatMsgSysGroupMemberMuted,
			constant.ChatMsgSysGroupMemberCancelMuted,
			constant.ChatMsgSysGroupAds:
			systemIds = append(systemIds, item.Id)
		}
	}

	hashSystem := make(map[int]*postgresModel.MessageSystem)
	if len(systemIds) > 0 {
		var systemItems []*postgresModel.MessageSystem
		if err := m.Source.Postgres().
			Model(&postgresModel.MessageSystem{}).
			Where("message_id IN ?", systemIds).
			Scan(&systemItems).
			Error; err != nil {
			log.Printf("Error - IMessageUseCase - HandleMessages.MessageSystem: %v", err)
		}
		for i := range systemItems {
			uIds = append(uIds, systemItems[i].UserId, systemItems[i].TargetUserId)
			hashSystem[systemItems[i].MessageId] = systemItems[i]
		}
	}

	hashVotes := make(map[int]*postgresModel.MessageVote)
	if len(voteId) > 0 {
		var voteItems []*postgresModel.MessageVote
		if err := m.Source.Postgres().
			Model(&postgresModel.MessageVote{}).
			Where("message_id IN ?", voteId).
			Scan(&voteItems).Error; err != nil {
			log.Printf("Error - IMessageUseCase - HandleMessages.MessageVote: %v", err)
		}

		for i := range voteItems {
			hashVotes[voteItems[i].MessageId] = voteItems[i]
		}
	}

	hashLogin := make(map[int]*postgresModel.MessageLogin)
	if len(loginId) > 0 {
		var loginItems []*postgresModel.MessageLogin
		if err := m.Source.Postgres().
			Model(&postgresModel.MessageLogin{}).
			Where("message_id IN ?", loginId).
			Scan(&loginItems).
			Error; err != nil {
			log.Printf("Error - IMessageUseCase - HandleMessages.MessageLogin: %v", err)
		}

		for i := range loginItems {
			hashLogin[loginItems[i].MessageId] = loginItems[i]
		}
	}

	hashMedia := make(map[int]*postgresModel.MessageMedia)
	if len(mediaId) > 0 {
		var mediaItems []*postgresModel.MessageMedia
		if err := m.Source.Postgres().
			Model(&postgresModel.MessageMedia{}).
			Where("message_id IN ?", mediaId).
			Scan(&mediaItems).
			Error; err != nil {
			log.Printf("Error - IMessageUseCase - HandleMessages.MessageMedia: %v", err)
		}
		for i := range mediaItems {
			hashMedia[mediaItems[i].MessageId] = mediaItems[i]
		}
	}

	var forwardedMessages []*entity.MessageReply
	if err := m.Source.Postgres().Select([]string{
		"messages.id AS id",
		"messages.user_id AS user_id",
		"messages.content AS content",
		"messages.msg_type AS msg_type",
		"message_forwarded.new_message_id AS new_message_id",
	}).
		Table("message_forwarded").
		Joins("LEFT JOIN messages on message_forwarded.original_message_id = messages.id").
		Where("message_forwarded.new_message_id IN ?", ids).
		Scan(&forwardedMessages).
		Error; err != nil {
		log.Printf("Error - IMessageUseCase - HandleMessages.MessageReply: %v", err)
	}

	forwardedMap := make(map[int]*entity.MessageReply)
	for i := range forwardedMessages {
		uIds = append(uIds, forwardedMessages[i].UserId)
		forwardedMap[forwardedMessages[i].NewMessageId] = forwardedMessages[i]
	}

	hashUsers := make(map[int]*postgresModel.User)
	if len(uIds) > 0 {
		var users []*postgresModel.User
		if err := m.Source.Postgres().
			Where("id IN ?", uIds).
			Find(&users).
			Error; err != nil {
			log.Printf("Error - IMessageUseCase - HandleMessages.MessageMedia: %v", err)
		}

		for i := range users {
			hashUsers[users[i].Id] = users[i]
		}
	}

	newItems := make([]*entity.MessageItem, 0, len(items))
	for _, item := range items {
		data := &entity.MessageItem{
			Id:         item.Id,
			Sequence:   item.Sequence,
			ChatType:   item.ChatType,
			MsgType:    item.MsgType,
			UserId:     item.UserId,
			ReceiverId: item.ReceiverId,
			IsRevoke:   item.IsRevoke,
			IsMark:     item.IsMark,
			IsRead:     item.IsRead,
			CreatedAt:  timeutil.FormatDatetime(item.CreatedAt),
		}

		if u, ok := hashUsers[item.UserId]; ok {
			data.Username = u.Username
			data.Name = u.Name
			data.Surname = u.Surname
			data.Avatar = u.Avatar
		}

		if item.IsRevoke == 0 {
			data.Content = item.Content
		}

		if fw, ok := forwardedMap[item.Id]; ok {
			reply := &entity.Reply{
				Id:      fw.Id,
				Content: fw.Content,
				MsgType: fw.MsgType,
			}

			if u, ok := hashUsers[fw.UserId]; ok {
				reply.Username = u.Username
				reply.UserId = u.Id
			}

			data.Reply = reply
		}

		switch item.MsgType {
		case constant.ChatMsgTypeImage:
			if val, ok := hashMedia[item.Id]; ok {
				data.Media = &entity.MediaItem{
					Type:   entity.MediaTypeImage,
					FileId: val.FileId,
					Url:    val.Url,
					Width:  val.Width,
					Height: val.Height,
				}
			}
		case constant.ChatMsgTypeVideo:
			if val, ok := hashMedia[item.Id]; ok {
				data.Media = &entity.MediaItem{
					Type:     entity.MediaTypeVideo,
					FileId:   val.FileId,
					Url:      val.Url,
					Size:     val.Size,
					Cover:    val.Cover,
					MimeType: val.MimeType,
					Duration: val.Duration,
				}
			}
		case constant.ChatMsgTypeAudio:
			if val, ok := hashMedia[item.Id]; ok {
				data.Media = &entity.MediaItem{
					Type:     entity.MediaTypeAudio,
					FileId:   val.FileId,
					Url:      val.Url,
					Name:     val.Name,
					Size:     val.Size,
					MimeType: val.MimeType,
					Duration: val.Duration,
				}
			}

		case constant.ChatMsgTypeFile:
			if val, ok := hashMedia[item.Id]; ok {
				data.Media = &entity.MediaItem{
					Type:     entity.MediaTypeFile,
					FileId:   val.FileId,
					Url:      val.Url,
					Name:     val.Name,
					Size:     val.Size,
					MimeType: val.MimeType,
					Drive:    val.Drive,
				}
			}
		case constant.ChatMsgTypeVote:
			if value, ok := hashVotes[item.Id]; ok {
				options := make(map[string]any)
				opts := make([]any, 0)
				if err := jsonutil.Decode(value.AnswerOption, &options); err == nil {
					arr := make([]string, 0, len(options))
					for k := range options {
						arr = append(arr, k)
					}
					sort.Strings(arr)
					for _, v := range arr {
						opts = append(opts, map[string]any{
							"key":   v,
							"value": options[v],
						})
					}
				}

				users := make([]int, 0)
				if uids, err := m.MessageVoteRepo.GetVoteAnswerUser(ctx, value.Id); err == nil {
					users = uids
				}

				var statistics any
				if res, err := m.MessageVoteRepo.GetVoteStatistics(ctx, value.Id); err != nil {
					statistics = map[string]any{
						"count":   0,
						"options": map[string]int{},
					}
				} else {
					statistics = res
				}

				data.Media = &entity.MediaItem{
					Type: entity.MediaTypeVote,
					Detail: entity.DetailVote{
						Id:           value.Id,
						MessageId:    value.MessageId,
						Title:        value.Title,
						AnswerMode:   value.AnswerMode,
						Status:       value.Status,
						AnswerOption: opts,
						AnswerNum:    value.AnswerNum,
						AnsweredNum:  value.AnsweredNum,
					},
					Statistics: statistics,
					VoteUsers:  users,
				}
			}

		case constant.ChatMsgTypeLogin:
			if val, ok := hashLogin[item.Id]; ok {
				data.Service = &entity.ServiceItem{
					Type:      entity.ServiceTypeChatMsgTypeLogin,
					CreatedAt: val.CreatedAt,
					Action: entity.ServiceLogin{
						Ip:      val.IpAddress,
						Agent:   val.UserAgent,
						Address: val.Address,
					},
				}
			}

		case constant.ChatMsgSysGroupCreate:
			if val, ok := hashSystem[item.Id]; ok {
				action := entity.ServiceItemGroupCreateMessage{}

				if u, ok := hashUsers[val.UserId]; ok {
					action.OwnerId = u.Id
					action.OwnerName = u.Name
				}

				data.Service = &entity.ServiceItem{
					Type:      entity.ServiceTypeChatMsgSysGroupCreate,
					CreatedAt: val.CreatedAt,
					Action:    action,
				}
			}

		case constant.ChatMsgSysGroupMemberJoin:
			if val, ok := hashSystem[item.Id]; ok {
				member := entity.ServiceItemMember{}
				if u, ok := hashUsers[val.TargetUserId]; ok {
					member = entity.ServiceItemMember{
						UserId:   u.Id,
						Username: u.Username,
						Name:     u.Name,
					}
				}
				action := entity.ServiceItemGroupJoinMessage{
					Member: member,
				}

				if u, ok := hashUsers[val.UserId]; ok {
					action.OwnerId = u.Id
					action.OwnerName = u.Name
				}

				data.Service = &entity.ServiceItem{
					Type:      entity.ServiceTypeChatMsgSysGroupMemberJoin,
					CreatedAt: val.CreatedAt,
					Action:    action,
				}
			}

			// TODO
		case constant.ChatMsgSysGroupMemberQuit:
			// TODO

		case constant.ChatMsgSysGroupMemberKicked:
			if val, ok := hashSystem[item.Id]; ok {
				member := entity.ServiceItemMember{}
				if u, ok := hashUsers[val.TargetUserId]; ok {
					member = entity.ServiceItemMember{
						UserId:   u.Id,
						Username: u.Username,
						Name:     u.Name,
					}
				}
				action := entity.ServiceItemGroupMemberKickedMessage{
					Member: member,
				}

				if u, ok := hashUsers[val.UserId]; ok {
					action.OwnerId = u.Id
					action.OwnerName = u.Name
				}

				data.Service = &entity.ServiceItem{
					Type:      entity.ServiceTypeChatMsgSysGroupMemberKicked,
					CreatedAt: val.CreatedAt,
					Action:    action,
				}
			}

		case constant.ChatMsgSysGroupMessageRevoke:
			// TODO
		case constant.ChatMsgSysGroupDismissed:
			// TODO
		case constant.ChatMsgSysGroupMuted:
			// TODO
		case constant.ChatMsgSysGroupCancelMuted:
			// TODO
		case constant.ChatMsgSysGroupMemberMuted:
			// TODO
		case constant.ChatMsgSysGroupMemberCancelMuted:
			// TODO
		case constant.ChatMsgSysGroupAds:
			// TODO

		}
		newItems = append(newItems, data)
	}

	return newItems, nil
}

func (m *MessageUseCase) GetMessageByMessageId(ctx context.Context, messageId int64) (*postgresModel.Message, error) {
	record, err := m.MessageRepo.FindById(ctx, int(messageId))
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (m *MessageUseCase) SendText(ctx context.Context, uid int, req *entity.SendText) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeText,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    strutil.EscapeHtml(req.Content),
		ReplyTo:    int(req.ReplyToMsgId),
	})

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendSystemText(ctx context.Context, uid int, req *entity.TextMessageRequest) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgSysText,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    html.EscapeString(req.Content),
	})

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendImage(ctx context.Context, uid int, req *entity.SendImage) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeImage,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		ReplyTo:    int(req.ReplyToMsgId),
		Content:    req.Content,
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageMedia{
		MessageId: message.Id,
		FileId:    req.FileId,
		Url:       req.Url,
		Width:     req.Width,
		Height:    req.Height,
		CreatedAt: time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendImage.MessageMedia: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendVideo(ctx context.Context, uid int, req *entity.SendVideo) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeVideo,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    req.Content,
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageMedia{
		MessageId: message.Id,
		FileId:    req.FileId,
		Cover:     req.Cover,
		MimeType:  "video/" + strutil.FileSuffix(req.Url),
		Size:      int(req.Size),
		Url:       req.Url,
		Duration:  int(req.Duration),
		CreatedAt: time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendVideo.MessageMedia: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendAudio(ctx context.Context, uid int, req *entity.SendAudio) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeAudio,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    req.Content,
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageMedia{
		MessageId: message.Id,
		FileId:    req.FileId,
		MimeType:  "audio/" + strutil.FileSuffix(req.Url),
		Size:      int(req.Size),
		Url:       req.Url,
		CreatedAt: time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendAudio.MessageMedia: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendFile(ctx context.Context, uid int, req *entity.SendFile) error {
	file, err := m.FileSplitRepo.GetFile(ctx, uid, req.UploadId)
	if err != nil {
		return err
	}

	filePath := strutil.GenMediaObjectName(file.FileExt, 0, 0)
	publicUrl := ""
	if entity.GetMediaType(file.FileExt) <= 3 {
		if err := m.Minio.CopyObject(
			m.Conf.Minio.GetBucket(), file.Path,
			m.Conf.Minio.GetBucket(), filePath,
		); err != nil {
			return err
		}

		publicUrl = m.Minio.PublicUrl(m.Conf.Minio.GetBucket(), filePath)
	} else {
		if err := m.Minio.Copy(m.Conf.Minio.GetBucket(), file.Path, filePath); err != nil {
			return err
		}
	}

	data := &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	media := &postgresModel.MessageMedia{}

	switch entity.GetMediaType(file.FileExt) {
	case constant.MediaFileAudio:
		data.MsgType = constant.ChatMsgTypeAudio
		media = &postgresModel.MessageMedia{
			MimeType: "audio/" + file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
		}
	case constant.MediaFileVideo:
		data.MsgType = constant.ChatMsgTypeVideo
		media = &postgresModel.MessageMedia{
			MimeType: "video/" + file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
		}
	case constant.MediaFileOther:
		media = &postgresModel.MessageMedia{
			Drive:    file.Drive,
			Name:     file.OriginalName,
			MimeType: "application/" + file.FileExt,
			Size:     int(file.FileSize),
			Url:      filePath,
		}
		data.MsgType = constant.ChatMsgTypeFile
	}

	message, err := m.save(ctx, data)

	media.MessageId = message.Id

	if err := m.Source.Postgres().WithContext(ctx).Create(media).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendFile.MessageMedia: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendBotFile(ctx context.Context, uid int, req *entity.SendBotFile) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		MsgType:    constant.ChatMsgTypeFile,
		Content:    req.Content,
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageMedia{
		MessageId: message.Id,
		FileId:    req.FileId,
		Drive:     req.Drive,
		Name:      req.OriginalName,
		MimeType:  "application/" + req.FileExt,
		Size:      req.FileSize,
		Url:       req.FilePath,
		CreatedAt: time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendBotFile.MessageMedia: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendVote(ctx context.Context, uid int, req *v1Pb.VoteMessageRequest) error {
	data := &postgresModel.Message{
		ChatType:   constant.ChatGroupMode,
		MsgType:    constant.ChatMsgTypeVote,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	m.loadSequence(ctx, data)

	options := make(map[int]string)
	for i, value := range req.Options {
		options[i+1] = value
	}

	num := m.GroupChatMemberRepo.CountMemberTotal(ctx, int(req.Receiver.ReceiverId))
	err := m.Source.Postgres().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}
		return tx.Create(&postgresModel.MessageVote{
			MessageId:    data.Id,
			UserId:       uid,
			Title:        req.Title,
			AnswerMode:   int(req.Mode),
			AnswerOption: jsonutil.Encode(options),
			AnswerNum:    int(num),
			IsAnonymous:  int(req.Anonymous),
		}).Error
	})
	if err == nil {
		data.Content = m.Locale.Localize("vote")
		m.afterHandle(ctx, data)
	}

	return err
}

func (m *MessageUseCase) SendForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) error {
	if err := m.MessageForward.Verify(ctx, uid, req); err != nil {
		return err
	}

	var (
		err   error
		items []*logic.ForwardRecord
	)
	if req.Mode == 1 {
		items, err = m.MessageForward.MultiSplitForward(ctx, uid, req)
	} else {
		items, err = m.MessageForward.MultiMergeForward(ctx, uid, req)
	}
	if err != nil {
		return err
	}

	for _, record := range items {
		if record.ChatType == constant.ChatPrivateMode {
			m.UnreadCache.Incr(ctx, constant.ChatPrivateMode, uid, record.ReceiverId)
		} else if record.ChatType == constant.ChatGroupMode {
			pipe := m.Source.Redis().Pipeline()
			for _, uid := range m.GroupChatMemberRepo.GetMemberIds(ctx, record.ReceiverId) {
				m.UnreadCache.PipeIncr(ctx, pipe, constant.ChatGroupMode, record.ReceiverId, uid)
			}
			_, _ = pipe.Exec(ctx)
		}

		_ = m.MessageCache.Set(ctx, record.ChatType, uid, record.ReceiverId, &redisModel.LastCacheMessage{
			Content:  m.Locale.Localize("forwarded_message"),
			Datetime: timeutil.DateTime(),
		})
	}
	_, _ = m.Source.Redis().Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, item := range items {
			data := jsonutil.Encode(map[string]any{
				"event": constant.SubEventImMessage,
				"data": jsonutil.Encode(map[string]any{
					"sender_id":   uid,
					"receiver_id": item.ReceiverId,
					"chat_type":   item.ChatType,
					"message_id":  item.MessageId,
				}),
			})
			pipe.Publish(ctx, constant.ImTopicChat, data)
		}

		return nil
	})

	return nil
}

func (m *MessageUseCase) SendMixedMessage(ctx context.Context, uid int, req *v1Pb.MixedMessageRequest) error {
	items := make([]*entity.MessageExtraMixedItem, 0)
	for _, item := range req.Items {
		items = append(items, &entity.MessageExtraMixedItem{
			Type:    int(item.Type),
			Content: item.Content,
		})
	}

	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeMixed,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(entity.MessageExtraMixed{
			Items: items,
		}),
		ReplyTo: int(req.ReplyToMsgId),
	})

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendSysOther(ctx context.Context, data *postgresModel.Message) error {
	message, err := m.save(ctx, data)

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendLogin(ctx context.Context, uid int, req *entity.SendLogin) error {
	bot, err := m.BotRepo.GetLoginBot(ctx)
	if err != nil {
		return err
	}

	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   constant.ChatPrivateMode,
		MsgType:    constant.ChatMsgTypeLogin,
		UserId:     bot.UserId,
		ReceiverId: uid,
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageLogin{
		MessageId: message.Id,
		IpAddress: req.Ip,
		UserAgent: req.Agent,
		Address:   req.Address,
		UserId:    uid,
		CreatedAt: time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendLogin.MessageLogin: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendSticker(ctx context.Context, uid int, req *v1Pb.StickerMessageRequest) error {
	var sticker postgresModel.StickerItem
	if err := m.Source.Postgres().First(&sticker, "id = ? AND user_id = ?", req.StickerId, uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(m.Locale.Localize("emoji_info_not_found"))
		}
		return err
	}

	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeImage,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageMedia{
		MessageId: message.Id,
		Url:       sticker.Url,
		CreatedAt: time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendImage.MessageMedia: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendCode(ctx context.Context, uid int, req *v1Pb.CodeMessageRequest) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeCode,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageCode{
		MessageId: message.Id,
		Lang:      req.Lang,
		Code:      req.Code,
		CreatedAt: time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendCode.MessageCode: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) SendLocation(ctx context.Context, uid int, req *v1Pb.LocationMessageRequest) error {
	message, err := m.save(ctx, &postgresModel.Message{
		ChatType:   int(req.Receiver.ChatType),
		MsgType:    constant.ChatMsgTypeLocation,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	})

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageLocation{
		MessageId:   message.Id,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - SendCode.MessageCode: %v", err)
	}

	m.afterHandle(ctx, message)

	return err
}

func (m *MessageUseCase) Vote(ctx context.Context, uid int, msgId int64, optionsValue string) (*postgresRepo.VoteStatistics, error) {
	db := m.Source.Postgres().WithContext(ctx)
	fields := []string{
		"messages.receiver_id",
		"messages.chat_type",
		"messages.msg_type",
		"vote.id as vote_id",
		"vote.id as message_id",
		"vote.answer_mode",
		"vote.answer_option",
		"vote.answer_num",
		"vote.status as vote_status",
	}
	query := db.Select(fields).
		Table("messages").
		Joins("LEFT JOIN message_votes as vote ON vote.message_id = messages.id").
		Where("messages.id = ?", msgId)

	var vote entity.QueryVoteModel
	if err := query.Take(&vote).Error; err != nil {
		return nil, err
	}

	if vote.MsgType != constant.ChatMsgTypeVote {
		return nil, fmt.Errorf(m.Locale.Localize("record_not_related_to_voting_info"), vote.MsgType)
	}

	if vote.ChatType == constant.ChatGroupMode {
		var count int64
		db.Table("group_chat_members").Where("group_id = ? AND user_id = ? AND is_quit = 0", vote.ReceiverId, uid).Count(&count)
		if count == 0 {
			return nil, errors.New(m.Locale.Localize("no_permission_to_vote"))
		}
	}

	var count int64
	db.Table("message_vote_answers").Where("vote_id = ? AND user_id = ?", vote.VoteId, uid).Count(&count)
	if count > 0 {
		return nil, fmt.Errorf(m.Locale.Localize("duplicate_vote"), vote.VoteId)
	}

	options := strings.Split(optionsValue, ",")
	sort.Strings(options)
	var answerOptions map[string]any
	if err := jsonutil.Decode(vote.AnswerOption, &answerOptions); err != nil {
		return nil, err
	}

	for _, option := range options {
		if _, ok := answerOptions[option]; !ok {
			return nil, fmt.Errorf(m.Locale.Localize("invalid_vote_option"), option)
		}
	}

	if vote.AnswerMode == constant.VoteAnswerModeSingleChoice {
		options = options[:1]
	}
	answers := make([]*postgresModel.MessageVoteAnswer, 0, len(options))
	for _, option := range options {
		answers = append(answers, &postgresModel.MessageVoteAnswer{
			VoteId: vote.VoteId,
			UserId: uid,
			Option: option,
		})
	}

	err := m.Source.Postgres().Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("message_votes").
			Where("id = ?", vote.VoteId).
			Updates(map[string]any{
				"answered_num": gorm.Expr("answered_num + 1"),
				"status":       gorm.Expr("CASE WHEN answered_num >= answer_num THEN 1 ELSE 0 END"),
			}).Error; err != nil {
			return err
		}
		return tx.Create(answers).Error
	})
	if err != nil {
		return nil, err
	}

	_, _ = m.MessageVoteRepo.SetVoteAnswerUser(ctx, vote.VoteId)
	_, _ = m.MessageVoteRepo.SetVoteStatistics(ctx, vote.VoteId)
	info, _ := m.MessageVoteRepo.GetVoteStatistics(ctx, vote.VoteId)

	return info, nil
}

func (m *MessageUseCase) Revoke(ctx context.Context, uid int, msgId int64) error {
	var record postgresModel.Message
	if err := m.Source.Postgres().First(&record, "id = ?", msgId).Error; err != nil {
		return err
	}

	if record.IsRevoke == 1 {
		return nil
	}

	if record.UserId != uid {
		return errors.New(m.Locale.Localize("no_permission_to_revoke_message"))
	}

	if time.Now().Unix() > record.CreatedAt.Add(3*time.Minute).Unix() {
		return errors.New(m.Locale.Localize("revoke_time_exceeded"))
	}

	if err := m.Source.Postgres().Model(&postgresModel.Message{Id: record.Id}).Update("is_revoke", 1).Error; err != nil {
		return err
	}

	_ = m.MessageCache.Set(ctx, record.ChatType, record.UserId, record.ReceiverId, &redisModel.LastCacheMessage{
		Content:  m.Locale.Localize("message_deleted"),
		Datetime: timeutil.DateTime(),
	})

	body := map[string]any{
		"event": constant.SubEventImMessageRevoke,
		"data": jsonutil.Encode(map[string]any{
			"id": record.Id,
		}),
	}

	m.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(body))

	return nil
}

func (m *MessageUseCase) save(ctx context.Context, data *postgresModel.Message) (*postgresModel.Message, error) {
	m.loadSequence(ctx, data)

	if err := m.Source.Postgres().WithContext(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	m.loadReply(ctx, data)

	switch data.MsgType {
	case constant.ChatMsgTypeText:
		data.Content = strutil.MtSubstr(strutil.ReplaceImgAll(data.Content), 0, 300)
	default:
		data.Content = entity.GetChatMsgTypeMapping(m.Locale, data.MsgType)
	}

	return data, nil
}

func (m *MessageUseCase) loadReply(ctx context.Context, data *postgresModel.Message) {
	if data.ReplyTo == 0 {
		return
	}

	var message postgresModel.Message
	if err := m.Source.Postgres().
		Table("messages").
		Find(&message, "id = ?", data.ReplyTo).Error; err != nil {
		log.Printf("Error - IMessageUseCase - loadReply.Message: %v", err)
		return
	}

	if err := m.Source.Postgres().WithContext(ctx).Create(&postgresModel.MessageForwarded{
		OriginalMessageId: message.Id,
		NewMessageId:      data.Id,
		UserId:            data.UserId,
		CreatedAt:         time.Now(),
	}).Error; err != nil {
		log.Printf("Error - IMessageUseCase - loadReply.MessageForwarded: %v", err)
	}
}

func (m *MessageUseCase) loadSequence(ctx context.Context, data *postgresModel.Message) {
	if data.ChatType == constant.ChatGroupMode {
		data.Sequence = m.Sequence.Get(ctx, 0, data.ReceiverId)
	} else {
		data.Sequence = m.Sequence.Get(ctx, data.UserId, data.ReceiverId)
	}
}

func (m *MessageUseCase) writeMessageToQueue(uid int, userIds []int, message *postgresModel.Message) {
	userIds = append(userIds, uid)
	content := jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessage,
		"data": jsonutil.Encode(entity.ConsumeMessage{
			UserIds: userIds,
			Message: entity.Message{
				Id:         int64(message.Id),
				ChatType:   message.ChatType,
				MsgType:    message.MsgType,
				ReceiverId: message.ReceiverId,
				UserId:     message.UserId,
				Content:    message.Content,
				IsRead:     message.IsRead == 1,
				CreatedAt:  message.CreatedAt.String(),
			},
		}),
	})

	if err := m.Nats.Publish(constant.ImTopicChat, []byte(content)); err != nil {
		log.Println(err)
	}
}

func (m *MessageUseCase) writeMessageToPushQueue(userIds []int, message *postgresModel.Message) {
	pushPayload := &entity.PushPayload{
		UserIds: userIds,
		Message: message.Content,
	}

	pushData, err := json.Marshal(pushPayload)
	if err != nil {
		log.Fatal(m.Locale.Localize("json_serialization_error"), err)
	}

	if err := m.Nats.Publish(constant.QueuePush, pushData); err != nil {
		log.Println(err)
	}
}

func (m *MessageUseCase) afterHandle(ctx context.Context, record *postgresModel.Message) {
	userIds := make([]int, 0)
	if record.ChatType == constant.ChatPrivateMode {
		m.UnreadCache.Incr(ctx, constant.ChatPrivateMode, record.UserId, record.ReceiverId)
		userIds = append(userIds, record.ReceiverId)
		if record.MsgType == constant.ChatMsgSysText {
			m.UnreadCache.Incr(ctx, 1, record.ReceiverId, record.UserId)
		}
	} else if record.ChatType == constant.ChatGroupMode {
		pipe := m.Source.Redis().Pipeline()
		for _, uid := range m.GroupChatMemberRepo.GetMemberIds(ctx, record.ReceiverId) {
			if uid != record.UserId {
				m.UnreadCache.PipeIncr(ctx, pipe, constant.ChatGroupMode, record.ReceiverId, uid)
				userIds = append(userIds, uid)
			}
		}

		_, _ = pipe.Exec(ctx)
	}

	_ = m.MessageCache.Set(ctx, record.ChatType, record.UserId, record.ReceiverId, &redisModel.LastCacheMessage{
		Content:  record.Content,
		Datetime: timeutil.DateTime(),
	})

	content := jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessage,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"chat_type":   record.ChatType,
			"message_id":  record.Id,
		}),
	})

	if record.ChatType == constant.ChatPrivateMode {
		sids := m.ServerCache.All(ctx, 1)
		if len(sids) > 3 {
			pipe := m.Source.Redis().Pipeline()
			for _, sid := range sids {
				for _, uid := range []int{record.UserId, record.ReceiverId} {
					if !m.ClientCache.IsCurrentServerOnline(ctx, sid, constant.ImChannelChat, strconv.Itoa(uid)) {
						continue
					}

					pipe.Publish(ctx, fmt.Sprintf(constant.ImTopicChatPrivate, sid), content)
				}
			}

			if _, err := pipe.Exec(ctx); err == nil {
				return
			}
		}
	}

	if err := m.Source.Redis().Publish(ctx, constant.ImTopicChat, content).Err(); err != nil {
		log.Fatalf(m.Locale.Localize("notification_sending_error"), err.Error())
	}

	m.writeMessageToQueue(record.UserId, userIds, record)
	if record.ChatType == constant.ChatPrivateMode {
		m.writeMessageToPushQueue(userIds, record)
	}
}

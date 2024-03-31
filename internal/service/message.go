package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"html"
	"sort"
	"strconv"
	"strings"
	"time"
	"voo.su/api/pb/v1"
	"voo.su/internal/entity"
	"voo.su/internal/logic"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/filesystem"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/logger"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

var _ MessageSendService = (*MessageService)(nil)

type MessageSendService interface {
	SendSystemText(ctx context.Context, uid int, req *api_v1.TextMessageRequest) error
	SendText(ctx context.Context, uid int, req *api_v1.TextMessageRequest) error
	SendImage(ctx context.Context, uid int, req *api_v1.ImageMessageRequest) error
	SendVoice(ctx context.Context, uid int, req *api_v1.VoiceMessageRequest) error
	SendVideo(ctx context.Context, uid int, req *api_v1.VideoMessageRequest) error
	SendFile(ctx context.Context, uid int, req *api_v1.FileMessageRequest) error
	SendCode(ctx context.Context, uid int, req *api_v1.CodeMessageRequest) error
	SendVote(ctx context.Context, uid int, req *api_v1.VoteMessageRequest) error
	SendForward(ctx context.Context, uid int, req *api_v1.ForwardMessageRequest) error
	SendLocation(ctx context.Context, uid int, req *api_v1.LocationMessageRequest) error
	SendBusinessCard(ctx context.Context, uid int, req *api_v1.CardMessageRequest) error
	SendSysOther(ctx context.Context, data *model.Message) error
	SendMixedMessage(ctx context.Context, uid int, req *api_v1.MixedMessageRequest) error
	Revoke(ctx context.Context, uid int, msgId string) error
	Vote(ctx context.Context, uid int, msgId int, optionsValue string) (*repo.VoteStatistics, error)
	SendLogin(ctx context.Context, uid int, req *api_v1.LoginMessageRequest) error
}

type MessageService struct {
	*repo.Source
	MessageForwardLogic *logic.MessageForwardLogic
	GroupChatMemberRepo *repo.GroupChatMember
	SplitRepo           *repo.Split
	MessageVoteRepo     *repo.MessageVote
	Filesystem          *filesystem.Filesystem
	UnreadStorage       *cache.UnreadStorage
	MessageStorage      *cache.MessageStorage
	ServerStorage       *cache.ServerStorage
	ClientStorage       *cache.ClientStorage
	Sequence            *repo.Sequence
	DialogVoteCache     *cache.Vote
	MessageRepo         *repo.Message
	BotRepo             *repo.Bot
}

type DialogRecordsItem struct {
	Id         int    `json:"id"`
	Sequence   int    `json:"sequence"`
	MsgId      string `json:"msg_id"`
	DialogType int    `json:"dialog_type"`
	MsgType    int    `json:"msg_type"`
	UserId     int    `json:"user_id"`
	ReceiverId int    `json:"receiver_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Avatar     string `json:"avatar"`
	IsRevoke   int    `json:"is_revoke"`
	IsMark     int    `json:"is_mark"`
	IsRead     int    `json:"is_read"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	Extra      any    `json:"extra"`
}

type QueryDialogRecordsOpt struct {
	DialogType int
	UserId     int
	ReceiverId int
	MsgType    []int
	RecordId   int
	Limit      int
}

type QueryDialogRecordsItem struct {
	Id         int       `json:"id"`
	MsgId      string    `json:"msg_id"`
	Sequence   int64     `json:"sequence"`
	DialogType int       `json:"dialog_type"`
	MsgType    int       `json:"msg_type"`
	UserId     int       `json:"user_id"`
	ReceiverId int       `json:"receiver_id"`
	IsRevoke   int       `json:"is_revoke"`
	IsMark     int       `json:"is_mark"`
	IsRead     int       `json:"is_read"`
	QuoteId    int       `json:"quote_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Avatar     string    `json:"avatar"`
	Extra      string    `json:"extra"`
}

func (m *MessageService) GetDialogRecords(ctx context.Context, opt *QueryDialogRecordsOpt) ([]*DialogRecordsItem, error) {
	var (
		items  = make([]*QueryDialogRecordsItem, 0, opt.Limit)
		fields = []string{
			"messages.id",
			"messages.sequence",
			"messages.dialog_type",
			"messages.msg_type",
			"messages.msg_id",
			"messages.user_id",
			"messages.receiver_id",
			"messages.is_revoke",
			"messages.is_read",
			"messages.content",
			"messages.extra",
			"messages.created_at",
			"users.username",
			"users.name as name",
			"users.surname as surname",
			"users.avatar as avatar",
		}
	)
	query := m.Source.Db().WithContext(ctx).Table("messages")
	query.Joins("LEFT JOIN users on messages.user_id = users.id")
	query.Joins("LEFT JOIN message_delete on messages.id = message_delete.record_id and message_delete.user_id = ?", opt.UserId)
	if opt.RecordId > 0 {
		query.Where("messages.sequence < ?", opt.RecordId)
	}

	if opt.DialogType == entity.ChatPrivateMode {
		subQuery := m.Source.Db().Where("messages.user_id = ? and messages.receiver_id = ?", opt.UserId, opt.ReceiverId)
		subQuery.Or("messages.user_id = ? and messages.receiver_id = ?", opt.ReceiverId, opt.UserId)
		query.Where(subQuery)
	} else {
		query.Where("messages.receiver_id = ?", opt.ReceiverId)
	}

	if opt.MsgType != nil && len(opt.MsgType) > 0 {
		query.Where("messages.msg_type in ?", opt.MsgType)
	}

	query.Where("messages.dialog_type = ?", opt.DialogType)
	query.Where("COALESCE(message_delete.id,0) = 0")
	query.Select(fields).Order("messages.sequence desc").Limit(opt.Limit)
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return make([]*DialogRecordsItem, 0), nil
	}

	return m.HandleDialogRecords(ctx, items)
}

func (m *MessageService) GetDialogRecord(ctx context.Context, recordId int64) (*DialogRecordsItem, error) {
	var (
		err    error
		item   *QueryDialogRecordsItem
		fields = []string{
			"messages.id",
			"messages.msg_id",
			"messages.sequence",
			"messages.dialog_type",
			"messages.msg_type",
			"messages.user_id",
			"messages.receiver_id",
			"messages.is_revoke",
			"messages.content",
			"messages.extra",
			"messages.created_at",
			"users.username",
			"users.avatar as avatar",
		}
	)
	query := m.Source.Db().Table("messages")
	query.Joins("LEFT JOIN users on messages.user_id = users.id")
	query.Where("messages.id = ?", recordId)
	if err = query.Select(fields).Take(&item).Error; err != nil {
		return nil, err
	}

	list, err := m.HandleDialogRecords(ctx, []*QueryDialogRecordsItem{item})
	if err != nil {
		return nil, err
	}

	return list[0], nil
}

func (m *MessageService) GetForwardRecords(ctx context.Context, uid int, recordId int64) ([]*DialogRecordsItem, error) {
	record, err := m.MessageRepo.FindById(ctx, int(recordId))
	if err != nil {
		return nil, err
	}

	//if record.DialogType == entity.ChatPrivateMode {
	//	if record.UserId != uid && record.ReceiverId != uid {
	//		return nil, entity.ErrPermissionDenied
	//	}
	//} else if record.DialogType == entity.ChatGroupMode {
	//	if !s.GroupMemberRepo.IsMember(ctx, record.ReceiverId, uid, true) {
	//		return nil, entity.ErrPermissionDenied
	//	}
	//} else {
	//	return nil, entity.ErrPermissionDenied
	//}

	var extra model.DialogRecordExtraForward
	if err := jsonutil.Decode(record.Extra, &extra); err != nil {
		return nil, err
	}
	var (
		items  = make([]*QueryDialogRecordsItem, 0)
		fields = []string{
			"messages.id",
			"messages.msg_id",
			"messages.sequence",
			"messages.dialog_type",
			"messages.msg_type",
			"messages.user_id",
			"messages.receiver_id",
			"messages.is_revoke",
			"messages.content",
			"messages.extra",
			"messages.created_at",
			"users.username",
			"users.avatar as avatar",
		}
	)
	query := m.Source.Db().Table("messages")
	query.Select(fields)
	query.Joins("LEFT JOIN users on messages.user_id = users.id")
	query.Where("messages.id in ?", extra.MsgIds)
	query.Order("messages.sequence asc")
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return m.HandleDialogRecords(ctx, items)
}

func (m *MessageService) HandleDialogRecords(ctx context.Context, items []*QueryDialogRecordsItem) ([]*DialogRecordsItem, error) {
	var (
		votes     []int
		voteItems []*model.MessageVote
	)
	for _, item := range items {
		switch item.MsgType {
		case entity.ChatMsgTypeVote:
			votes = append(votes, item.Id)
		}
	}

	hashVotes := make(map[int]*model.MessageVote)
	if len(votes) > 0 {
		m.Source.Db().Model(&model.MessageVote{}).Where("record_id in ?", votes).Scan(&voteItems)
		for i := range voteItems {
			hashVotes[voteItems[i].RecordId] = voteItems[i]
		}
	}

	newItems := make([]*DialogRecordsItem, 0, len(items))
	for _, item := range items {
		data := &DialogRecordsItem{
			Id:         item.Id,
			MsgId:      item.MsgId,
			Sequence:   int(item.Sequence),
			DialogType: item.DialogType,
			MsgType:    item.MsgType,
			UserId:     item.UserId,
			ReceiverId: item.ReceiverId,
			Username:   item.Username,
			Name:       item.Name,
			Surname:    item.Surname,
			Avatar:     item.Avatar,
			IsRevoke:   item.IsRevoke,
			IsMark:     item.IsMark,
			IsRead:     item.IsRead,
			Content:    item.Content,
			CreatedAt:  timeutil.FormatDatetime(item.CreatedAt),
			Extra:      make(map[string]any),
		}
		_ = jsonutil.Decode(item.Extra, &data.Extra)
		switch item.MsgType {
		//case entity.ChatMsgSysGroupCreate:
		//    fmt.Println(item.Extra.OwnerId)
		case entity.ChatMsgTypeVote:
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

				data.Extra = map[string]any{
					"detail": map[string]any{
						"id":            value.Id,
						"record_id":     value.RecordId,
						"title":         value.Title,
						"answer_mode":   value.AnswerMode,
						"status":        value.Status,
						"answer_option": opts,
						"answer_num":    value.AnswerNum,
						"answered_num":  value.AnsweredNum,
					},
					"statistics": statistics,
					"vote_users": users,
				}
			}
		}
		newItems = append(newItems, data)
	}

	return newItems, nil
}

func (m *MessageService) SendSystemText(ctx context.Context, uid int, req *api_v1.TextMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgSysText,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    html.EscapeString(req.Content),
	}
	return m.save(ctx, data)
}

func (m *MessageService) SendText(ctx context.Context, uid int, req *api_v1.TextMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeText,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    strutil.EscapeHtml(req.Content),
		QuoteId:    req.QuoteId,
	}

	return m.save(ctx, data)
}

func (m *MessageService) SendImage(ctx context.Context, uid int, req *api_v1.ImageMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeImage,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&model.DialogRecordExtraImage{
			Url:    req.Url,
			Width:  int(req.Width),
			Height: int(req.Height),
		}),
		QuoteId: req.QuoteId,
	}

	return m.save(ctx, data)
}

func (m *MessageService) SendVoice(ctx context.Context, uid int, req *api_v1.VoiceMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeAudio,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&model.DialogRecordExtraAudio{
			Suffix:   strutil.FileSuffix(req.Url),
			Size:     int(req.Size),
			Url:      req.Url,
			Duration: 0,
		}),
	}

	return m.save(ctx, data)
}

func (m *MessageService) SendVideo(ctx context.Context, uid int, req *api_v1.VideoMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeVideo,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&model.DialogRecordExtraVideo{
			Cover:    req.Cover,
			Suffix:   strutil.FileSuffix(req.Url),
			Size:     int(req.Size),
			Url:      req.Url,
			Duration: int(req.Duration),
		}),
	}
	return m.save(ctx, data)
}

func (m *MessageService) SendFile(ctx context.Context, uid int, req *api_v1.FileMessageRequest) error {
	file, err := m.SplitRepo.GetFile(ctx, uid, req.UploadId)
	if err != nil {
		return err
	}
	publicUrl := ""
	filePath := fmt.Sprintf("private-dialog/%s/%s.%s", timeutil.DateNumber(), encrypt.Md5(strutil.Random(16)), file.FileExt)
	if entity.GetMediaType(file.FileExt) <= 3 {
		filePath = fmt.Sprintf("file/%s/%s.%s", timeutil.DateNumber(), encrypt.Md5(strutil.Random(16)), file.FileExt)
		publicUrl = m.Filesystem.Default.PublicUrl(filePath)
	}

	if err := m.Filesystem.Default.Copy(file.Path, filePath); err != nil {
		return err
	}
	data := &model.Message{
		MsgId:      encrypt.Md5(req.UploadId),
		DialogType: int(req.Receiver.DialogType),
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}

	switch entity.GetMediaType(file.FileExt) {
	case entity.MediaFileAudio:
		data.MsgType = entity.ChatMsgTypeAudio
		data.Extra = jsonutil.Encode(&model.DialogRecordExtraAudio{
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case entity.MediaFileVideo:
		data.MsgType = entity.ChatMsgTypeVideo
		data.Extra = jsonutil.Encode(&model.DialogRecordExtraVideo{
			Cover:    "",
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case entity.MediaFileOther:
		data.MsgType = entity.ChatMsgTypeFile
		data.Extra = jsonutil.Encode(&model.DialogRecordExtraFile{
			Drive:  file.Drive,
			Name:   file.OriginalName,
			Suffix: file.FileExt,
			Size:   int(file.FileSize),
			Path:   filePath,
		})
	}

	return m.save(ctx, data)
}

func (m *MessageService) SendCode(ctx context.Context, uid int, req *api_v1.CodeMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeCode,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&model.DialogRecordExtraCode{
			Lang: req.Lang,
			Code: req.Code,
		}),
	}
	return m.save(ctx, data)
}

func (m *MessageService) SendVote(ctx context.Context, uid int, req *api_v1.VoteMessageRequest) error {
	data := &model.Message{
		MsgId:      strutil.NewMsgId(),
		DialogType: entity.ChatGroupMode,
		MsgType:    entity.ChatMsgTypeVote,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
	}
	m.loadSequence(ctx, data)
	options := make(map[int]string)
	for i, value := range req.Options {
		options[i+1] = value
	}
	num := m.GroupChatMemberRepo.CountMemberTotal(ctx, int(req.Receiver.ReceiverId))
	err := m.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}
		return tx.Create(&model.MessageVote{
			RecordId:     data.Id,
			UserId:       uid,
			Title:        req.Title,
			AnswerMode:   int(req.Mode),
			AnswerOption: jsonutil.Encode(options),
			AnswerNum:    int(num),
			IsAnonymous:  int(req.Anonymous),
		}).Error
	})
	if err == nil {
		m.afterHandle(ctx, data, map[string]string{"text": "Опрос"})
	}

	return err
}

func (m *MessageService) SendForward(ctx context.Context, uid int, req *api_v1.ForwardMessageRequest) error {
	if err := m.MessageForwardLogic.Verify(ctx, uid, req); err != nil {
		return err
	}

	var (
		err   error
		items []*logic.ForwardRecord
	)
	if req.Mode == 1 {
		items, err = m.MessageForwardLogic.MultiSplitForward(ctx, uid, req)
	} else {
		items, err = m.MessageForwardLogic.MultiMergeForward(ctx, uid, req)
	}
	if err != nil {
		return err
	}

	for _, record := range items {
		if record.DialogType == entity.ChatPrivateMode {
			m.UnreadStorage.Incr(ctx, entity.ChatPrivateMode, uid, record.ReceiverId)
		} else if record.DialogType == entity.ChatGroupMode {
			pipe := m.Source.Redis().Pipeline()
			for _, uid := range m.GroupChatMemberRepo.GetMemberIds(ctx, record.ReceiverId) {
				m.UnreadStorage.PipeIncr(ctx, pipe, entity.ChatGroupMode, record.ReceiverId, uid)
			}
			_, _ = pipe.Exec(ctx)
		}

		_ = m.MessageStorage.Set(ctx, record.DialogType, uid, record.ReceiverId, &cache.LastCacheMessage{
			Content:  "Пересланное сообщение",
			Datetime: timeutil.DateTime(),
		})
	}
	_, _ = m.Source.Redis().Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, item := range items {
			data := jsonutil.Encode(map[string]any{
				"event": entity.SubEventImMessage,
				"data": jsonutil.Encode(map[string]any{
					"sender_id":   uid,
					"receiver_id": item.ReceiverId,
					"dialog_type": item.DialogType,
					"record_id":   item.RecordId,
				}),
			})
			pipe.Publish(ctx, entity.ImTopicChat, data)
		}

		return nil
	})

	return nil
}

func (m *MessageService) SendLocation(ctx context.Context, uid int, req *api_v1.LocationMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeLocation,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&model.DialogRecordExtraLocation{
			Longitude:   req.Longitude,
			Latitude:    req.Latitude,
			Description: req.Description,
		}),
	}
	return m.save(ctx, data)
}

func (m *MessageService) SendBusinessCard(ctx context.Context, uid int, req *api_v1.CardMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeCard,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&model.DialogRecordExtraCard{
			UserId: int(req.UserId),
		}),
	}

	return m.save(ctx, data)
}

func (m *MessageService) SendMixedMessage(ctx context.Context, uid int, req *api_v1.MixedMessageRequest) error {

	items := make([]*model.DialogRecordExtraMixedItem, 0)

	for _, item := range req.Items {
		items = append(items, &model.DialogRecordExtraMixedItem{
			Type:    int(item.Type),
			Content: item.Content,
		})
	}

	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    entity.ChatMsgTypeMixed,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra:      jsonutil.Encode(model.DialogRecordExtraMixed{Items: items}),
		QuoteId:    req.QuoteId,
	}

	return m.save(ctx, data)
}

func (m *MessageService) SendSysOther(ctx context.Context, data *model.Message) error {
	return m.save(ctx, data)
}

func (m *MessageService) Revoke(ctx context.Context, uid int, msgId string) error {
	var record model.Message
	if err := m.Source.Db().First(&record, "msg_id = ?", msgId).Error; err != nil {
		return err
	}
	if record.IsRevoke == 1 {
		return nil
	}

	if record.UserId != uid {
		return errors.New("нет прав на отзыв данного сообщения")
	}

	if time.Now().Unix() > record.CreatedAt.Add(3*time.Minute).Unix() {
		return errors.New("превышено допустимое время для отзыва. Отзыв невозможен")
	}

	if err := m.Source.Db().Model(&model.Message{Id: record.Id}).Update("is_revoke", 1).Error; err != nil {
		return err
	}

	_ = m.MessageStorage.Set(ctx, record.DialogType, record.UserId, record.ReceiverId, &cache.LastCacheMessage{
		Content:  "Данное сообщение удалено",
		Datetime: timeutil.DateTime(),
	})

	body := map[string]any{
		"event": entity.SubEventImMessageRevoke,
		"data": jsonutil.Encode(map[string]any{
			"msg_id": record.MsgId,
		}),
	}

	m.Source.Redis().Publish(ctx, entity.ImTopicChat, jsonutil.Encode(body))

	return nil
}

func (m *MessageService) Vote(ctx context.Context, uid int, msgId int, optionsValue string) (*repo.VoteStatistics, error) {
	db := m.Source.Db().WithContext(ctx)
	query := db.Table("messages")
	query.Select([]string{
		"messages.receiver_id", "messages.dialog_type", "messages.msg_type",
		"vote.id as vote_id", "vote.id as record_id", "vote.answer_mode", "vote.answer_option",
		"vote.answer_num", "vote.status as vote_status",
	})
	query.Joins("LEFT JOIN message_votes as vote on vote.record_id = messages.id")
	query.Where("messages.id = ?", msgId)

	var vote model.QueryVoteModel
	if err := query.Take(&vote).Error; err != nil {
		return nil, err
	}

	if vote.MsgType != entity.ChatMsgTypeVote {
		return nil, fmt.Errorf("текущая запись не относится к информации о голосовании %d", vote.MsgType)
	}

	if vote.DialogType == entity.ChatGroupMode {
		var count int64
		db.Table("group_chat_members").Where("group_id = ? and user_id = ? and is_quit = 0", vote.ReceiverId, uid).Count(&count)
		if count == 0 {
			return nil, errors.New("нет прав на голосование")
		}
	}

	var count int64
	db.Table("message_vote_answers").Where("vote_id = ? and user_id = ?", vote.VoteId, uid).Count(&count)
	if count > 0 {
		return nil, fmt.Errorf("повторное голосование %d", vote.VoteId)
	}

	options := strings.Split(optionsValue, ",")
	sort.Strings(options)
	var answerOptions map[string]any
	if err := jsonutil.Decode(vote.AnswerOption, &answerOptions); err != nil {
		return nil, err
	}

	for _, option := range options {
		if _, ok := answerOptions[option]; !ok {
			return nil, fmt.Errorf("недопустимый вариант голосования %s", option)
		}
	}

	if vote.AnswerMode == model.VoteAnswerModeSingleChoice {
		options = options[:1]
	}
	answers := make([]*model.MessageVoteAnswer, 0, len(options))
	for _, option := range options {
		answers = append(answers, &model.MessageVoteAnswer{
			VoteId: vote.VoteId,
			UserId: uid,
			Option: option,
		})
	}

	err := m.Source.Db().Transaction(func(tx *gorm.DB) error {
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

func (m *MessageService) save(ctx context.Context, data *model.Message) error {
	if data.MsgId == "" {
		data.MsgId = strutil.NewMsgId()
	}
	m.loadReply(ctx, data)
	m.loadSequence(ctx, data)
	if err := m.Source.Db().WithContext(ctx).Create(data).Error; err != nil {
		return err
	}
	option := make(map[string]string)
	switch data.MsgType {
	case entity.ChatMsgTypeText:
		option["text"] = strutil.MtSubstr(strutil.ReplaceImgAll(data.Content), 0, 300)
	default:
		if value, ok := entity.ChatMsgTypeMapping[data.MsgType]; ok {
			option["text"] = value
		} else {
			option["text"] = "Неизвестно"
		}
	}
	m.afterHandle(ctx, data, option)
	return nil
}

func (m *MessageService) loadReply(_ context.Context, data *model.Message) {
	if data.QuoteId == "" {
		return
	}
	if data.Extra == "" {
		data.Extra = "{}"
	}
	extra := make(map[string]any)
	err := jsonutil.Decode(data.Extra, &extra)
	if err != nil {
		logger.Errorf("MessageService Json Decode err: %s", err.Error())
		return
	}
	var record model.Message
	err = m.Source.Db().Table("messages").Find(&record, "msg_id = ?", data.QuoteId).Error
	if err != nil {
		return
	}
	var user model.User
	err = m.Source.Db().Table("users").Select("username").Find(&user, "id = ?", record.UserId).Error
	if err != nil {
		return
	}
	reply := model.Reply{
		UserId:   record.UserId,
		Username: user.Username,
		MsgType:  1,
		Content:  record.Content,
		MsgId:    record.MsgId,
	}
	if record.MsgType != entity.ChatMsgTypeText {
		reply.Content = "Неизвестно"
		if value, ok := entity.ChatMsgTypeMapping[record.MsgType]; ok {
			reply.Content = value
		}
	}
	extra["reply"] = reply
	data.Extra = jsonutil.Encode(extra)
}

func (m *MessageService) loadSequence(ctx context.Context, data *model.Message) {
	if data.DialogType == entity.ChatGroupMode {
		data.Sequence = m.Sequence.Get(ctx, 0, data.ReceiverId)
	} else {
		data.Sequence = m.Sequence.Get(ctx, data.UserId, data.ReceiverId)
	}
}

func (m *MessageService) afterHandle(ctx context.Context, record *model.Message, opt map[string]string) {
	if record.DialogType == entity.ChatPrivateMode {
		m.UnreadStorage.Incr(ctx, entity.ChatPrivateMode, record.UserId, record.ReceiverId)
		if record.MsgType == entity.ChatMsgSysText {
			m.UnreadStorage.Incr(ctx, 1, record.ReceiverId, record.UserId)
		}
	} else if record.DialogType == entity.ChatGroupMode {
		pipe := m.Source.Redis().Pipeline()
		for _, uid := range m.GroupChatMemberRepo.GetMemberIds(ctx, record.ReceiverId) {
			if uid != record.UserId {
				m.UnreadStorage.PipeIncr(ctx, pipe, entity.ChatGroupMode, record.ReceiverId, uid)
			}
		}
		_, _ = pipe.Exec(ctx)
	}

	_ = m.MessageStorage.Set(ctx, record.DialogType, record.UserId, record.ReceiverId, &cache.LastCacheMessage{
		Content:  opt["text"],
		Datetime: timeutil.DateTime(),
	})

	content := jsonutil.Encode(map[string]any{
		"event": entity.SubEventImMessage,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"dialog_type": record.DialogType,
			"record_id":   record.Id,
		}),
	})

	if record.DialogType == entity.ChatPrivateMode {
		sids := m.ServerStorage.All(ctx, 1)
		if len(sids) > 3 {
			pipe := m.Source.Redis().Pipeline()
			for _, sid := range sids {
				for _, uid := range []int{record.UserId, record.ReceiverId} {
					if !m.ClientStorage.IsCurrentServerOnline(ctx, sid, entity.ImChannelChat, strconv.Itoa(uid)) {
						continue
					}
					pipe.Publish(ctx, fmt.Sprintf(entity.ImTopicChatPrivate, sid), content)
				}
			}
			if _, err := pipe.Exec(ctx); err == nil {
				return
			}
		}
	}
	if err := m.Source.Redis().Publish(ctx, entity.ImTopicChat, content).Err(); err != nil {
		logger.Errorf("Ошибка отправки уведомления %s", err.Error())
	}
}

func (m *MessageService) SendLogin(ctx context.Context, uid int, req *api_v1.LoginMessageRequest) error {
	bot, err := m.BotRepo.GetLoginBot(ctx)
	if err != nil {
		return err
	}

	data := &model.Message{
		DialogType: entity.ChatPrivateMode,
		MsgType:    entity.ChatMsgTypeLogin,
		UserId:     bot.UserId,
		ReceiverId: uid,
		Extra: jsonutil.Encode(&model.DialogRecordExtraLogin{
			IP:       req.Ip,
			Agent:    req.Agent,
			Datetime: timeutil.DateTime(),
			//Platform: req.Platform,
			//Address:  req.Address,
			//Reason:   req.Reason,
		}),
	}

	return m.save(ctx, data)
}

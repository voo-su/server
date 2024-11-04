package usecase

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
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/domain/logic"
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

var _ MessageSendUseCase = (*MessageUseCase)(nil)

type MessageSendUseCase interface {
	SendSystemText(ctx context.Context, uid int, req *v1Pb.TextMessageRequest) error
	SendText(ctx context.Context, uid int, req *SendText) error
	SendImage(ctx context.Context, uid int, req *v1Pb.ImageMessageRequest) error
	SendVoice(ctx context.Context, uid int, req *v1Pb.VoiceMessageRequest) error
	SendVideo(ctx context.Context, uid int, req *v1Pb.VideoMessageRequest) error
	SendFile(ctx context.Context, uid int, req *v1Pb.FileMessageRequest) error
	SendVote(ctx context.Context, uid int, req *v1Pb.VoteMessageRequest) error
	SendForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) error
	SendSysOther(ctx context.Context, data *model.Message) error
	SendMixedMessage(ctx context.Context, uid int, req *v1Pb.MixedMessageRequest) error
	Revoke(ctx context.Context, uid int, msgId string) error
	Vote(ctx context.Context, uid int, msgId int, optionsValue string) (*repo.VoteStatistics, error)
	SendLogin(ctx context.Context, uid int, req *SendLogin) error
	SendSticker(ctx context.Context, uid int, req *v1Pb.StickerMessageRequest) error
	SendCode(ctx context.Context, uid int, req *v1Pb.CodeMessageRequest) error
}

type MessageUseCase struct {
	*repo.Source
	MessageForwardLogic *logic.MessageForwardLogic
	Filesystem          *filesystem.Filesystem
	GroupChatMemberRepo *repo.GroupChatMember
	SplitRepo           *repo.Split
	MessageVoteRepo     *repo.MessageVote
	Sequence            *repo.Sequence
	MessageRepo         *repo.Message
	BotRepo             *repo.Bot
	UnreadStorage       *cache.UnreadStorage
	MessageStorage      *cache.MessageStorage
	ServerStorage       *cache.ServerStorage
	ClientStorage       *cache.ClientStorage
	DialogVoteCache     *cache.Vote
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

func (m *MessageUseCase) GetDialogRecords(ctx context.Context, opt *QueryDialogRecordsOpt) ([]*DialogRecordsItem, error) {
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

	if opt.DialogType == constant.ChatPrivateMode {
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

func (m *MessageUseCase) GetDialogRecord(ctx context.Context, recordId int64) (*DialogRecordsItem, error) {
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

func (m *MessageUseCase) GetForwardRecords(ctx context.Context, uid int, recordId int64) ([]*DialogRecordsItem, error) {
	record, err := m.MessageRepo.FindById(ctx, int(recordId))
	if err != nil {
		return nil, err
	}

	//if record.DialogType == domain.ChatPrivateMode {
	//	if record.UserId != uid && record.ReceiverId != uid {
	//		return nil, domain.ErrPermissionDenied
	//	}
	//} else if record.DialogType == domain.ChatGroupMode {
	//	if !s.GroupMemberRepo.IsMember(ctx, record.ReceiverId, uid, true) {
	//		return nil, domain.ErrPermissionDenied
	//	}
	//} else {
	//	return nil, domain.ErrPermissionDenied
	//}

	var extra entity.DialogRecordExtraForward
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

func (m *MessageUseCase) HandleDialogRecords(ctx context.Context, items []*QueryDialogRecordsItem) ([]*DialogRecordsItem, error) {
	var (
		votes     []int
		voteItems []*model.MessageVote
	)
	for _, item := range items {
		switch item.MsgType {
		case constant.ChatMsgTypeVote:
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
		//case constant.ChatMsgSysGroupCreate:
		//    fmt.Println(item.Extra.OwnerId)
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

func (m *MessageUseCase) SendSystemText(ctx context.Context, uid int, req *v1Pb.TextMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgSysText,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    html.EscapeString(req.Content),
	}
	return m.save(ctx, data)
}

type Receiver struct {
	DialogType int32
	ReceiverId int32
}

type SendText struct {
	Receiver Receiver
	Content  string
	QuoteId  string
}

func (m *MessageUseCase) SendText(ctx context.Context, uid int, req *SendText /*req1 *v1Pb.TextMessageRequest*/) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgTypeText,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Content:    strutil.EscapeHtml(req.Content),
		QuoteId:    req.QuoteId,
	}

	return m.save(ctx, data)
}

func (m *MessageUseCase) SendImage(ctx context.Context, uid int, req *v1Pb.ImageMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgTypeImage,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&entity.DialogRecordExtraImage{
			Url:    req.Url,
			Width:  int(req.Width),
			Height: int(req.Height),
		}),
		QuoteId: req.QuoteId,
	}

	return m.save(ctx, data)
}

func (m *MessageUseCase) SendVoice(ctx context.Context, uid int, req *v1Pb.VoiceMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgTypeAudio,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&entity.DialogRecordExtraAudio{
			Suffix:   strutil.FileSuffix(req.Url),
			Size:     int(req.Size),
			Url:      req.Url,
			Duration: 0,
		}),
	}

	return m.save(ctx, data)
}

func (m *MessageUseCase) SendVideo(ctx context.Context, uid int, req *v1Pb.VideoMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgTypeVideo,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&entity.DialogRecordExtraVideo{
			Cover:    req.Cover,
			Suffix:   strutil.FileSuffix(req.Url),
			Size:     int(req.Size),
			Url:      req.Url,
			Duration: int(req.Duration),
		}),
	}
	return m.save(ctx, data)
}

func (m *MessageUseCase) SendFile(ctx context.Context, uid int, req *v1Pb.FileMessageRequest) error {
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
	case constant.MediaFileAudio:
		data.MsgType = constant.ChatMsgTypeAudio
		data.Extra = jsonutil.Encode(&entity.DialogRecordExtraAudio{
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case constant.MediaFileVideo:
		data.MsgType = constant.ChatMsgTypeVideo
		data.Extra = jsonutil.Encode(&entity.DialogRecordExtraVideo{
			Cover:    "",
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case constant.MediaFileOther:
		data.MsgType = constant.ChatMsgTypeFile
		data.Extra = jsonutil.Encode(&entity.DialogRecordExtraFile{
			Drive:  file.Drive,
			Name:   file.OriginalName,
			Suffix: file.FileExt,
			Size:   int(file.FileSize),
			Path:   filePath,
		})
	}

	return m.save(ctx, data)
}

func (m *MessageUseCase) SendVote(ctx context.Context, uid int, req *v1Pb.VoteMessageRequest) error {
	data := &model.Message{
		MsgId:      strutil.NewMsgId(),
		DialogType: constant.ChatGroupMode,
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

func (m *MessageUseCase) SendForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) error {
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
		if record.DialogType == constant.ChatPrivateMode {
			m.UnreadStorage.Incr(ctx, constant.ChatPrivateMode, uid, record.ReceiverId)
		} else if record.DialogType == constant.ChatGroupMode {
			pipe := m.Source.Redis().Pipeline()
			for _, uid := range m.GroupChatMemberRepo.GetMemberIds(ctx, record.ReceiverId) {
				m.UnreadStorage.PipeIncr(ctx, pipe, constant.ChatGroupMode, record.ReceiverId, uid)
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
				"event": constant.SubEventImMessage,
				"data": jsonutil.Encode(map[string]any{
					"sender_id":   uid,
					"receiver_id": item.ReceiverId,
					"dialog_type": item.DialogType,
					"record_id":   item.RecordId,
				}),
			})
			pipe.Publish(ctx, constant.ImTopicChat, data)
		}

		return nil
	})

	return nil
}

func (m *MessageUseCase) SendMixedMessage(ctx context.Context, uid int, req *v1Pb.MixedMessageRequest) error {

	items := make([]*entity.DialogRecordExtraMixedItem, 0)

	for _, item := range req.Items {
		items = append(items, &entity.DialogRecordExtraMixedItem{
			Type:    int(item.Type),
			Content: item.Content,
		})
	}

	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgTypeMixed,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra:      jsonutil.Encode(entity.DialogRecordExtraMixed{Items: items}),
		QuoteId:    req.QuoteId,
	}

	return m.save(ctx, data)
}

func (m *MessageUseCase) SendSysOther(ctx context.Context, data *model.Message) error {
	return m.save(ctx, data)
}

func (m *MessageUseCase) Revoke(ctx context.Context, uid int, msgId string) error {
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
		"event": constant.SubEventImMessageRevoke,
		"data": jsonutil.Encode(map[string]any{
			"msg_id": record.MsgId,
		}),
	}

	m.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(body))

	return nil
}

func (m *MessageUseCase) Vote(ctx context.Context, uid int, msgId int, optionsValue string) (*repo.VoteStatistics, error) {
	db := m.Source.Db().WithContext(ctx)
	query := db.Table("messages")
	query.Select([]string{
		"messages.receiver_id", "messages.dialog_type", "messages.msg_type",
		"vote.id as vote_id", "vote.id as record_id", "vote.answer_mode", "vote.answer_option",
		"vote.answer_num", "vote.status as vote_status",
	})
	query.Joins("LEFT JOIN message_votes as vote on vote.record_id = messages.id")
	query.Where("messages.id = ?", msgId)

	var vote entity.QueryVoteModel
	if err := query.Take(&vote).Error; err != nil {
		return nil, err
	}

	if vote.MsgType != constant.ChatMsgTypeVote {
		return nil, fmt.Errorf("текущая запись не относится к информации о голосовании %d", vote.MsgType)
	}

	if vote.DialogType == constant.ChatGroupMode {
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

	if vote.AnswerMode == constant.VoteAnswerModeSingleChoice {
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

func (m *MessageUseCase) save(ctx context.Context, data *model.Message) error {
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
	case constant.ChatMsgTypeText:
		option["text"] = strutil.MtSubstr(strutil.ReplaceImgAll(data.Content), 0, 300)
	default:
		if value, ok := constant.ChatMsgTypeMapping[data.MsgType]; ok {
			option["text"] = value
		} else {
			option["text"] = "Неизвестно"
		}
	}

	m.afterHandle(ctx, data, option)

	return nil
}

func (m *MessageUseCase) loadReply(_ context.Context, data *model.Message) {
	if data.QuoteId == "" {
		return
	}
	if data.Extra == "" {
		data.Extra = "{}"
	}
	extra := make(map[string]any)
	err := jsonutil.Decode(data.Extra, &extra)
	if err != nil {
		logger.Errorf("MessageUseCase Json Decode err: %s", err.Error())
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
	reply := entity.Reply{
		UserId:   record.UserId,
		Username: user.Username,
		MsgType:  1,
		Content:  record.Content,
		MsgId:    record.MsgId,
	}
	if record.MsgType != constant.ChatMsgTypeText {
		reply.Content = "Неизвестно"
		if value, ok := constant.ChatMsgTypeMapping[record.MsgType]; ok {
			reply.Content = value
		}
	}
	extra["reply"] = reply
	data.Extra = jsonutil.Encode(extra)
}

func (m *MessageUseCase) loadSequence(ctx context.Context, data *model.Message) {
	if data.DialogType == constant.ChatGroupMode {
		data.Sequence = m.Sequence.Get(ctx, 0, data.ReceiverId)
	} else {
		data.Sequence = m.Sequence.Get(ctx, data.UserId, data.ReceiverId)
	}
}

func (m *MessageUseCase) afterHandle(ctx context.Context, record *model.Message, opt map[string]string) {
	if record.DialogType == constant.ChatPrivateMode {
		m.UnreadStorage.Incr(ctx, constant.ChatPrivateMode, record.UserId, record.ReceiverId)
		if record.MsgType == constant.ChatMsgSysText {
			m.UnreadStorage.Incr(ctx, 1, record.ReceiverId, record.UserId)
		}
	} else if record.DialogType == constant.ChatGroupMode {
		pipe := m.Source.Redis().Pipeline()
		for _, uid := range m.GroupChatMemberRepo.GetMemberIds(ctx, record.ReceiverId) {
			if uid != record.UserId {
				m.UnreadStorage.PipeIncr(ctx, pipe, constant.ChatGroupMode, record.ReceiverId, uid)
			}
		}
		_, _ = pipe.Exec(ctx)
	}

	_ = m.MessageStorage.Set(ctx, record.DialogType, record.UserId, record.ReceiverId, &cache.LastCacheMessage{
		Content:  opt["text"],
		Datetime: timeutil.DateTime(),
	})

	content := jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessage,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"dialog_type": record.DialogType,
			"record_id":   record.Id,
		}),
	})

	if record.DialogType == constant.ChatPrivateMode {
		sids := m.ServerStorage.All(ctx, 1)
		if len(sids) > 3 {
			pipe := m.Source.Redis().Pipeline()
			for _, sid := range sids {
				for _, uid := range []int{record.UserId, record.ReceiverId} {
					if !m.ClientStorage.IsCurrentServerOnline(ctx, sid, constant.ImChannelChat, strconv.Itoa(uid)) {
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
		logger.Errorf("Ошибка отправки уведомления %s", err.Error())
	}
}

type SendLogin struct {
	Ip      string
	Agent   string
	Address string
}

func (m *MessageUseCase) SendLogin(ctx context.Context, uid int, req *SendLogin) error {
	bot, err := m.BotRepo.GetLoginBot(ctx)
	if err != nil {
		return err
	}

	data := &model.Message{
		DialogType: constant.ChatPrivateMode,
		MsgType:    constant.ChatMsgTypeLogin,
		UserId:     bot.UserId,
		ReceiverId: uid,
		Extra: jsonutil.Encode(&entity.DialogRecordExtraLogin{
			IP:       req.Ip,
			Agent:    req.Agent,
			Address:  req.Address,
			Datetime: timeutil.DateTime(),
		}),
	}

	return m.save(ctx, data)
}

func (m *MessageUseCase) SendSticker(ctx context.Context, uid int, req *v1Pb.StickerMessageRequest) error {
	var sticker model.StickerItem
	if err := m.Source.Db().First(&sticker, "id = ? and user_id = ?", req.StickerId, uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("информация о смайлик не существует")
		}
		return err
	}

	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgTypeImage,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&entity.DialogRecordExtraImage{
			Url:    sticker.Url,
			Width:  0,
			Height: 0,
		}),
	}

	return m.save(ctx, data)
}

func (m *MessageUseCase) SendCode(ctx context.Context, uid int, req *v1Pb.CodeMessageRequest) error {
	data := &model.Message{
		DialogType: int(req.Receiver.DialogType),
		MsgType:    constant.ChatMsgTypeCode,
		UserId:     uid,
		ReceiverId: int(req.Receiver.ReceiverId),
		Extra: jsonutil.Encode(&entity.DialogRecordExtraCode{
			Lang: req.Lang,
			Code: req.Code,
		}),
	}
	return m.save(ctx, data)
}

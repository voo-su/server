package usecase

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/minio"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type GroupChatUseCase struct {
	Conf               *config.Config
	Locale             locale.ILocale
	Source             *infrastructure.Source
	GroupChatRepo      *postgresRepo.GroupChatRepository
	MemberRepo         *postgresRepo.GroupChatMemberRepository
	SequenceRepo       *postgresRepo.SequenceRepository
	RelationCache      *redisRepo.RelationCacheRepository
	RedisLockCacheRepo *redisRepo.RedisLockCacheRepository
	Minio              minio.IMinio
}

func NewGroupChatUseCase(
	conf *config.Config,
	locale locale.ILocale,
	source *infrastructure.Source,
	groupChatRepo *postgresRepo.GroupChatRepository,
	memberRepo *postgresRepo.GroupChatMemberRepository,
	sequenceRepo *postgresRepo.SequenceRepository,
	relationCache *redisRepo.RelationCacheRepository,
	redisLockCacheRepo *redisRepo.RedisLockCacheRepository,
	minio minio.IMinio,
) *GroupChatUseCase {
	return &GroupChatUseCase{
		Conf:               conf,
		Locale:             locale,
		Source:             source,
		GroupChatRepo:      groupChatRepo,
		MemberRepo:         memberRepo,
		SequenceRepo:       sequenceRepo,
		RelationCache:      relationCache,
		RedisLockCacheRepo: redisLockCacheRepo,
		Minio:              minio,
	}
}

type GroupCreateOpt struct {
	UserId      int
	Name        string
	Avatar      string
	Description string
	MemberIds   []int
}

func (g *GroupChatUseCase) Create(ctx context.Context, opt *GroupCreateOpt) (int, error) {
	var (
		err     error
		members []*postgresModel.GroupChatMember
		chats   []*postgresModel.Chat
	)
	uids := sliceutil.Unique(append(opt.MemberIds, opt.UserId))
	group := &postgresModel.GroupChat{
		CreatorId:   opt.UserId,
		Name:        opt.Name,
		Description: opt.Description,
		Avatar:      opt.Avatar,
		MaxNum:      constant.GroupMemberMaxNum,
	}
	joinTime := time.Now()
	err = g.Source.Postgres().Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(group).Error; err != nil {
			return err
		}

		addMembers := make([]entity.MessageExtraGroupMembers, 0, len(opt.MemberIds))
		tx.Table("users").
			Select("id AS user_id", "username").
			Where("id IN ?", opt.MemberIds).
			Scan(&addMembers)
		for _, val := range uids {
			leader := 0
			if opt.UserId == val {
				leader = 2
			}

			members = append(members, &postgresModel.GroupChatMember{
				GroupId:  group.Id,
				UserId:   val,
				Leader:   leader,
				JoinTime: joinTime,
			})
			chats = append(chats, &postgresModel.Chat{
				ChatType:   2,
				UserId:     val,
				ReceiverId: group.Id,
			})
		}
		if err = tx.Create(members).Error; err != nil {
			return err
		}

		if err = tx.Create(chats).Error; err != nil {
			return err
		}

		var user postgresModel.User
		if err = tx.Table("users").
			Where("id = ?", opt.UserId).
			Scan(&user).
			Error; err != nil {
			return err
		}

		message := &postgresModel.Message{
			MsgId:      strutil.NewMsgId(),
			ChatType:   constant.ChatGroupMode,
			ReceiverId: group.Id,
			MsgType:    constant.ChatMsgSysGroupCreate,
			Sequence:   g.SequenceRepo.Get(ctx, 0, group.Id),
			Extra: jsonutil.Encode(entity.MessageExtraGroupCreate{
				OwnerId:   user.Id,
				OwnerName: user.Username,
				Members:   addMembers,
			}),
		}
		if err = tx.Create(message).Error; err != nil {
			return err
		}

		return nil
	})

	body := map[string]any{
		"event": constant.SubEventGroupChatJoin,
		"data": jsonutil.Encode(map[string]any{
			"group_id": group.Id,
			"uids":     uids,
		}),
	}

	g.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(body))

	return group.Id, err
}

type GroupUpdateOpt struct {
	GroupId     int
	Name        string
	Description string
}

func (g *GroupChatUseCase) Update(ctx context.Context, opt *GroupUpdateOpt) error {
	_, err := g.GroupChatRepo.UpdateById(ctx, opt.GroupId, map[string]any{
		"group_name":  opt.Name,
		"description": opt.Description,
	})
	return err
}

func (g *GroupChatUseCase) Dismiss(ctx context.Context, groupId int, uid int) error {
	err := g.Source.Postgres().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&postgresModel.GroupChat{Id: groupId, CreatorId: uid}).Updates(&postgresModel.GroupChat{
			IsDismiss: 1,
		}).Error; err != nil {
			return err
		}

		if err := g.Source.Postgres().
			Model(&postgresModel.GroupChatMember{}).
			Where("group_id = ?", groupId).
			Updates(&postgresModel.GroupChatMember{
				IsQuit: 1,
			}).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (g *GroupChatUseCase) Secede(ctx context.Context, groupId int, uid int) error {
	var info postgresModel.GroupChatMember
	if err := g.Source.Postgres().
		Where("group_id = ? AND user_id = ? AND is_quit = 0", groupId, uid).
		First(&info).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(g.Locale.Localize("data_not_found"))
		}
		return err
	}
	if info.Leader == 2 {
		return errors.New(g.Locale.Localize("group_owner_cannot_leave"))
	}

	var user postgresModel.User
	if err := g.Source.Postgres().
		Table("users").
		Select("id, username").
		Where("id = ?", uid).
		First(&user).
		Error; err != nil {
		return err
	}

	message := &postgresModel.Message{
		MsgId:      strutil.NewMsgId(),
		ChatType:   constant.ChatGroupMode,
		ReceiverId: groupId,
		MsgType:    constant.ChatMsgSysGroupMemberQuit,
		Sequence:   g.SequenceRepo.Get(ctx, 0, groupId),
		Extra: jsonutil.Encode(&entity.MessageExtraGroupMemberQuit{
			OwnerId:   user.Id,
			OwnerName: user.Username,
		}),
	}
	err := g.Source.Postgres().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&postgresModel.GroupChatMember{}).
			Where("group_id = ? AND user_id = ?", groupId, uid).
			Updates(&postgresModel.GroupChatMember{
				IsQuit: 1,
			}).Error
		if err != nil {
			return err
		}

		if err = tx.Create(message).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	g.RelationCache.DelGroupRelation(ctx, uid, groupId)
	g.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventGroupChatJoin,
		"data": jsonutil.Encode(map[string]any{
			"type":     2,
			"group_id": groupId,
			"uids":     []int{uid},
		}),
	}))

	g.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessage,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   message.UserId,
			"receiver_id": message.ReceiverId,
			"chat_type":   message.ChatType,
			"message_id":  message.Id,
		}),
	}))

	return nil
}

type GroupInviteOpt struct {
	UserId    int
	GroupId   int
	MemberIds []int
}

func (g *GroupChatUseCase) Invite(ctx context.Context, opt *GroupInviteOpt) error {
	var (
		err        error
		addMembers []*postgresModel.GroupChatMember
		addChat    []*postgresModel.Chat
		updateChat []int
		chats      []*postgresModel.Chat
		db         = g.Source.Postgres().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range g.MemberRepo.GetMemberIds(ctx, opt.GroupId) {
		m[value] = struct{}{}
	}

	listHash := make(map[int]*postgresModel.Chat)
	db.Select("id", "user_id", "is_delete").
		Where("user_id IN ? AND receiver_id = ? AND chat_type = ?", opt.MemberIds, opt.GroupId, 2).
		Find(&chats)
	for _, item := range chats {
		listHash[item.UserId] = item
	}

	mids := make([]int, 0)
	mids = append(mids, opt.MemberIds...)
	mids = append(mids, opt.UserId)

	memberItems := make([]*postgresModel.User, 0)
	if err = db.Table("users").
		Select("id, username").
		Where("id in ?", mids).
		Scan(&memberItems).
		Error; err != nil {
		return err
	}

	memberMaps := make(map[int]*postgresModel.User)
	for _, item := range memberItems {
		memberMaps[item.Id] = item
	}

	members := make([]entity.MessageExtraGroupMembers, 0)
	for _, value := range opt.MemberIds {
		members = append(members, entity.MessageExtraGroupMembers{
			UserId:   value,
			Username: memberMaps[value].Username,
		})

		if _, ok := m[value]; !ok {
			addMembers = append(addMembers, &postgresModel.GroupChatMember{
				GroupId:  opt.GroupId,
				UserId:   value,
				JoinTime: time.Now(),
			})
		}

		if item, ok := listHash[value]; !ok {
			addChat = append(addChat, &postgresModel.Chat{
				ChatType:   constant.ChatGroupMode,
				UserId:     value,
				ReceiverId: opt.GroupId,
			})
		} else if item.IsDelete == 1 {
			updateChat = append(updateChat, item.Id)
		}
	}
	if len(addMembers) == 0 {
		return errors.New(g.Locale.Localize("all_invited_contacts_are_group_members"))
	}

	message := &postgresModel.Message{
		MsgId:      strutil.NewMsgId(),
		ChatType:   constant.ChatGroupMode,
		ReceiverId: opt.GroupId,
		MsgType:    constant.ChatMsgSysGroupMemberJoin,
		Sequence:   g.SequenceRepo.Get(ctx, 0, opt.GroupId),
	}

	message.Extra = jsonutil.Encode(&entity.MessageExtraGroupJoin{
		OwnerId:   memberMaps[opt.UserId].Id,
		OwnerName: memberMaps[opt.UserId].Username,
		Members:   members,
	})

	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&postgresModel.GroupChatMember{}, "group_id = ? AND user_id IN ? AND is_quit = ?", opt.GroupId, opt.MemberIds, constant.GroupMemberQuitStatusYes)
		if err = tx.Create(&addMembers).Error; err != nil {
			return err
		}

		if len(addChat) > 0 {
			if err = tx.Create(&addChat).Error; err != nil {
				return err
			}
		}

		if len(updateChat) > 0 {
			tx.Model(&postgresModel.Chat{}).Where("id IN ?", updateChat).Updates(map[string]any{
				"is_delete":  0,
				"created_at": timeutil.DateTime(),
			})
		}
		if err = tx.Create(message).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	g.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventGroupChatJoin,
		"data": jsonutil.Encode(map[string]any{
			"type":     1,
			"group_id": opt.GroupId,
			"uids":     opt.MemberIds,
		}),
	}))

	g.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessage,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   message.UserId,
			"receiver_id": message.ReceiverId,
			"chat_type":   message.ChatType,
			"message_id":  message.Id,
		}),
	}))
	return nil
}

type GroupRemoveMembersOpt struct {
	UserId    int
	GroupId   int
	MemberIds []int
}

func (g *GroupChatUseCase) RemoveMember(ctx context.Context, opt *GroupRemoveMembersOpt) error {
	var num int64
	if err := g.Source.Postgres().
		Model(&postgresModel.GroupChatMember{}).
		Where("group_id = ? AND user_id in ? AND is_quit = 0", opt.GroupId, opt.MemberIds).
		Count(&num).
		Error; err != nil {
		return err
	}
	if int(num) != len(opt.MemberIds) {
		return errors.New(g.Locale.Localize("delete_failed"))
	}

	mids := make([]int, 0)
	mids = append(mids, opt.MemberIds...)
	mids = append(mids, opt.UserId)
	memberItems := make([]*postgresModel.User, 0)

	if err := g.Source.Postgres().
		Table("users").
		Select("id, username").
		Where("id IN ?", mids).
		Scan(&memberItems).
		Error; err != nil {
		return err
	}

	memberMaps := make(map[int]*postgresModel.User)
	for _, item := range memberItems {
		memberMaps[item.Id] = item
	}

	members := make([]entity.MessageExtraGroupMembers, 0)
	for _, value := range opt.MemberIds {
		members = append(members, entity.MessageExtraGroupMembers{
			UserId:   value,
			Username: memberMaps[value].Username,
		})
	}

	message := &postgresModel.Message{
		MsgId:      strutil.NewMsgId(),
		Sequence:   g.SequenceRepo.Get(ctx, 0, opt.GroupId),
		ChatType:   constant.ChatGroupMode,
		ReceiverId: opt.GroupId,
		MsgType:    constant.ChatMsgSysGroupMemberKicked,
		Extra: jsonutil.Encode(&entity.MessageExtraGroupMemberKicked{
			OwnerId:   memberMaps[opt.UserId].Id,
			OwnerName: memberMaps[opt.UserId].Username,
			Members:   members,
		}),
	}
	err := g.Source.Postgres().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&postgresModel.GroupChatMember{}).
			Where("group_id = ? AND user_id in ? AND is_quit = 0", opt.GroupId, opt.MemberIds).
			Updates(map[string]any{
				"is_quit":    1,
				"updated_at": time.Now(),
			}).Error
		if err != nil {
			return err
		}

		return tx.Create(message).Error
	})
	if err != nil {
		return err
	}

	g.RelationCache.BatchDelGroupRelation(ctx, opt.MemberIds, opt.GroupId)

	_, _ = g.Source.Redis().Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
			"event": constant.SubEventGroupChatJoin,
			"data": jsonutil.Encode(map[string]any{
				"type":     2,
				"group_id": opt.GroupId,
				"uids":     opt.MemberIds,
			}),
		}))

		pipe.Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
			"event": constant.SubEventImMessage,
			"data": jsonutil.Encode(map[string]any{
				"sender_id":   int64(message.UserId),
				"receiver_id": int64(message.ReceiverId),
				"chat_type":   message.ChatType,
				"message_id":  int64(message.Id),
			}),
		}))
		return nil
	})

	return nil
}

type Session struct {
	ReceiverID int `json:"receiver_id"`
	IsDisturb  int `json:"is_disturb"`
}

func (g *GroupChatUseCase) List(userId int) ([]*entity.GroupItem, error) {
	tx := g.Source.Postgres().
		Table("group_chat_members").
		Select("gc.id AS id, gc.group_name AS group_name, gc.avatar AS avatar, gc.description AS description, group_chat_members.leader AS leader, gc.creator_id AS creator_id").
		Joins("LEFT JOIN group_chats gc on gc.id = group_chat_members.group_id").
		Where("group_chat_members.user_id = ? AND group_chat_members.is_quit = ?", userId, 0).
		Order("group_chat_members.created_at desc")

	items := make([]*entity.GroupItem, 0)
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	length := len(items)
	if length == 0 {
		return items, nil
	}

	ids := make([]int, 0, length)
	for i := range items {
		ids = append(ids, items[i].Id)
	}

	query := g.Source.Postgres().
		Table("chats").
		Select("receiver_id,is_disturb").
		Where("chat_type = ? AND receiver_id in ?", 2, ids)
	list := make([]*Session, 0)
	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}

	hash := make(map[int]*Session)
	for i := range list {
		hash[list[i].ReceiverID] = list[i]
	}
	for i := range items {
		if value, ok := hash[items[i].Id]; ok {
			items[i].IsDisturb = value.IsDisturb
		}
	}

	return items, nil
}
func (g *GroupChatUseCase) UpdateAvatar(ctx context.Context, uid int, avatar string) error {
	_, err := g.GroupChatRepo.UpdateById(ctx, uid, map[string]any{
		"avatar": avatar,
	})
	if err != nil {
		return errors.New(g.Locale.Localize("personal_info_update_error"))
	}

	return nil
}

func (g *GroupChatUseCase) UpdateAvatarFullPath(ctx context.Context, uid int, avatarPath string) error {
	_, err := g.GroupChatRepo.UpdateById(ctx, uid, map[string]any{
		"avatar": g.Minio.PublicUrl(g.Conf.Minio.GetBucket(), avatarPath),
	})
	if err != nil {
		return errors.New(g.Locale.Localize("personal_info_update_error"))
	}

	return nil
}

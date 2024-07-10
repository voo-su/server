package service

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type GroupChatService struct {
	*repo.Source
	Repo     *repo.GroupChat
	Member   *repo.GroupChatMember
	Relation *cache.Relation
	Sequence *repo.Sequence
}

func NewGroupChatService(
	source *repo.Source,
	repo *repo.GroupChat,
	member *repo.GroupChatMember,
	relation *cache.Relation,
	sequence *repo.Sequence,
) *GroupChatService {
	return &GroupChatService{
		Source:   source,
		Repo:     repo,
		Member:   member,
		Relation: relation,
		Sequence: sequence,
	}
}

type GroupCreateOpt struct {
	UserId      int
	Name        string
	Avatar      string
	Description string
	MemberIds   []int
}

func (g *GroupChatService) Create(ctx context.Context, opt *GroupCreateOpt) (int, error) {
	var (
		err        error
		members    []*model.GroupChatMember
		dialogList []*model.Dialog
	)
	uids := sliceutil.Unique(append(opt.MemberIds, opt.UserId))
	group := &model.GroupChat{
		CreatorId:   opt.UserId,
		Name:        opt.Name,
		Description: opt.Description,
		Avatar:      opt.Avatar,
		MaxNum:      model.GroupMemberMaxNum,
	}
	joinTime := time.Now()
	err = g.Source.Db().Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(group).Error; err != nil {
			return err
		}

		addMembers := make([]model.DialogRecordExtraGroupMembers, 0, len(opt.MemberIds))
		tx.Table("users").Select("id as user_id", "username").Where("id in ?", opt.MemberIds).
			Scan(&addMembers)
		for _, val := range uids {
			leader := 0
			if opt.UserId == val {
				leader = 2
			}

			members = append(members, &model.GroupChatMember{
				GroupId:  group.Id,
				UserId:   val,
				Leader:   leader,
				JoinTime: joinTime,
			})
			dialogList = append(dialogList, &model.Dialog{
				DialogType: 2,
				UserId:     val,
				ReceiverId: group.Id,
			})
		}
		if err = tx.Create(members).Error; err != nil {
			return err
		}

		if err = tx.Create(dialogList).Error; err != nil {
			return err
		}

		var user model.User
		err = tx.Table("users").Where("id = ?", opt.UserId).Scan(&user).Error
		if err != nil {
			return err
		}

		record := &model.Message{
			MsgId:      strutil.NewMsgId(),
			DialogType: entity.ChatGroupMode,
			ReceiverId: group.Id,
			MsgType:    entity.ChatMsgSysGroupCreate,
			Sequence:   g.Sequence.Get(ctx, 0, group.Id),
			Extra: jsonutil.Encode(model.DialogRecordExtraGroupCreate{
				OwnerId:   user.Id,
				OwnerName: user.Username,
				Members:   addMembers,
			}),
		}
		if err = tx.Create(record).Error; err != nil {
			return err
		}

		return nil
	})

	body := map[string]any{
		"event": entity.SubEventGroupChatJoin,
		"data": jsonutil.Encode(map[string]any{
			"group_id": group.Id,
			"uids":     uids,
		}),
	}

	g.Source.Redis().Publish(ctx, entity.ImTopicChat, jsonutil.Encode(body))

	return group.Id, err
}

type GroupUpdateOpt struct {
	GroupId     int
	Name        string
	Avatar      string
	Description string
}

func (g *GroupChatService) Update(ctx context.Context, opt *GroupUpdateOpt) error {
	_, err := g.Repo.UpdateById(ctx, opt.GroupId, map[string]any{
		"group_name":  opt.Name,
		"avatar":      opt.Avatar,
		"description": opt.Description,
	})
	return err
}

func (g *GroupChatService) Dismiss(ctx context.Context, groupId int, uid int) error {
	err := g.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.GroupChat{Id: groupId, CreatorId: uid}).Updates(&model.GroupChat{
			IsDismiss: 1,
		}).Error; err != nil {
			return err
		}

		if err := g.Source.Db().Model(&model.GroupChatMember{}).Where("group_id = ?", groupId).Updates(&model.GroupChatMember{
			IsQuit: 1,
		}).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (g *GroupChatService) Secede(ctx context.Context, groupId int, uid int) error {
	var info model.GroupChatMember
	if err := g.Source.Db().Where("group_id = ? AND user_id = ? and is_quit = 0", groupId, uid).First(&info).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("данных не существует")
		}
		return err
	}
	if info.Leader == 2 {
		return errors.New("владелец группы не может покинуть группу")
	}

	var user model.User
	err := g.Source.Db().
		Table("users").
		Select("id, username").
		Where("id = ?", uid).
		First(&user).
		Error
	if err != nil {
		return err
	}

	record := &model.Message{
		MsgId:      strutil.NewMsgId(),
		DialogType: entity.ChatGroupMode,
		ReceiverId: groupId,
		MsgType:    entity.ChatMsgSysGroupMemberQuit,
		Sequence:   g.Sequence.Get(ctx, 0, groupId),
		Extra: jsonutil.Encode(&model.DialogRecordExtraGroupMemberQuit{
			OwnerId:   user.Id,
			OwnerName: user.Username,
		}),
	}
	err = g.Source.Db().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.GroupChatMember{}).
			Where("group_id = ? AND user_id = ?", groupId, uid).
			Updates(&model.GroupChatMember{
				IsQuit: 1,
			}).Error
		if err != nil {
			return err
		}

		if err = tx.Create(record).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	g.Relation.DelGroupRelation(ctx, uid, groupId)
	g.Source.Redis().Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventGroupChatJoin,
		"data": jsonutil.Encode(map[string]any{
			"type":     2,
			"group_id": groupId,
			"uids":     []int{uid},
		}),
	}))

	g.Source.Redis().Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventImMessage,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"dialog_type": record.DialogType,
			"record_id":   record.Id,
		}),
	}))

	return nil
}

type GroupInviteOpt struct {
	UserId    int
	GroupId   int
	MemberIds []int
}

func (g *GroupChatService) Invite(ctx context.Context, opt *GroupInviteOpt) error {
	var (
		err              error
		addMembers       []*model.GroupChatMember
		addDialogList    []*model.Dialog
		updateDialogList []int
		dialogList       []*model.Dialog
		db               = g.Source.Db().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range g.Member.GetMemberIds(ctx, opt.GroupId) {
		m[value] = struct{}{}
	}

	listHash := make(map[int]*model.Dialog)
	db.Select("id", "user_id", "is_delete").
		Where("user_id in ? and receiver_id = ? and dialog_type = 2", opt.MemberIds, opt.GroupId).
		Find(&dialogList)
	for _, item := range dialogList {
		listHash[item.UserId] = item
	}
	mids := make([]int, 0)
	mids = append(mids, opt.MemberIds...)
	mids = append(mids, opt.UserId)

	memberItems := make([]*model.User, 0)
	err = db.Table("users").Select("id, username").Where("id in ?", mids).Scan(&memberItems).Error
	if err != nil {
		return err
	}

	memberMaps := make(map[int]*model.User)
	for _, item := range memberItems {
		memberMaps[item.Id] = item
	}

	members := make([]model.DialogRecordExtraGroupMembers, 0)
	for _, value := range opt.MemberIds {
		members = append(members, model.DialogRecordExtraGroupMembers{
			UserId:   value,
			Username: memberMaps[value].Username,
		})
		if _, ok := m[value]; !ok {
			addMembers = append(addMembers, &model.GroupChatMember{
				GroupId:  opt.GroupId,
				UserId:   value,
				JoinTime: time.Now(),
			})
		}
		if item, ok := listHash[value]; !ok {
			addDialogList = append(addDialogList, &model.Dialog{
				DialogType: entity.ChatGroupMode,
				UserId:     value,
				ReceiverId: opt.GroupId,
			})
		} else if item.IsDelete == 1 {
			updateDialogList = append(updateDialogList, item.Id)
		}
	}
	if len(addMembers) == 0 {
		return errors.New("все приглашенные контакты стали участниками группы")
	}

	record := &model.Message{
		MsgId:      strutil.NewMsgId(),
		DialogType: entity.ChatGroupMode,
		ReceiverId: opt.GroupId,
		MsgType:    entity.ChatMsgSysGroupMemberJoin,
		Sequence:   g.Sequence.Get(ctx, 0, opt.GroupId),
	}
	record.Extra = jsonutil.Encode(&model.DialogRecordExtraGroupJoin{
		OwnerId:   memberMaps[opt.UserId].Id,
		OwnerName: memberMaps[opt.UserId].Username,
		Members:   members,
	})
	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&model.GroupChatMember{}, "group_id = ? and user_id in ? and is_quit = ?", opt.GroupId, opt.MemberIds, model.GroupMemberQuitStatusYes)
		if err = tx.Create(&addMembers).Error; err != nil {
			return err
		}

		if len(addDialogList) > 0 {
			if err = tx.Create(&addDialogList).Error; err != nil {
				return err
			}
		}

		if len(updateDialogList) > 0 {
			tx.Model(&model.Dialog{}).Where("id in ?", updateDialogList).Updates(map[string]any{
				"is_delete":  0,
				"created_at": timeutil.DateTime(),
			})
		}
		if err = tx.Create(record).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	g.Source.Redis().Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventGroupChatJoin,
		"data": jsonutil.Encode(map[string]any{
			"type":     1,
			"group_id": opt.GroupId,
			"uids":     opt.MemberIds,
		}),
	}))

	g.Source.Redis().Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventImMessage,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"dialog_type": record.DialogType,
			"record_id":   record.Id,
		}),
	}))
	return nil
}

type GroupRemoveMembersOpt struct {
	UserId    int
	GroupId   int
	MemberIds []int
}

func (g *GroupChatService) RemoveMember(ctx context.Context, opt *GroupRemoveMembersOpt) error {
	var num int64
	if err := g.Source.Db().Model(&model.GroupChatMember{}).Where("group_id = ? and user_id in ? and is_quit = 0", opt.GroupId, opt.MemberIds).Count(&num).Error; err != nil {
		return err
	}
	if int(num) != len(opt.MemberIds) {
		return errors.New("не удалось удалить")
	}

	mids := make([]int, 0)
	mids = append(mids, opt.MemberIds...)
	mids = append(mids, opt.UserId)
	memberItems := make([]*model.User, 0)
	err := g.Source.Db().
		Table("users").
		Select("id, username").
		Where("id in ?", mids).
		Scan(&memberItems).Error
	if err != nil {
		return err
	}

	memberMaps := make(map[int]*model.User)
	for _, item := range memberItems {
		memberMaps[item.Id] = item
	}

	members := make([]model.DialogRecordExtraGroupMembers, 0)
	for _, value := range opt.MemberIds {
		members = append(members, model.DialogRecordExtraGroupMembers{
			UserId:   value,
			Username: memberMaps[value].Username,
		})
	}

	record := &model.Message{
		MsgId:      strutil.NewMsgId(),
		Sequence:   g.Sequence.Get(ctx, 0, opt.GroupId),
		DialogType: entity.ChatGroupMode,
		ReceiverId: opt.GroupId,
		MsgType:    entity.ChatMsgSysGroupMemberKicked,
		Extra: jsonutil.Encode(&model.DialogRecordExtraGroupMemberKicked{
			OwnerId:   memberMaps[opt.UserId].Id,
			OwnerName: memberMaps[opt.UserId].Username,
			Members:   members,
		}),
	}
	err = g.Source.Db().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.GroupChatMember{}).
			Where("group_id = ? and user_id in ? and is_quit = 0", opt.GroupId, opt.MemberIds).
			Updates(map[string]any{
				"is_quit":    1,
				"updated_at": time.Now(),
			}).Error
		if err != nil {
			return err
		}

		return tx.Create(record).Error
	})
	if err != nil {
		return err
	}

	g.Relation.BatchDelGroupRelation(ctx, opt.MemberIds, opt.GroupId)

	_, _ = g.Source.Redis().Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
			"event": entity.SubEventGroupChatJoin,
			"data": jsonutil.Encode(map[string]any{
				"type":     2,
				"group_id": opt.GroupId,
				"uids":     opt.MemberIds,
			}),
		}))
		pipe.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
			"event": entity.SubEventImMessage,
			"data": jsonutil.Encode(map[string]any{
				"sender_id":   int64(record.UserId),
				"receiver_id": int64(record.ReceiverId),
				"dialog_type": record.DialogType,
				"record_id":   int64(record.Id),
			}),
		}))
		return nil
	})

	return nil
}

type session struct {
	ReceiverID int `json:"receiver_id"`
	IsDisturb  int `json:"is_disturb"`
}

func (g *GroupChatService) List(userId int) ([]*model.GroupItem, error) {
	tx := g.Source.Db().Table("group_chat_members")
	tx.Select("gc.id AS id, gc.group_name AS group_name, gc.avatar AS avatar, gc.description AS description, group_chat_members.leader AS leader, gc.creator_id AS creator_id")
	tx.Joins("LEFT JOIN group_chats gc on gc.id = group_chat_members.group_id")
	tx.Where("group_chat_members.user_id = ? and group_chat_members.is_quit = ?", userId, 0)
	tx.Order("group_chat_members.created_at desc")

	items := make([]*model.GroupItem, 0)
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
	query := g.Source.Db().Table("dialogs")
	query.Select("receiver_id,is_disturb")
	query.Where("dialog_type = ? and receiver_id in ?", 2, ids)
	list := make([]*session, 0)
	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}

	hash := make(map[int]*session)
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

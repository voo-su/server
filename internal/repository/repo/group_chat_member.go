package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type GroupChatMember struct {
	repo.Repo[model.GroupChatMember]
	RelationCache *cache.RelationCache
}

func NewGroupMember(db *gorm.DB, relationCache *cache.RelationCache) *GroupChatMember {
	return &GroupChatMember{Repo: repo.NewRepo[model.GroupChatMember](db), RelationCache: relationCache}
}

func (g *GroupChatMember) IsMaster(ctx context.Context, gid, uid int) bool {
	exist, err := g.Repo.QueryExist(ctx, "group_id = ? AND user_id = ? AND leader = 2 AND is_quit = ?", gid, uid, constant.GroupMemberQuitStatusNo)

	return err == nil && exist
}

func (g *GroupChatMember) IsLeader(ctx context.Context, gid, uid int) bool {
	exist, err := g.Repo.QueryExist(ctx, "group_id = ? AND user_id = ? AND leader in (1,2) AND is_quit = ?", gid, uid, constant.GroupMemberQuitStatusNo)

	return err == nil && exist
}

func (g *GroupChatMember) IsMember(ctx context.Context, gid, uid int, cache bool) bool {
	if cache && g.RelationCache.IsGroupRelation(ctx, uid, gid) == nil {
		return true
	}

	exist, err := g.Repo.QueryExist(ctx, "group_id = ? AND user_id = ? AND is_quit = ?", gid, uid, constant.GroupMemberQuitStatusNo)
	if err != nil {
		return false
	}
	if exist {
		g.RelationCache.SetGroupRelation(ctx, uid, gid)
	}

	return exist
}

func (g *GroupChatMember) FindByUserId(ctx context.Context, gid, uid int) (*model.GroupChatMember, error) {
	member := &model.GroupChatMember{}
	err := g.Repo.Model(ctx).
		Where("group_id = ? AND user_id = ?", gid, uid).
		First(member).
		Error

	return member, err
}

func (g *GroupChatMember) GetMemberIds(ctx context.Context, groupId int) []int {
	var ids []int
	_ = g.Repo.Model(ctx).
		Select("user_id").
		Where("group_id = ? AND is_quit = ?", groupId, constant.GroupMemberQuitStatusNo).
		Scan(&ids)

	return ids
}

func (g *GroupChatMember) GetUserGroupIds(ctx context.Context, uid int) []int {
	var ids []int
	_ = g.Repo.Model(ctx).
		Where("user_id = ? AND is_quit = ?", uid, constant.GroupMemberQuitStatusNo).
		Pluck("group_id", &ids)

	return ids
}

func (g *GroupChatMember) CountMemberTotal(ctx context.Context, gid int) int64 {
	count, _ := g.Repo.QueryCount(ctx, "group_id = ? AND is_quit = ?", gid, constant.GroupMemberQuitStatusNo)
	return count
}

//func (g *GroupChatMember) GetMemberRemark(ctx context.Context, groupId int, userId int) string {
//	var remarks string
//	g.Repo.Model(ctx).Select("user_card").Where("group_id = ? AND user_id = ?", groupId, userId).Scan(&remarks)
//
//	return remarks
//}

func (g *GroupChatMember) GetMembers(ctx context.Context, groupId int) []*entity.MemberItem {
	fields := []string{
		"group_chat_members.id",
		"group_chat_members.leader",
		//"group_chat_members.user_card",
		"group_chat_members.user_id",
		"group_chat_members.is_mute",
		"users.avatar",
		"users.username",
		"users.gender",
		"users.about",
	}
	tx := g.Repo.Db.WithContext(ctx).Table("group_chat_members").
		Joins("LEFT JOIN users on users.id = group_chat_members.user_id").
		Where("group_chat_members.group_id = ? AND group_chat_members.is_quit = ?", groupId, constant.GroupMemberQuitStatusNo).
		Order("group_chat_members.leader desc")
	var items []*entity.MemberItem
	tx.Unscoped().
		Select(fields).
		Scan(&items)

	return items
}

type CountGroupMember struct {
	GroupId int `gorm:"column:group_id;"`
	Count   int `gorm:"column:count;"`
}

func (g *GroupChatMember) CountGroupMemberNum(ids []int) ([]*CountGroupMember, error) {
	var items []*CountGroupMember
	if err := g.Repo.Model(context.TODO()).
		Select("group_id,count(*) as count").
		Where("group_id in ? AND is_quit = ?", ids, constant.GroupMemberQuitStatusNo).
		Group("group_id").
		Scan(&items).
		Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (g *GroupChatMember) CheckUserGroup(ids []int, userId int) ([]int, error) {
	items := make([]int, 0)
	err := g.Repo.Model(context.TODO()).
		Select("group_id").
		Where("group_id in ? AND user_id = ? AND is_quit = ?", ids, userId, constant.GroupMemberQuitStatusNo).
		Scan(&items).
		Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

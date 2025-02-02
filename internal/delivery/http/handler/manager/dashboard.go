package manager

import (
	"log"
	managerPb "voo.su/api/http/pb/manager"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

type Dashboard struct {
	Locale           locale.ILocale
	UserUseCase      *usecase.UserUseCase
	MessageUseCase   *usecase.MessageUseCase
	GroupChatUseCase *usecase.GroupChatUseCase
}

func (d *Dashboard) Dashboard(ctx *ginutil.Context) error {
	users, err := d.UserUseCase.UserRepo.QueryCount(ctx.Ctx(), "is_bot = ?", 0)
	if err != nil {
		log.Println(err)
	}

	bots, err := d.UserUseCase.UserRepo.QueryCount(ctx.Ctx(), "is_bot = ?", 1)
	if err != nil {
		log.Println(err)
	}

	totalMessages, err := d.MessageUseCase.MessageRepo.QueryCount(ctx.Ctx(), "id > 1")
	if err != nil {
		log.Println(err)
	}

	groupChats, err := d.GroupChatUseCase.GroupChatRepo.QueryCount(ctx.Ctx(), "id > 1")
	if err != nil {
		log.Println(err)
	}

	groupMessages, err := d.MessageUseCase.MessageRepo.QueryCount(ctx.Ctx(), "chat_type = ?", 2)
	if err != nil {
		log.Println(err)
	}

	privateMessages, err := d.MessageUseCase.MessageRepo.QueryCount(ctx.Ctx(), "chat_type = ?", 1)
	if err != nil {
		log.Println(err)
	}

	return ctx.Success(&managerPb.ManagerDashboardResponse{
		Users:           users,
		Bots:            bots,
		TotalMessages:   totalMessages,
		GroupChats:      groupChats,
		GroupMessages:   groupMessages,
		PrivateMessages: privateMessages,
	})
}

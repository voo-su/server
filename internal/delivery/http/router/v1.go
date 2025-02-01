package router

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/delivery/http/handler"
	"voo.su/internal/delivery/http/middleware"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

func NewV1(conf *config.Config, locale locale.ILocale, router *gin.Engine, handler *handler.Handler, session *redisRepo.JwtTokenCacheRepository) {
	authorize := middleware.Auth(locale, constant.GuardHttpAuth, conf.App.Jwt.Secret, session)
	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", ginutil.HandlerFunc(handler.V1.Auth.Login))
			auth.POST("/verify", ginutil.HandlerFunc(handler.V1.Auth.Verify))
			auth.POST("/refresh", authorize, ginutil.HandlerFunc(handler.V1.Auth.Refresh))
			auth.POST("/logout", authorize, ginutil.HandlerFunc(handler.V1.Auth.Logout))
		}
		account := v1.Group("/account").Use(authorize)
		{
			account.GET("", ginutil.HandlerFunc(handler.V1.Account.Get))
			account.PUT("", ginutil.HandlerFunc(handler.V1.Account.ChangeDetail))
			account.PUT("/username", ginutil.HandlerFunc(handler.V1.Account.ChangeUsername))
			account.POST("/push", ginutil.HandlerFunc(handler.V1.Account.Push))
		}
		user := v1.Group("/search").Use(authorize)
		{
			user.GET("/users", ginutil.HandlerFunc(handler.V1.Search.Users))
			user.GET("/group-chats", ginutil.HandlerFunc(handler.V1.Search.GroupChats))
		}
		contact := v1.Group("/contacts").Use(authorize)
		{
			contact.GET("", ginutil.HandlerFunc(handler.V1.Contact.List))
			contact.GET("/get", ginutil.HandlerFunc(handler.V1.Contact.Get))
			contact.POST("/delete", ginutil.HandlerFunc(handler.V1.Contact.Delete))
			contact.POST("/requests", ginutil.HandlerFunc(handler.V1.ContactRequest.Create))
			contact.GET("/requests", ginutil.HandlerFunc(handler.V1.ContactRequest.List))
			contact.POST("/requests/accept", ginutil.HandlerFunc(handler.V1.ContactRequest.Accept))
			contact.POST("/requests/decline", ginutil.HandlerFunc(handler.V1.ContactRequest.Decline))
			contact.GET("/requests/unread-num", ginutil.HandlerFunc(handler.V1.ContactRequest.ApplyUnreadNum))
			contact.GET("/folders", ginutil.HandlerFunc(handler.V1.ContactFolder.List))
			contact.POST("/folders", ginutil.HandlerFunc(handler.V1.ContactFolder.Save))
			contact.POST("/folders/move", ginutil.HandlerFunc(handler.V1.ContactFolder.Move))
		}
		chat := v1.Group("/chats").Use(authorize)
		{
			chat.GET("", ginutil.HandlerFunc(handler.V1.Chat.List))
			chat.POST("/create", ginutil.HandlerFunc(handler.V1.Chat.Create))
			chat.POST("/delete", ginutil.HandlerFunc(handler.V1.Chat.Delete))
			chat.POST("/topping", ginutil.HandlerFunc(handler.V1.Chat.Top))
			chat.POST("/disturb", ginutil.HandlerFunc(handler.V1.Chat.Disturb))
			chat.POST("/unread/clear", ginutil.HandlerFunc(handler.V1.Chat.ClearUnreadMessage))
		}
		groupChat := v1.Group("/group-chats").Use(authorize)
		{
			groupChat.GET("", ginutil.HandlerFunc(handler.V1.GroupChat.GroupList))
			groupChat.POST("/create", ginutil.HandlerFunc(handler.V1.GroupChat.Create))
			groupChat.GET("/get", ginutil.HandlerFunc(handler.V1.GroupChat.Get))
			groupChat.POST("/invite", ginutil.HandlerFunc(handler.V1.GroupChat.Invite))
			groupChat.POST("/leave-chat", ginutil.HandlerFunc(handler.V1.GroupChat.SignOut))
			groupChat.POST("/setting", ginutil.HandlerFunc(handler.V1.GroupChat.Setting))
			groupChat.POST("/assign-admin", ginutil.HandlerFunc(handler.V1.GroupChat.AssignAdmin))
			groupChat.GET("/members", ginutil.HandlerFunc(handler.V1.GroupChat.Members))
			groupChat.GET("/members/invites", ginutil.HandlerFunc(handler.V1.GroupChat.GetInviteFriends))
			groupChat.POST("/members/remove", ginutil.HandlerFunc(handler.V1.GroupChat.RemoveMembers))
			groupChat.POST("/dismiss", ginutil.HandlerFunc(handler.V1.GroupChat.Dismiss))
			groupChat.POST("/overt", ginutil.HandlerFunc(handler.V1.GroupChat.Overt))
			groupChat.POST("/mute", ginutil.HandlerFunc(handler.V1.GroupChat.Mute))
			groupChat.GET("/ads", ginutil.HandlerFunc(handler.V1.GroupChatAds.List))
			groupChat.POST("/ads/edit", ginutil.HandlerFunc(handler.V1.GroupChatAds.CreateAndUpdate))
			groupChat.POST("/ads/delete", ginutil.HandlerFunc(handler.V1.GroupChatAds.Delete))
			groupChat.GET("/requests", ginutil.HandlerFunc(handler.V1.GroupChatRequest.List))
			groupChat.POST("/requests/create", ginutil.HandlerFunc(handler.V1.GroupChatRequest.Create))
			groupChat.POST("/requests/decline", ginutil.HandlerFunc(handler.V1.GroupChatRequest.Decline))
			groupChat.POST("/requests/agree", ginutil.HandlerFunc(handler.V1.GroupChatRequest.Agree))
			groupChat.GET("/requests/all", ginutil.HandlerFunc(handler.V1.GroupChatRequest.All))
			groupChat.GET("/requests/unread", ginutil.HandlerFunc(handler.V1.GroupChatRequest.RequestUnreadNum))
		}
		message := v1.Group("/messages").Use(authorize)
		{
			message.GET("", ginutil.HandlerFunc(handler.V1.Message.GetHistory))
			message.POST("/send", ginutil.HandlerFunc(handler.V1.Message.Send))
			message.GET("/file/download", ginutil.HandlerFunc(handler.V1.Message.Download))
			message.POST("/delete", ginutil.HandlerFunc(handler.V1.Message.Delete))
			message.POST("/revoke", ginutil.HandlerFunc(handler.V1.Message.Revoke))
			message.POST("/vote", ginutil.HandlerFunc(handler.V1.Message.Vote))
			message.POST("/vote/handle", ginutil.HandlerFunc(handler.V1.Message.HandleVote))
			message.GET("/stickers", ginutil.HandlerFunc(handler.V1.Sticker.CollectList))
			message.POST("/stickers", ginutil.HandlerFunc(handler.V1.Sticker.Upload))
			message.POST("/stickers/delete", ginutil.HandlerFunc(handler.V1.Sticker.DeleteCollect))
			message.GET("/stickers/system/list", ginutil.HandlerFunc(handler.V1.Sticker.SystemList))
			message.POST("/stickers/system/install", ginutil.HandlerFunc(handler.V1.Sticker.SetSystemSticker))
			message.POST("/collect", ginutil.HandlerFunc(handler.V1.Message.Collect))
		}
		upload := v1.Group("/upload").Use(authorize)
		{
			upload.POST("", ginutil.HandlerFunc(handler.V1.Upload.Upload))
			upload.POST("/multipart/initiate", ginutil.HandlerFunc(handler.V1.Upload.InitiateMultipart))
			upload.POST("/multipart", ginutil.HandlerFunc(handler.V1.Upload.MultipartUpload))
			upload.POST("/avatar", ginutil.HandlerFunc(handler.V1.Upload.Avatar))
		}
		bot := v1.Group("/bots").Use(authorize)
		{
			bot.POST("", ginutil.HandlerFunc(handler.V1.Bot.Create))
			bot.GET("", ginutil.HandlerFunc(handler.V1.Bot.List))
		}
		project := v1.Group("/projects").Use(authorize)
		{
			project.GET("", ginutil.HandlerFunc(handler.V1.Project.Projects))
			project.POST("/create", ginutil.HandlerFunc(handler.V1.Project.Create))
			project.GET("/detail", ginutil.HandlerFunc(handler.V1.Project.Detail))
			project.GET("/members", ginutil.HandlerFunc(handler.V1.Project.Members))
			project.GET("/members/invites", ginutil.HandlerFunc(handler.V1.Project.GetInviteFriends))
			project.POST("/invite", ginutil.HandlerFunc(handler.V1.Project.Invite))
			project.PUT("/task-type-name", ginutil.HandlerFunc(handler.V1.ProjectTask.TaskTypeName))
			project.GET("/tasks", ginutil.HandlerFunc(handler.V1.ProjectTask.Tasks))
			project.POST("/tasks/create", ginutil.HandlerFunc(handler.V1.ProjectTask.Create))
			project.GET("/tasks/detail", ginutil.HandlerFunc(handler.V1.ProjectTask.TaskDetail))
			project.PUT("/tasks/executor", ginutil.HandlerFunc(handler.V1.ProjectTask.Executor))
			project.PUT("/tasks/move", ginutil.HandlerFunc(handler.V1.ProjectTask.TaskMove))
			project.GET("/tasks/coexecutors", ginutil.HandlerFunc(handler.V1.ProjectTask.TaskCoexecutors))
			project.POST("/tasks/coexecutors/invite", ginutil.HandlerFunc(handler.V1.ProjectTask.TaskCoexecutorInvite))
			project.GET("/tasks/watchers", ginutil.HandlerFunc(handler.V1.ProjectTask.TaskWatchers))
			project.POST("/tasks/watchers/invite", ginutil.HandlerFunc(handler.V1.ProjectTask.TaskWatcherInvite))
			project.GET("/tasks/comments", ginutil.HandlerFunc(handler.V1.ProjectTaskComment.Comments))
			project.POST("/tasks/comments/create", ginutil.HandlerFunc(handler.V1.ProjectTaskComment.Create))
		}
	}
}

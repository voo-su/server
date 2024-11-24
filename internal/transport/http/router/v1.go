package router

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/http/handler"
	"voo.su/pkg/core"
	"voo.su/pkg/middleware"
)

func NewV1(router *gin.Engine, conf *config.Config, handler *handler.Handler, session *cache.JwtTokenStorage) {
	authorize := middleware.Auth(conf.App.Jwt.Secret, "api", session)
	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", core.HandlerFunc(handler.V1.Auth.Login))
			auth.POST("/verify", core.HandlerFunc(handler.V1.Auth.Verify))
			//auth.POST("/refresh", authorize, core.HandlerFunc(handler.V1.Auth.Refresh))
			auth.POST("/logout", authorize, core.HandlerFunc(handler.V1.Auth.Logout))
		}
		account := v1.Group("/account").Use(authorize)
		{
			account.GET("", core.HandlerFunc(handler.V1.Account.Get))
			account.PUT("", core.HandlerFunc(handler.V1.Account.ChangeDetail))
			account.PUT("/username", core.HandlerFunc(handler.V1.Account.ChangeUsername))
			account.POST("/push", core.HandlerFunc(handler.V1.Account.Push))
		}
		user := v1.Group("/search").Use(authorize)
		{
			user.GET("/users", core.HandlerFunc(handler.V1.Search.Users))
			user.GET("/group-chats", core.HandlerFunc(handler.V1.Search.GroupChats))
		}
		contact := v1.Group("/contacts").Use(authorize)
		{
			contact.GET("", core.HandlerFunc(handler.V1.Contact.List))
			contact.GET("/get", core.HandlerFunc(handler.V1.Contact.Get))
			contact.PUT("/remark", core.HandlerFunc(handler.V1.Contact.EditRemark))
			contact.POST("/delete", core.HandlerFunc(handler.V1.Contact.Delete))
			contact.POST("/requests", core.HandlerFunc(handler.V1.ContactRequest.Create))
			contact.GET("/requests", core.HandlerFunc(handler.V1.ContactRequest.List))
			contact.POST("/requests/accept", core.HandlerFunc(handler.V1.ContactRequest.Accept))
			contact.POST("/requests/decline", core.HandlerFunc(handler.V1.ContactRequest.Decline))
			contact.GET("/requests/unread-num", core.HandlerFunc(handler.V1.ContactRequest.ApplyUnreadNum))
			contact.GET("/folders", core.HandlerFunc(handler.V1.ContactFolder.List))
			contact.POST("/folders", core.HandlerFunc(handler.V1.ContactFolder.Save))
			contact.POST("/folders/move", core.HandlerFunc(handler.V1.ContactFolder.Move))
		}
		chat := v1.Group("/chats").Use(authorize)
		{
			chat.GET("", core.HandlerFunc(handler.V1.Chat.List))
			chat.POST("/create", core.HandlerFunc(handler.V1.Chat.Create))
			chat.POST("/delete", core.HandlerFunc(handler.V1.Chat.Delete))
			chat.POST("/topping", core.HandlerFunc(handler.V1.Chat.Top))
			chat.POST("/disturb", core.HandlerFunc(handler.V1.Chat.Disturb))
			chat.POST("/unread/clear", core.HandlerFunc(handler.V1.Chat.ClearUnreadMessage))
		}
		groupChat := v1.Group("/group-chats").Use(authorize)
		{
			groupChat.GET("", core.HandlerFunc(handler.V1.GroupChat.GroupList))
			groupChat.POST("/create", core.HandlerFunc(handler.V1.GroupChat.Create))
			groupChat.GET("/get", core.HandlerFunc(handler.V1.GroupChat.Get))
			groupChat.POST("/invite", core.HandlerFunc(handler.V1.GroupChat.Invite))
			groupChat.POST("/leave-chat", core.HandlerFunc(handler.V1.GroupChat.SignOut))
			groupChat.POST("/setting", core.HandlerFunc(handler.V1.GroupChat.Setting))
			groupChat.POST("/assign-admin", core.HandlerFunc(handler.V1.GroupChat.AssignAdmin))
			groupChat.GET("/members", core.HandlerFunc(handler.V1.GroupChat.Members))
			groupChat.GET("/members/invites", core.HandlerFunc(handler.V1.GroupChat.GetInviteFriends))
			groupChat.POST("/members/remove", core.HandlerFunc(handler.V1.GroupChat.RemoveMembers))
			groupChat.POST("/dismiss", core.HandlerFunc(handler.V1.GroupChat.Dismiss))
			groupChat.POST("/overt", core.HandlerFunc(handler.V1.GroupChat.Overt))
			groupChat.POST("/mute", core.HandlerFunc(handler.V1.GroupChat.Mute))
			groupChat.GET("/ads", core.HandlerFunc(handler.V1.GroupChatAds.List))
			groupChat.POST("/ads/edit", core.HandlerFunc(handler.V1.GroupChatAds.CreateAndUpdate))
			groupChat.POST("/ads/delete", core.HandlerFunc(handler.V1.GroupChatAds.Delete))
			groupChat.GET("/requests", core.HandlerFunc(handler.V1.GroupChatRequest.List))
			groupChat.POST("/requests/create", core.HandlerFunc(handler.V1.GroupChatRequest.Create))
			groupChat.POST("/requests/decline", core.HandlerFunc(handler.V1.GroupChatRequest.Decline))
			groupChat.POST("/requests/agree", core.HandlerFunc(handler.V1.GroupChatRequest.Agree))
			groupChat.GET("/requests/all", core.HandlerFunc(handler.V1.GroupChatRequest.All))
			groupChat.GET("/requests/unread", core.HandlerFunc(handler.V1.GroupChatRequest.RequestUnreadNum))
		}
		message := v1.Group("/messages").Use(authorize)
		{
			message.GET("", core.HandlerFunc(handler.V1.Message.GetRecords))
			message.POST("/send", core.HandlerFunc(handler.V1.MessagePublish.Publish))
			message.GET("/file/download", core.HandlerFunc(handler.V1.Message.Download))
			message.POST("/delete", core.HandlerFunc(handler.V1.Message.Delete))
			message.POST("/revoke", core.HandlerFunc(handler.V1.Message.Revoke))
			message.POST("/vote", core.HandlerFunc(handler.V1.Message.Vote))
			message.POST("/vote/handle", core.HandlerFunc(handler.V1.Message.HandleVote))
			message.GET("/stickers", core.HandlerFunc(handler.V1.Sticker.CollectList))
			message.POST("/stickers", core.HandlerFunc(handler.V1.Sticker.Upload))
			message.POST("/stickers/delete", core.HandlerFunc(handler.V1.Sticker.DeleteCollect))
			message.GET("/stickers/system/list", core.HandlerFunc(handler.V1.Sticker.SystemList))
			message.POST("/stickers/system/install", core.HandlerFunc(handler.V1.Sticker.SetSystemSticker))
			message.POST("/collect", core.HandlerFunc(handler.V1.Message.Collect))
		}
		upload := v1.Group("/upload").Use(authorize)
		{
			upload.POST("", core.HandlerFunc(handler.V1.Upload.Upload))
			upload.POST("/multipart/initiate", core.HandlerFunc(handler.V1.Upload.InitiateMultipart))
			upload.POST("/multipart", core.HandlerFunc(handler.V1.Upload.MultipartUpload))
			upload.POST("/avatar", core.HandlerFunc(handler.V1.Upload.Avatar))
		}
		bot := v1.Group("/bots").Use(authorize)
		{
			bot.POST("", core.HandlerFunc(handler.V1.Bot.Create))
			bot.GET("", core.HandlerFunc(handler.V1.Bot.List))
		}
		project := v1.Group("/projects").Use(authorize)
		{
			project.GET("", core.HandlerFunc(handler.V1.Project.Projects))
			project.POST("/create", core.HandlerFunc(handler.V1.Project.Create))
			project.GET("/members", core.HandlerFunc(handler.V1.Project.Members))
			project.POST("/invite", core.HandlerFunc(handler.V1.Project.Invite))
			project.PUT("/task-type-name", core.HandlerFunc(handler.V1.ProjectTask.TaskTypeName))
			project.GET("/tasks", core.HandlerFunc(handler.V1.ProjectTask.Tasks))
			project.POST("/tasks/create", core.HandlerFunc(handler.V1.ProjectTask.Create))
			project.GET("/tasks/detail", core.HandlerFunc(handler.V1.ProjectTask.TaskDetail))
			project.PUT("/tasks/executor", core.HandlerFunc(handler.V1.ProjectTask.Executor))
			project.PUT("/tasks/move", core.HandlerFunc(handler.V1.ProjectTask.TaskMove))
			project.GET("/tasks/coexecutors", core.HandlerFunc(handler.V1.ProjectTask.TaskCoexecutors))
			project.POST("/tasks/coexecutors/invite", core.HandlerFunc(handler.V1.ProjectTask.TaskCoexecutorInvite))
			project.GET("/tasks/watchers", core.HandlerFunc(handler.V1.ProjectTask.TaskWatchers))
			project.POST("/tasks/watchers/invite", core.HandlerFunc(handler.V1.ProjectTask.TaskWatcherInvite))
			project.GET("/tasks/comments", core.HandlerFunc(handler.V1.ProjectComment.Comments))
			project.POST("/tasks/comments/create", core.HandlerFunc(handler.V1.ProjectComment.Create))
		}
	}
}

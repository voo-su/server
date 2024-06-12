package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/http/handler"
	"voo.su/pkg/core"
	"voo.su/pkg/core/middleware"
)

func NewRouter(conf *config.Config, handler *handler.Handler, session *cache.JwtTokenStorage) *gin.Engine {
	router := gin.New()
	src, err := os.OpenFile(fmt.Sprintf("%s/http_access.log", conf.App.Log), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	router.Use(middleware.Cors(conf.Cors))
	router.Use(middleware.AccessLog(src))

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": "Ошибка системы, повторите попытку",
		})
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"code":    http.StatusOK,
			"message": "v1",
		})
	})

	authorize := middleware.Auth(conf.Jwt.Secret, "api", session)
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
			//account.PUT("/email", core.HandlerFunc(handler.V1.Account.ChangeEmail))
		}
		user := v1.Group("/users").Use(authorize)
		{
			user.GET("/search", core.HandlerFunc(handler.V1.User.Search))
		}
		contact := v1.Group("/contacts").Use(authorize)
		{
			contact.GET("", core.HandlerFunc(handler.V1.Contact.List))
			contact.GET("/get", core.HandlerFunc(handler.V1.Contact.Get))
			//contact.POST("/delete", core.HandlerFunc(handler.V1.Contact.Delete))
			//contact.POST("/edit-remark", core.HandlerFunc(handler.V1.Contact.Remark))
			contact.GET("/requests", core.HandlerFunc(handler.V1.ContactRequest.List))
			contact.POST("/requests/create", core.HandlerFunc(handler.V1.ContactRequest.Create))
			contact.POST("/requests/accept", core.HandlerFunc(handler.V1.ContactRequest.Accept))
			contact.POST("/requests/decline", core.HandlerFunc(handler.V1.ContactRequest.Decline))
			contact.GET("/requests/unread-num", core.HandlerFunc(handler.V1.ContactRequest.ApplyUnreadNum))
			contact.GET("/folders", core.HandlerFunc(handler.V1.ContactFolder.List))
			contact.POST("/folders", core.HandlerFunc(handler.V1.ContactFolder.Save))
			contact.POST("/folders/move", core.HandlerFunc(handler.V1.ContactFolder.Move))
		}
		chat := v1.Group("/chats").Use(authorize)
		{
			chat.GET("", core.HandlerFunc(handler.V1.Dialog.List))
			chat.POST("/create", core.HandlerFunc(handler.V1.Dialog.Create))
			chat.POST("/delete", core.HandlerFunc(handler.V1.Dialog.Delete))
			chat.POST("/topping", core.HandlerFunc(handler.V1.Dialog.Top))
			chat.POST("/disturb", core.HandlerFunc(handler.V1.Dialog.Disturb))
			chat.POST("/unread/clear", core.HandlerFunc(handler.V1.Dialog.ClearUnreadMessage))
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
		}
		message := v1.Group("/messages").Use(authorize)
		{
			message.GET("", core.HandlerFunc(handler.V1.Message.GetRecords))
			message.GET("/file/download", core.HandlerFunc(handler.V1.Message.Download))
			message.POST("/publish", core.HandlerFunc(handler.V1.MessagePublish.Publish))
			message.POST("/file", core.HandlerFunc(handler.V1.Message.File))
			message.POST("/delete", core.HandlerFunc(handler.V1.Message.Delete))
			message.POST("/revoke", core.HandlerFunc(handler.V1.Message.Revoke))
			message.POST("/text", core.HandlerFunc(handler.V1.Message.Text))
			message.POST("/image", core.HandlerFunc(handler.V1.Message.Image))
			message.POST("/vote", core.HandlerFunc(handler.V1.Message.Vote))
			message.POST("/vote/handle", core.HandlerFunc(handler.V1.Message.HandleVote))
			message.POST("/stickers", core.HandlerFunc(handler.V1.Message.Sticker))
			message.GET("/stickers", core.HandlerFunc(handler.V1.Sticker.CollectList))
			message.POST("/stickers/create", core.HandlerFunc(handler.V1.Sticker.Upload))
			message.POST("/stickers/delete", core.HandlerFunc(handler.V1.Sticker.DeleteCollect))
			message.GET("/stickers/system/list", core.HandlerFunc(handler.V1.Sticker.SystemList))
			message.POST("/stickers/system/install", core.HandlerFunc(handler.V1.Sticker.SetSystemSticker))
		}
		upload := v1.Group("/upload").Use(authorize)
		{
			upload.POST("/avatar", core.HandlerFunc(handler.V1.Upload.Avatar))
			upload.POST("/image", core.HandlerFunc(handler.V1.Upload.Image))
			upload.POST("/multipart/initiate", core.HandlerFunc(handler.V1.Upload.InitiateMultipart))
			upload.POST("/multipart", core.HandlerFunc(handler.V1.Upload.MultipartUpload))
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]any{
			"code":    http.StatusNotFound,
			"message": "Метод не найден",
		})
	})

	return router
}

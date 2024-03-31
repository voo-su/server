package handler

import (
	"fmt"
	"github.com/bytedance/sonic"
	"net"
	"strconv"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/entity"
	"voo.su/pkg/core/socket/adapter"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/jwt"
	"voo.su/pkg/logger"
)

type Handler struct {
	Chat   *ChatChannel
	Config *config.Config
}

type AuthConn struct {
	Uid     int    `json:"uid"`
	Channel string `json:"channel"`
	conn    *adapter.TcpAdapter
}

type Authorize struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

func (h *Handler) Dispatch(conn net.Conn) {
	ch := make(chan *AuthConn)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Восстановить", err)
		}
	}()

	fmt.Println(conn.RemoteAddr())

	go h.auth(conn, ch)

	fmt.Println(conn.RemoteAddr(), "начать аутентификацию", time.Now().Unix())

	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println(conn.RemoteAddr(), "тайм-аут аутентификации", time.Now().Unix())
		_ = conn.Close()
		return
	case info := <-ch:
		fmt.Println(conn.RemoteAddr(), "аутентификация успешна", time.Now().Unix())
		if info.Channel == entity.ImChannelChat {
			_ = h.Chat.NewClient(info.Uid, info.conn)
		}
	}
}

func (h *Handler) auth(connect net.Conn, data chan *AuthConn) {
	conn, err := adapter.NewTcpAdapter(connect)
	if err != nil {
		logger.Std().Error(fmt.Sprintf("ошибка подключения TCP: %s", err.Error()))
	}

	fmt.Println(connect.RemoteAddr(), "ожидание аутентификации", time.Now().Unix())

	read, err := conn.Read()
	if err != nil {
		fmt.Println(connect.RemoteAddr(), "исключение аутентификации", time.Now().Unix(), err.Error())
		return
	}

	if _, err := sonic.Get(read, "token"); err == nil {
		return
	}

	detail := &Authorize{}
	if err := jsonutil.Decode(read, detail); err != nil {
		return
	}

	claims, err := jwt.ParseToken(detail.Token, h.Config.Jwt.Secret)
	if err != nil || claims.Valid() != nil {
		return
	}

	uid, err := strconv.Atoi(claims.ID)
	if err != nil {
		return
	}

	data <- &AuthConn{Uid: uid, conn: conn, Channel: detail.Channel}
}

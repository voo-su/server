package socket

import (
	"context"
	"errors"
	"strconv"
	"time"
	"voo.su/pkg/timewheel"
)

const (
	heartbeatInterval = 30
	heartbeatTimeout  = 75
)

var health *heartbeat

type heartbeat struct {
	timeWheel *timewheel.SimpleTimeWheel[*Client]
}

func init() {
	health = &heartbeat{}
	health.timeWheel = timewheel.NewSimpleTimeWheel[*Client](1*time.Second, 100, health.handle)
}

func (h *heartbeat) Start(ctx context.Context) error {
	go h.timeWheel.Start()
	<-ctx.Done()
	h.timeWheel.Stop()

	return errors.New("выход из сердцебиения")
}

func (h *heartbeat) insert(c *Client) {
	h.timeWheel.Add(strconv.FormatInt(c.cid, 10), c, time.Duration(heartbeatInterval)*time.Second)
}

func (h *heartbeat) delete(c *Client) {
	h.timeWheel.Remove(strconv.FormatInt(c.cid, 10))
}

func (h *heartbeat) handle(timeWheel *timewheel.SimpleTimeWheel[*Client], key string, c *Client) {
	if c.Closed() {
		return
	}

	interval := int(time.Now().Unix() - c.lastTime)
	if interval > heartbeatTimeout {
		c.Close(2000, "Превышено время ожидания проверки сердцебиения, соединение автоматически закрыто")
		return
	}

	if interval > heartbeatInterval {
		_ = c.Write(&ClientResponse{Event: "ping"})
	}

	timeWheel.Add(key, c, time.Duration(heartbeatInterval)*time.Second)
}

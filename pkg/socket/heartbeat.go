package socket

import (
	"context"
	"errors"
	"strconv"
	"time"
	"voo.su/pkg/timeutil"
)

const (
	heartbeatInterval = 30
	heartbeatTimeout  = 75
)

var health *heartbeat

type heartbeat struct {
	TimeWheel *timeutil.SimpleTimeWheel[*Client]
}

func init() {
	health = &heartbeat{}
	health.TimeWheel = timeutil.NewSimpleTimeWheel[*Client](1*time.Second, 100, health.handle)
}

func (h *heartbeat) Start(ctx context.Context) error {
	go h.TimeWheel.Start()
	<-ctx.Done()
	h.TimeWheel.Stop()

	return errors.New("выход из сердцебиения")
}

func (h *heartbeat) insert(c *Client) {
	h.TimeWheel.Add(strconv.FormatInt(c.cid, 10), c, time.Duration(heartbeatInterval)*time.Second)
}

func (h *heartbeat) delete(c *Client) {
	h.TimeWheel.Remove(strconv.FormatInt(c.cid, 10))
}

func (h *heartbeat) handle(timeWheel *timeutil.SimpleTimeWheel[*Client], key string, c *Client) {
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

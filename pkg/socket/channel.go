package socket

import (
	"context"
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/sourcegraph/conc/pool"
	"log"
	"strconv"
	"sync/atomic"
	"time"
)

type IChannel interface {
	Name() string
	Count() int64
	Client(cid int64) (*Client, bool)
	Write(data *SenderContent)
	addClient(client *Client)
	delClient(client *Client)
}

type Channel struct {
	name    string
	count   int64
	node    cmap.ConcurrentMap[string, *Client]
	outChan chan *SenderContent
}

func NewChannel(name string, outChan chan *SenderContent) *Channel {
	return &Channel{name: name, node: cmap.New[*Client](), outChan: outChan}
}

func (c *Channel) Name() string {
	return c.name
}

func (c *Channel) Count() int64 {
	return c.count
}

func (c *Channel) Client(cid int64) (*Client, bool) {
	return c.node.Get(strconv.FormatInt(cid, 10))
}

func (c *Channel) Write(data *SenderContent) {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	select {
	case c.outChan <- data:
	case <-timer.C:
		log.Printf("%s Таймаут записи в канал OutChan, длина канала: %d \n", c.name, len(c.outChan))
	}
}

func (c *Channel) Start(ctx context.Context) error {
	var (
		worker = pool.New().WithMaxGoroutines(10)
		timer  = time.NewTicker(15 * time.Second)
	)

	defer log.Println(fmt.Errorf("Выход из канала: %s", c.Name()))

	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("Выход из канала: %s", c.Name())
		case <-timer.C:
			fmt.Printf("Канал пустой name:%s unix:%d len:%d\n", c.name, time.Now().Unix(), len(c.outChan))
		case val, ok := <-c.outChan:
			if !ok {
				return fmt.Errorf("закрытие outchan: %s", c.Name())
			}

			c.consume(worker, val, func(data *SenderContent, value *Client) {
				_ = value.Write(&ClientResponse{
					IsAck:   data.IsAck,
					Event:   data.message.Event,
					Content: data.message.Content,
					Retry:   3,
				})
			})
		}
	}
}

func (c *Channel) consume(worker *pool.Pool, data *SenderContent, fn func(data *SenderContent, value *Client)) {
	worker.Go(func() {
		if data.IsBroadcast() {
			c.node.IterCb(func(_ string, client *Client) {
				fn(data, client)
			})
			return
		}
		for _, cid := range data.receives {
			if client, ok := c.Client(cid); ok {
				fn(data, client)
			}
		}
	})
}

func (c *Channel) addClient(client *Client) {
	c.node.Set(strconv.FormatInt(client.cid, 10), client)
	atomic.AddInt64(&c.count, 1)
}

func (c *Channel) delClient(client *Client) {
	cid := strconv.FormatInt(client.cid, 10)
	if !c.node.Has(cid) {
		return
	}

	c.node.Remove(cid)
	atomic.AddInt64(&c.count, -1)
}

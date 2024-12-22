package socket

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"log"
	"sync/atomic"
	"time"
	"voo.su/pkg/strutil"
)

const (
	_MsgEventPing = "ping"
	_MsgEventPong = "pong"
	_MsgEventAck  = "ack"
)

type IClient interface {
	Cid() int64

	Uid() int

	Close(code int, text string)

	Write(data *ClientResponse) error

	Channel() IChannel
}

type IStorage interface {
	Bind(ctx context.Context, channel string, cid int64, uid int) error

	UnBind(ctx context.Context, channel string, cid int64) error
}

type Client struct {
	conn     IConn
	cid      int64
	uid      int
	lastTime int64
	closed   int32
	channel  IChannel
	storage  IStorage
	event    IEvent
	outChan  chan *ClientResponse
}

type ClientOption struct {
	Uid         int
	Channel     IChannel
	Storage     IStorage
	IdGenerator IIdGenerator
	Buffer      int
}

type ClientResponse struct {
	IsAck   bool   `json:"-"`
	Sid     string `json:"sid,omitempty"`
	Event   string `json:"event"`
	Content any    `json:"content,omitempty"`
	Retry   int    `json:"-"`
}

func NewClient(conn IConn, option *ClientOption, event IEvent) error {
	if option.Buffer <= 0 {
		option.Buffer = 10
	}
	if event == nil {
		panic("empty")
	}

	client := &Client{
		conn:     conn,
		uid:      option.Uid,
		lastTime: time.Now().Unix(),
		channel:  option.Channel,
		storage:  option.Storage,
		outChan:  make(chan *ClientResponse, option.Buffer),
		event:    event,
	}
	if option.IdGenerator != nil {
		client.cid = option.IdGenerator.IdGen()
	} else {
		client.cid = defaultIdGenerator.IdGen()
	}

	conn.SetCloseHandler(client.hookClose)
	if client.storage != nil {
		err := client.storage.Bind(context.TODO(), client.channel.Name(), client.cid, client.uid)
		if err != nil {
			log.Printf("Client binding error: %s", err)
			return err
		}
	}

	client.channel.addClient(client)
	client.event.Open(client)
	health.insert(client)

	return client.init()
}

func (c *Client) Channel() IChannel {
	return c.channel
}

func (c *Client) Cid() int64 {
	return c.cid
}

func (c *Client) Uid() int {
	return c.uid
}

func (c *Client) Close(code int, message string) {
	defer func() {
		if err := c.conn.Close(); err != nil {
			log.Printf("error closing connection: %s \n", err.Error())
		}
	}()
	if err := c.hookClose(code, message); err != nil {
		log.Printf("%s-%d-%d client close error: %s", c.channel.Name(), c.cid, c.uid, err)
	}
}

func (c *Client) Closed() bool {
	return atomic.LoadInt32(&c.closed) == 1
}

func (c *Client) Write(data *ClientResponse) error {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s-%d-%d error writing to pipe: %v", c.channel.Name(), c.cid, c.uid, err)
		}
	}()
	if c.Closed() {
		return fmt.Errorf("connection closed")
	}

	if data.IsAck {
		data.Sid = strutil.NewMsgId()
	}

	c.outChan <- data

	return nil
}

func (c *Client) loopAccept() {
	defer c.Close(1000, "cycle reception closed")
	for {
		data, err := c.conn.Read()
		if err != nil {
			break
		}
		c.lastTime = time.Now().Unix()
		c.handleMessage(data)
	}
}

func (c *Client) loopWrite() {
	timer := time.NewTimer(15 * time.Second)
	defer timer.Stop()
	for {
		timer.Reset(15 * time.Second)
		select {
		case <-timer.C:
			log.Printf("Client cid:%d uid:%d time:%d", c.cid, c.uid, time.Now().Unix())
		case data, ok := <-c.outChan:
			if !ok || c.Closed() {
				return
			}

			bt, err := sonic.Marshal(data)
			if err != nil {
				log.Printf("loopWrite json decode err: %v \n", err)
				break
			}

			if err := c.conn.Write(bt); err != nil {
				log.Printf("%s-%d-%d client write error: %v \n", c.channel.Name(), c.cid, c.uid, err)
				return
			}

			if data.IsAck && data.Retry > 0 {
				data.Retry--
				ackBufferContent := &AckBufferContent{}
				ackBufferContent.Cid = c.cid
				ackBufferContent.Uid = int64(c.uid)
				ackBufferContent.Channel = c.channel.Name()
				ackBufferContent.Response = data
				ack.insert(data.Sid, ackBufferContent)
			}
		}
	}
}

func (c *Client) init() error {
	_ = c.Write(&ClientResponse{
		Event: "connect",
		Content: map[string]any{
			"ping_interval": heartbeatInterval,
			"ping_timeout":  heartbeatTimeout,
		}},
	)
	go c.loopWrite()
	go c.loopAccept()
	return nil
}

func (c *Client) hookClose(code int, text string) error {
	if !atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		return nil
	}

	close(c.outChan)
	c.event.Close(c, code, text)
	if c.storage != nil {
		err := c.storage.UnBind(context.TODO(), c.channel.Name(), c.cid)
		if err != nil {
			log.Printf("Error unbinding client: %s", err)
			return err
		}
	}
	health.delete(c)
	c.channel.delClient(c)
	return nil
}

func (c *Client) handleMessage(data []byte) {
	event, err := c.validate(data)
	if err != nil {
		log.Printf("Verification error: %s \n", err.Error())

		return
	}
	switch event {
	case _MsgEventPing:
		_ = c.Write(&ClientResponse{Event: _MsgEventPong})
	case _MsgEventPong:
	case _MsgEventAck:
		val, err := sonic.Get(data, "sid")
		if err == nil {
			ackId, _ := val.String()
			if len(ackId) > 0 {
				ack.delete(ackId)
			}
		}
	default:
		c.event.Message(c, data)
	}
}

func (c *Client) validate(data []byte) (string, error) {
	if !sonic.Valid(data) {
		return "", fmt.Errorf("validate err")
	}

	value, err := sonic.Get(data, "event")
	if err != nil {
		return "", err
	}

	return value.String()
}

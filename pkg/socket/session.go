package socket

import (
	"context"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

var Session *session
var once sync.Once

type session struct {
	Chat     *Channel
	Channels map[string]*Channel
}

func (s *session) Channel(name string) (*Channel, bool) {
	val, ok := s.Channels[name]

	return val, ok
}

func Initialize(ctx context.Context, eg *errgroup.Group, fn func(name string)) {
	once.Do(func() {
		InitAck()
		initialize(ctx, eg, fn)
	})
}

func initialize(ctx context.Context, eg *errgroup.Group, fn func(name string)) {
	Session = &session{
		Chat:     NewChannel("chat", make(chan *SenderContent, 5<<20)),
		Channels: map[string]*Channel{},
	}

	Session.Channels["chat"] = Session.Chat

	time.AfterFunc(3*time.Second, func() {
		eg.Go(func() error {
			defer fn("health exit")
			return health.Start(ctx)
		})

		eg.Go(func() error {
			defer fn("ack exit")
			return ack.Start(ctx)
		})

		eg.Go(func() error {
			defer fn("chat exit")
			return Session.Chat.Start(ctx)
		})
	})
}

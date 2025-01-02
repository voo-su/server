// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package timewheel

import (
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/sourcegraph/conc/pool"
	"time"
)

type entry[T any] struct {
	Key    string
	Value  T
	Expire int64
}

type SimpleTimeWheel[T any] struct {
	interval  time.Duration
	ticker    *time.Ticker
	tickIndex int
	slot      []cmap.ConcurrentMap[string, *entry[T]]
	indicator cmap.ConcurrentMap[string, int]
	onTick    SimpleHandler[T]
	taskChan  chan *entry[T]
	quitChan  chan struct{}
}

type SimpleHandler[T any] func(*SimpleTimeWheel[T], string, T)

func NewSimpleTimeWheel[T any](delay time.Duration, numSlot int, handler SimpleHandler[T]) *SimpleTimeWheel[T] {
	timeWheel := &SimpleTimeWheel[T]{
		taskChan:  make(chan *entry[T], 100),
		quitChan:  make(chan struct{}),
		indicator: cmap.New[int](),
		interval:  delay,
		ticker:    time.NewTicker(delay),
		onTick:    handler,
	}
	for i := 0; i < numSlot; i++ {
		timeWheel.slot = append(timeWheel.slot, cmap.New[*entry[T]]())
	}

	return timeWheel
}

func (s *SimpleTimeWheel[T]) Start() {
	go s.run()
	for {
		select {
		case <-s.quitChan:
			return
		case el := <-s.taskChan:
			s.Remove(el.Key)
			slotIndex := s.getCircleAndSlot(el)
			s.slot[slotIndex].Set(el.Key, el)
			s.indicator.Set(el.Key, slotIndex)
		}
	}
}

func (s *SimpleTimeWheel[T]) Stop() {
	close(s.quitChan)
}

func (s *SimpleTimeWheel[T]) run() {
	worker := pool.New().WithMaxGoroutines(10)
	for {
		select {
		case <-s.quitChan:
			s.ticker.Stop()

			return
		case <-s.ticker.C:
			tickIndex := s.tickIndex
			s.tickIndex++
			if s.tickIndex >= len(s.slot) {
				s.tickIndex = 0
			}

			slot := s.slot[tickIndex]
			for item := range slot.IterBuffered() {
				v := item.Val
				slot.Remove(v.Key)
				s.indicator.Remove(v.Key)
				worker.Go(func() {
					unix := time.Now().Unix()
					if v.Expire <= unix {
						s.onTick(s, v.Key, v.Value)
					} else {
						s.Add(v.Key, v.Value, time.Duration(v.Expire-unix)*time.Second)
					}
				})
			}
		}
	}
}

func (s *SimpleTimeWheel[T]) Add(key string, value T, delay time.Duration) {
	s.taskChan <- &entry[T]{Key: key, Value: value, Expire: time.Now().Add(delay).Unix()}
}

func (s *SimpleTimeWheel[T]) Remove(key string) {
	if value, ok := s.indicator.Get(key); ok {
		s.slot[value].Remove(key)
		s.indicator.Remove(key)
	}
}

func (s *SimpleTimeWheel[T]) getCircleAndSlot(e *entry[T]) int {
	remainingTime := int(e.Expire - time.Now().Unix())
	if remainingTime <= 0 {
		remainingTime = 0
	}

	return (s.tickIndex + remainingTime/int(s.interval.Seconds())) % len(s.slot)
}

package timewheel

import (
	"errors"
	"fmt"
	"github.com/sourcegraph/conc/pool"
	"log"
	"sync"
	"time"
)

type element struct {
	key    string
	value  any
	expire int64
}

type slot struct {
	id       int
	elements *sync.Map
}

func newSlot(id int) *slot {
	s := &slot{id: id}
	s.elements = &sync.Map{}

	return s
}

func (s *slot) add(el *element) {
	s.elements.LoadOrStore(el.key, el)
}

func (s *slot) remove(key any) {
	s.elements.Delete(key)
}

type circle struct {
	index     int
	tickIndex int
	ticker    *time.Ticker
	slot      []*slot
}

func newCircle(index int, numSlots int, ticker *time.Ticker) *circle {
	c := &circle{
		index:     index,
		tickIndex: 0,
		ticker:    ticker,
		slot:      make([]*slot, 0, numSlots),
	}
	for i := 0; i < numSlots; i++ {
		c.slot = append(c.slot, newSlot(i))
	}

	return c
}

type Handler func(*TimeWheel, any)

type TimeWheel struct {
	circle    []*circle
	onTick    Handler
	taskChan  chan any
	quitChan  chan any
	indicator *sync.Map
}

func NewTimeWheel(handler Handler) *TimeWheel {
	timeWheel := &TimeWheel{
		taskChan:  make(chan any, 100),
		quitChan:  make(chan any),
		indicator: &sync.Map{},
		onTick:    handler,
	}
	timeWheel.circle = []*circle{
		newCircle(0, 60, time.NewTicker(time.Second)),
		newCircle(1, 60, time.NewTicker(time.Minute)),
		newCircle(2, 24, time.NewTicker(time.Hour)),
	}

	return timeWheel
}

func (t *TimeWheel) Start() {
	defer fmt.Println("TimeWheel Stop")
	for _, c := range t.circle {
		go func(c *circle) {
			t.runTimeWheel(c)
		}(c)
	}
	for {
		select {
		case <-t.quitChan:
			return
		case v := <-t.taskChan:
			el, ok := v.(*element)
			if !ok {
				continue
			}
			circleIndex, slotIndex := t.getCircleAndSlot(el)
			circleSlot := t.circle[circleIndex].slot[slotIndex]
			circleSlot.add(el)
			t.indicator.Store(el.value, circleSlot)
		}
	}
}

func (t *TimeWheel) getCircleAndSlot(el *element) (int, int64) {
	var (
		circleIndex   int
		slotIndex     int64
		remainingTime = int(el.expire - time.Now().Unix())
	)
	if remainingTime <= 0 {
		remainingTime = 0
	}

	if remainingTime < 60 {
		circleIndex = 0
		slotIndex = int64((t.getCurrentTickIndex(0) + remainingTime) % 60)
	} else if int(remainingTime/60) < 60 {
		circleIndex = 1
		slotIndex = int64((t.getCurrentTickIndex(1) + remainingTime/60) % 60)
	} else {
		circleIndex = 2
		slotIndex = int64((t.getCurrentTickIndex(1) + remainingTime/3600) % 60)
	}

	return circleIndex, slotIndex
}

func (t *TimeWheel) runTimeWheel(circle *circle) {
	defer fmt.Printf("[%d]RunTimeWheel Stop\n", circle.index)
	worker := pool.New().WithMaxGoroutines(10)
	for {
		select {
		case <-t.quitChan:
			circle.ticker.Stop()
			return
		case <-circle.ticker.C:
			tickIndex := circle.tickIndex
			circle.tickIndex++
			if circle.tickIndex >= len(circle.slot) {
				circle.tickIndex = 0
			}
			circleSlot := circle.slot[tickIndex]
			circleSlot.elements.Range(func(_, value any) bool {
				if el, ok := value.(*element); ok {
					t.indicator.Delete(el.value)
					circleSlot.remove(el.value)
					worker.Go(func() {
						if el.expire <= time.Now().Unix() {
							t.onTick(t, el.value)
						} else {
							second := el.expire - time.Now().Unix()
							if err := t.Add(el.value, time.Duration(second)*time.Second); err != nil {
								log.Printf("Ошибка при понижении уровня временного колеса: %s", err.Error())
							}
						}
					})
				}
				return true
			})
		}
	}
}

func (t *TimeWheel) Stop() {
	close(t.quitChan)
}

func (t *TimeWheel) Add(task any, delay time.Duration) error {
	if delay > 24*time.Hour {
		return errors.New("максимальная задержка 24 часа")
	}
	t.taskChan <- &element{value: task, expire: time.Now().Add(delay).Unix()}

	return nil
}

func (t *TimeWheel) Remove(task any) {
	if value, ok := t.indicator.Load(task); ok {
		if slot, ok := value.(*slot); ok {
			slot.remove(task)
			t.indicator.Delete(task)
		}
	}
}

func (t *TimeWheel) getCurrentTickIndex(circleIndex int) int {
	return t.circle[circleIndex].tickIndex
}

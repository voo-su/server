package socket

type Message struct {
	Event   string `json:"event"`
	Content any    `json:"content"`
}

func NewMessage(event string, content any) *Message {
	return &Message{
		Event:   event,
		Content: content,
	}
}

type SenderContent struct {
	IsAck     bool
	broadcast bool
	exclude   []int64
	receives  []int64
	message   *Message
}

func NewSenderContent() *SenderContent {
	return &SenderContent{
		broadcast: false,
		exclude:   make([]int64, 0),
		receives:  make([]int64, 0),
	}
}

func (s *SenderContent) SetAck(value bool) *SenderContent {
	s.IsAck = value

	return s
}

func (s *SenderContent) SetBroadcast(value bool) *SenderContent {
	s.broadcast = value

	return s
}

func (s *SenderContent) SetMessage(event string, content any) *SenderContent {
	s.message = NewMessage(event, content)

	return s
}

func (s *SenderContent) SetReceive(cid ...int64) *SenderContent {
	s.receives = append(s.receives, cid...)

	return s
}

func (s *SenderContent) SetExclude(cid ...int64) *SenderContent {
	s.exclude = append(s.exclude, cid...)

	return s
}

func (s *SenderContent) IsBroadcast() bool {
	return s.broadcast
}

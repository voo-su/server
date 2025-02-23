package push

type RTCCallType int

type Payload interface {
	// GetTitle Заголовок уведомления
	GetTitle() string

	// GetContent Текст уведомления
	GetContent() string

	// GetBadge Количество уведомлений
	GetBadge() int

	// GetRTCPayload Получить RTC нагрузку
	GetRTCPayload() RTCPayload
}

type RTCPayload interface {
	// GetCallType Тип видеозвонка
	GetCallType() RTCCallType

	// GetOperation Операция видеозвонка
	// invite: пригласить
	// cancel: отменить приглашение
	GetOperation() string

	// GetFromUID UID инициатора
	GetFromUID() string
}

type BasePayload struct {
	title   string
	content string
	badge   int
}

func (p *BasePayload) GetTitle() string {
	return p.title
}

func (p *BasePayload) GetContent() string {
	return p.content
}

func (p *BasePayload) GetBadge() int {
	return p.badge
}

func (p *BasePayload) GetRTCPayload() RTCPayload {
	return nil
}

type BaseRTCPayload struct {
	BasePayload
	callType  RTCCallType
	operation string
	fromUID   string
}

func (b *BaseRTCPayload) GetCallType() RTCCallType {
	return b.callType
}

func (b *BaseRTCPayload) GetOperation() string {
	return b.operation
}

func (b *BaseRTCPayload) GetFromUID() string {
	return b.fromUID
}

func (b *BaseRTCPayload) GetRTCPayload() RTCPayload {
	return b
}

type Push interface {
	GetPayload() (Payload, error)

	Push(deviceToken string, payload Payload) error
}

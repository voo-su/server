package push

type PayloadInfo struct {
	Title   string
	Content string
	Badge   int

	// RTC
	IsVideoCall bool   // Является ли это RTC-сообщением
	FromUID     string // Требуется для RTC-сообщений
	CallType    RTCCallType
	Operation   string
}

func (p *PayloadInfo) toPayload() Payload {
	var basePayload = BasePayload{
		title:   p.Title,
		content: p.Content,
		badge:   p.Badge,
	}
	var payload Payload

	if p.IsVideoCall {
		payload = &BaseRTCPayload{
			BasePayload: basePayload,
			fromUID:     p.FromUID,
			operation:   p.Operation,
			callType:    p.CallType,
		}
	} else {
		payload = &basePayload
	}

	return payload
}

func ParsePushInfo() (*PayloadInfo, error) {
	payloadInfo := &PayloadInfo{
		Badge:   1,
		Title:   "titleTitle",
		Content: "contentContent",
	}

	return payloadInfo, nil
}

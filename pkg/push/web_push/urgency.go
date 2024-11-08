package web_push

// Urgency указывает сервису push-уведомлений, насколько важно сообщение для пользователя.
type Urgency string

const (
	// UrgencyVeryLow требует состояния устройства: подключено к питанию и Wi-Fi
	UrgencyVeryLow Urgency = "very-low"
	// UrgencyLow требует состояния устройства: подключено либо к питанию, либо к Wi-Fi
	UrgencyLow Urgency = "low"
	// UrgencyNormal исключает состояние устройства: низкий заряд батареи
	UrgencyNormal Urgency = "normal"
	// UrgencyHigh допускает состояние устройства: низкий заряд батареи
	UrgencyHigh Urgency = "high"
)

// Проверка допустимых значений заголовка urgency
func isValidUrgency(urgency Urgency) bool {
	switch urgency {
	case UrgencyVeryLow, UrgencyLow, UrgencyNormal, UrgencyHigh:
		return true
	}

	return false
}

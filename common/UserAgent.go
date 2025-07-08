package common

// UserAgent - структура для хранения User-Agent
type UserAgent struct {
	Data string
}

// ParseUserAgentFromString - парсит User-Agent из строки
func ParseUserAgentFromString(data string) UserAgent {
	return UserAgent{
		Data: data,
	}
}

// ToString - возвращает строковое представление User-Agent
func (ua UserAgent) ToString() string {
	return ua.Data
}

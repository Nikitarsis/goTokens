package common

// UserAgentData - структура для хранения User-Agent
type UserAgentData struct {
	Data string
}

// ParseUserAgentFromString - парсит User-Agent из строки
func ParseUserAgentFromString(data string) UserAgentData {
	return UserAgentData{
		Data: data,
	}
}

// ToString - возвращает строковое представление User-Agent
func (ua UserAgentData) ToString() string {
	return ua.Data
}

package common

// IUserAgentRepository - интерфейс для работы с репозиторием User-Agent
type IUserAgentRepository interface {
	SaveUserAgent(kid UUID, userAgent UserAgentData) error
	CheckUserAgent(kid UUID, userAgent UserAgentData) bool
}

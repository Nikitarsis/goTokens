package common

// IUserAgentRepository - интерфейс для работы с репозиторием User-Agent
type IUserAgentRepository interface {
	SaveUserAgent(kid UUID, userAgent string) error
	CheckUserAgent(kid UUID, userAgent string) bool
}
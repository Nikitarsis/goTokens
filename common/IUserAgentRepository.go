package common

type IUserAgentRepository interface {
	SaveUserAgent(kid UUID, userAgent string) error
	CheckUserAgent(kid UUID) bool
}
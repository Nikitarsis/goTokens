package inmemory

import (
	co "github.com/Nikitarsis/goTokens/common"
)

// In-Memory репозиторий User-Agent
type UserAgentRepositoryInMemory struct {
	userAgents *co.SafeMap[co.UUID, string]
}

// CreateInMemoryUserRepository создает новый экземпляр UserAgentRepositoryInMemory.
func CreateInMemoryUserRepository() co.IUserAgentRepository {
	safeMap := co.CreateSafeMap[co.UUID, string]()
	return &UserAgentRepositoryInMemory{
		userAgents: safeMap,
	}
}

// SaveUserAgent сохраняет User-Agent в репозитории.
func (r *UserAgentRepositoryInMemory) SaveUserAgent(kid co.UUID, userAgent string) error {
	r.userAgents.Store(kid, userAgent)
	return nil
}

// CheckUserAgent проверяет User-Agent в репозитории.
func (r *UserAgentRepositoryInMemory) CheckUserAgent(kid co.UUID, userAgent string) bool {
	ua, ok := r.userAgents.Load(kid)
	if !ok {
		return false
	}
	return ua == userAgent
}
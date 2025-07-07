package inmemory

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type UserAgentRepositoryInMemory struct {
	userAgents *co.SafeMap[co.UUID, string]
}

func NewUserAgentRepositoryInMemory() co.IUserAgentRepository {
	safeMap := co.CreateSafeMap[co.UUID, string]()
	return &UserAgentRepositoryInMemory{
		userAgents: safeMap,
	}
}

func (r *UserAgentRepositoryInMemory) SaveUserAgent(kid co.UUID, userAgent string) error {
	r.userAgents.Store(kid, userAgent)
	return nil
}

func (r *UserAgentRepositoryInMemory) CheckUserAgent(kid co.UUID, userAgent string) bool {
	ua, ok := r.userAgents.Load(kid)
	if !ok {
		return false
	}
	return ua == userAgent
}
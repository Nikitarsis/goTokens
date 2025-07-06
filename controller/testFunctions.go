package controller

import (
	"sync"

	co "github.com/Nikitarsis/goTokens/common"
)

type testUserAgentRepository struct {
	ret *sync.Map
}

func (r *testUserAgentRepository) SaveUserAgent(kid co.UUID, userAgent string) error {
	r.ret.Store("userAgent", userAgent)
	r.ret.Store("kid-UA-save", kid.ToString())
	return nil
}

func (r *testUserAgentRepository) CheckUserAgent(kid co.UUID) bool {
	r.ret.Store("kid-UA-check", kid.ToString())
	return true
}

func getTestUserAgentRepository() (co.IUserAgentRepository, *sync.Map) {
	ret := &sync.Map{}
	repo := testUserAgentRepository{ret: ret}
	return &repo, ret
}

type testIpRepository struct {
	ret *sync.Map
}

func (r *testIpRepository) TraceIp(kid co.UUID, ip string) {
	r.ret.Store("kid-IP-trace", kid.ToString())
	r.ret.Store("ip", ip)
}

func getTestIpRepository() (co.IIpRepository, *sync.Map) {
	ret := &sync.Map{}
	repo := testIpRepository{ret: ret}
	return &repo, ret
}

func getTestTokenPairGetter(
	tokenPairFunc func(co.UUID) (map[string]co.TokenData, error),
) (*TokensPairGetter, func(string) string) {
	var mapUARepository *sync.Map
	var mapIPRepository *sync.Map
	checker := func(key string) string {
		if ret, ok := mapIPRepository.Load(key); ok {
			return ret.(string)
		}
	if ret, ok := mapUARepository.Load(key); ok {
		return ret.(string)
	}
	return ""
}
UARepository, mapUARepository := getTestUserAgentRepository()
IPRepository, mapIPRepository := getTestIpRepository()

return NewTokensPairGetter(
		tokenPairFunc,
		UARepository,
		IPRepository,
	), checker
}
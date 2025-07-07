package controller

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
)

type testUserAgentRepository struct {
	ret *co.SafeMap[string, string]
}

func (r *testUserAgentRepository) SaveUserAgent(kid co.UUID, userAgent string) error {
	r.ret.Store("userAgent", userAgent)
	r.ret.Store("kid-UA-save", kid.ToString())
	return nil
}

func (r *testUserAgentRepository) CheckUserAgent(kid co.UUID, userAgent string) bool {
	r.ret.Store("kid-UA-check", kid.ToString())
	return true
}

func getTestUserAgentRepository() (co.IUserAgentRepository, *co.SafeMap[string, string]) {
	ret := co.CreateSafeMap[string, string]()
	repo := testUserAgentRepository{ret: ret}
	return &repo, ret
}

type testIpRepository struct {
	ret *co.SafeMap[string, string]
}

func (r *testIpRepository) TraceIp(data co.DataIP) {
	var ip string
	if data.IP == nil {
		ip = ""
	} else {
		ip = data.IP.String()
	}
	r.ret.Store("kid-IP-trace", data.KeyId.ToString())
	r.ret.Store("ip", ip)
	r.ret.Store("port", fmt.Sprintf("%d", data.Port))
	r.ret.Store("uid", data.UserId.ToString())
}

func getTestIpRepository() (co.IIpRepository, *co.SafeMap[string, string]) {
	ret := co.CreateSafeMap[string, string]()
	repo := testIpRepository{ret: ret}
	return &repo, ret
}

func getTestTokenPairGetter(
	tokenPairFunc func(co.UUID) (map[string]co.TokenData, error),
) (*TokensPairGetter, func(string) string) {
	var mapUARepository *co.SafeMap[string, string]
	var mapIPRepository *co.SafeMap[string, string]
	checker := func(key string) string {
		if ret, ok := mapIPRepository.Load(key); ok {
			return ret
		}
		if ret, ok := mapUARepository.Load(key); ok {
			return ret
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

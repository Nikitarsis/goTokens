package controller

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
)
//Тестовая имплементация IUserAgentRepository
type testUserAgentRepository struct {
	ret *co.SafeMap[string, string]
}

// SaveUserAgent сохраняет userAgent для данного kid:
// userAgent -> userAgent(SaveUserAgent);
// kid -> kid(SaveUserAgent);
func (r *testUserAgentRepository) SaveUserAgent(kid co.UUID, userAgent co.UserAgentData) error {
	r.ret.Store("userAgent(SaveUserAgent)", userAgent.ToString())
	r.ret.Store("kid(SaveUserAgent)", kid.ToString())
	return nil
}

// CheckUserAgent проверяет userAgent для данного kid:
// userAgent -> userAgent(CheckUserAgent);
// kid -> kid(CheckUserAgent);
func (r *testUserAgentRepository) CheckUserAgent(kid co.UUID, userAgent co.UserAgentData) bool {
	r.ret.Store("userAgent(CheckUserAgent)", userAgent.ToString())
	r.ret.Store("kid(CheckUserAgent)", kid.ToString())
	return true
}
 
// Создание тестового IUserAgentRepository
func getTestUserAgentRepository() (co.IUserAgentRepository, *co.SafeMap[string, string]) {
	ret := co.CreateSafeMap[string, string]()
	repo := testUserAgentRepository{ret: ret}
	return &repo, ret
}

// Тестовая имплементация IIpRepository
type testIpRepository struct {
	ret *co.SafeMap[string, string]
}

// TraceIp отслеживает IP-адрес для данного kid.
// DataIp хранится как несколько ключей:
// DataIp.KeyId -> kid(TraceIp);
// DataIp.IP -> ip(TraceIp);
// DataIp.Port -> port(TraceIp);
// DataIp.UserId -> uid(TraceIp);
func (r *testIpRepository) TraceIp(data co.DataIP) {
	var ip string
	if data.IP == nil {
		ip = ""
	} else {
		ip = data.IP.String()
	}
	r.ret.Store("kid(TraceIp)", data.KeyId.ToString())
	r.ret.Store("ip(TraceIp)", ip)
	r.ret.Store("port(TraceIp)", fmt.Sprintf("%d", data.Port))
	r.ret.Store("uid(TraceIp)", data.UserId.ToString())
}

// Создание тестового IIpTracer
func getTestIpRepository() (co.IIpTracer, *co.SafeMap[string, string]) {
	ret := co.CreateSafeMap[string, string]()
	repo := testIpRepository{ret: ret}
	return &repo, ret
}

func getTestTokenPairGetter(
	tokenPairFunc func(co.UUID) (map[string]co.TokenData, error),
) (*TokensPairGetter, func(int, string) string) {
	var mapUARepository *co.SafeMap[string, string]
	var mapIPRepository *co.SafeMap[string, string]
	checker := func(code int, key string) string {
		if code == 0 {
			ret, _ := mapIPRepository.Load(key)
			return ret
		}
		if code == 1 {
			ret, _ := mapUARepository.Load(key)
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
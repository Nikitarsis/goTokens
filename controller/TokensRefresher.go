package controller

import (
	"encoding/json"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

// TokenRefresher - это класс для обновления токенов
type TokensRefresher struct {
	getPairTokens func(co.UUID) (map[string]co.TokenData, error)
	parseToken    func(co.Token) (co.TokenData, error)
	userAgent     co.IUserAgentRepository
	ipTracer      co.IIpRepository
	dropKey       func(co.UUID) bool
}

// NewTokensRefresher - это конструктор для TokensRefresher
func NewTokensRefresher(
	getPairTokens func(co.UUID) (map[string]co.TokenData, error),
	parseToken func(co.Token) (co.TokenData, error),
	userAgent co.IUserAgentRepository,
	ipTracer co.IIpRepository,
	dropKey func(co.UUID) bool,
	) *TokensRefresher {
	return &TokensRefresher{
		getPairTokens: getPairTokens,
		parseToken:    parseToken,
		userAgent:     userAgent,
		ipTracer:      ipTracer,
		dropKey:       dropKey,
	}
}

// RefreshTokens - это метод для обновления токенов
func (tr TokensRefresher) RefreshTokens(request *http.Request) (co.Response) {
	// Парсинг тела запроса
	token, uid, err := parseBodyWithId(request)
	if err != nil {
		return co.ParseError(err)
	}
	// Парсинг токена
	// Токен должен возвращать ошибку при !Valid
	parsedToken, err := tr.parseToken(token)
	if err != nil {
		return co.ParseError(co.ErrInvalidToken)
	}
	// Проверка типа токена. Должен быть RefreshToken
	if parsedToken.Type != co.RefreshToken {
		return co.ParseError(co.ErrWrongToken)
	}
	// Проверка соответствия UID
	if (parsedToken.UserId != uid) {
		// Удаление ключа как скомпроментированного
		go tr.dropKey(parsedToken.KeyId)
		return co.ParseError(co.ErrStealedToken)
	}
	// Если User-Agent не соответствует указанному при получении токена, ключи удаляются
	if !tr.userAgent.CheckUserAgent(uid, request.UserAgent()) {
		go tr.dropKey(parsedToken.KeyId)
		return co.ParseError(co.ErrInvalidUserAgent)
	}
	// Получение пары токенов
	pairRaw, err := tr.getPairTokens(uid)
	if err != nil {
		return co.ParseError(co.ErrInternalServerError)
	}
	pair := TokensPair{
		Access:  pairRaw["access"].Token.ToString(),
		Refresh: pairRaw["refresh"].Token.ToString(),
	}
	// Превращение пары токенов в JSON
	ret, err := json.Marshal(pair)
	if err != nil {
		return co.ParseError(co.ErrInternalServerError)
	}
	// Удаление ключа как устаревшего
	go tr.dropKey(parsedToken.KeyId)
	return co.Response{
		StatusCode: http.StatusOK,
		Message:    ret,
	}
}

func (tr TokensRefresher) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := tr.RefreshTokens(r)
		w.WriteHeader(response.StatusCode)
		w.Write(response.Message)
	})
}
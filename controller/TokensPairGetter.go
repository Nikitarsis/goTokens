package controller

import (
	"encoding/json"
	"io"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

// TokensPairGetter - структура для получения пары токенов
type TokensPairGetter struct {
	getPairTokens       func(co.UUID) (map[string]co.TokenData, error)
	userAgentRepository co.IUserAgentRepository
	ipRepository        co.IIpTracer
}

// NewTokensPairGetter - создает новый экземпляр TokensPairGetter
func NewTokensPairGetter(
	getPairTokens func(co.UUID) (map[string]co.TokenData, error),
	userAgentRepository co.IUserAgentRepository,
	ipRepository co.IIpTracer,
) *TokensPairGetter {
	return &TokensPairGetter{
		getPairTokens:       getPairTokens,
		userAgentRepository: userAgentRepository,
		ipRepository:        ipRepository,
	}
}

// parseRequestGet - парсит GET-запрос и извлекает UID из строки параметров
func (tg TokensPairGetter) parseRequestGet(request *http.Request) (co.UUID, error) {
	userIdStr := request.URL.Query().Get("uid")
	if userIdStr == "" {
		return co.UUID{}, co.ErrNoUserId
	}
	userId, err := co.GetUUIDFromString(userIdStr)
	if err != nil {
		return co.UUID{}, co.ErrInvalidUserId
	}
	return userId, nil
}

// parseRequestPost - парсит POST-запрос и извлекает UID из тела запроса
func (tg TokensPairGetter) parseRequestPost(request *http.Request) (co.UUID, error) {
	var userId UserId
	var body []byte
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return co.UUID{}, co.ErrJsonParsingError
	}
	err = json.Unmarshal(body, &userId)
	if err != nil {
		return co.UUID{}, co.ErrInvalidUserId
	}
	ret, err := co.GetUUIDFromString(userId.ID)
	if err != nil {
		return co.UUID{}, co.ErrInvalidUserId
	}
	return ret, nil
}

// GetTokensPair - получает пару токенов
func (tg TokensPairGetter) GetTokensPair(request *http.Request) co.Response {
	method := request.Method
	var userId co.UUID
	var err error
	// Разбор метода
	switch method {
	case http.MethodGet:
		userId, err = tg.parseRequestGet(request)
		if err != nil {
			return co.ParseError(err)
		}
	case http.MethodPost:
		userId, err = tg.parseRequestPost(request)
		if err != nil {
			return co.ParseError(err)
		}
	default:
		return co.ParseError(co.ErrInvalidMethod)
	}
	// Получение пары токенов
	result, err := tg.getPairTokens(userId)
	if err != nil {
		return co.ParseError(co.ErrInternalServerError)
	}
	pair := TokensPair{
		Access:  result["access"].Token.ToString(),
		Refresh: result["refresh"].Token.ToString(),
	}
	refresh := result["refresh"]
	ret, err := json.Marshal(pair)
	if err != nil {
		return co.ParseError(co.ErrInternalServerError)
	}
	// Сохранение userAgent
	go tg.userAgentRepository.SaveUserAgent(
		refresh.KeyId,
		co.ParseUserAgentFromString(request.UserAgent()),
	)
	// Трассировка IP
	errIp := traceIp(refresh, request.RemoteAddr, tg.ipRepository)
	if errIp != nil {
		return co.ParseError(errIp)
	}
	return co.Response{
		StatusCode: http.StatusOK,
		Message:    ret,
	}
}

// GetHandler - возвращает обработчик для получения пары токенов
//
// # HTTP-методы - GET, POST
//
// Возвращает ошибку, если токен невалиден или возникла ошибка при обработке тела или токена
func (tg TokensPairGetter) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := tg.GetTokensPair(r)
		w.WriteHeader(response.StatusCode)
		w.Write(response.Message)
	})
}

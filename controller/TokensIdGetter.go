package controller

import (
	"encoding/json"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

// TokensIdGetter - структура для получения user Id из токена
type TokensIdGetter struct {
	parseToken func(co.Token) (co.TokenData, error)
}

// NewTokensIdGetter - создает новый экземпляр TokensIdGetter
func NewTokensIdGetter(parseToken func(co.Token) (co.TokenData, error)) *TokensIdGetter {
	return &TokensIdGetter{
		parseToken: parseToken,
	}
}

// GetTokenId - получает user Id из токена
func (tid *TokensIdGetter) GetTokenId(request *http.Request) co.Response {
	//Проверка метода, должен быть POST
	if request.Method != http.MethodPost {
		return co.ParseError(co.ErrInvalidMethod)
	}
	token, err := parseBody(request)
	if err != nil {
		return co.ParseError(err)
	}
	// Парсинг токена
	// Токен должен возвращать ошибку при !Valid
	parsedToken, err := tid.parseToken(token)
	if err != nil {
		return co.ParseError(co.ErrInvalidToken)
	}
	// Получение User Id
	ret, err := json.Marshal(UserId{ID: parsedToken.UserId.ToString()})
	if err != nil {
		return co.ParseError(co.ErrInternalServerError)
	}
	return co.Response{
		StatusCode: http.StatusOK,
		Message:    ret,
	}
}

// GetHandler - возвращает обработчик для получения user Id
//
// HTTP-метод - POST
//
// Возвращает ошибку, если токен невалиден или возникла ошибка при обработке тела или токена
func (tid TokensIdGetter) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := tid.GetTokenId(r)
		w.WriteHeader(response.StatusCode)
		w.Write(response.Message)
	})
}

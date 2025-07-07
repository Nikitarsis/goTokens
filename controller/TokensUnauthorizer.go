package controller

import (
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

type TokensUnauthorizer struct {
	parseToken func(co.Token) (co.TokenData, error)
	dropKey   func(co.UUID) bool
}

func (tu TokensUnauthorizer) UnauthorizeTokens(request *http.Request) co.Response {
	// Парсинг тела запроса
	token, err := parseBody(request)
	if err != nil {
		return co.ParseError(err)
	}
	// Парсинг токена
	parsedToken, err := tu.parseToken(token)
	if err != nil {
		return co.ParseError(co.ErrInvalidToken)
	}
	//Токен должен быть Access
	if parsedToken.Type != co.AccessToken {
		return co.ParseError(co.ErrWrongToken)
	}
	// Удаление ключа
	go tu.dropKey(parsedToken.KeyId)
	return co.Response{
		StatusCode: http.StatusOK,
		Message:    []byte("Token unauthorized"),
	}
}

func (tu TokensUnauthorizer) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := tu.UnauthorizeTokens(r)
		w.WriteHeader(response.StatusCode)
		w.Write(response.Message)
	})
}
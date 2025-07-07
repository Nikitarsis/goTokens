package controller

import (
	"encoding/json"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

type TokensUnauthorizer struct {
	parseToken func(co.Token) (co.TokenData, error)
	dropKey   func(co.UUID) bool
}

func (tu TokensUnauthorizer) UnauthorizeTokens(request *http.Request) co.Response {
	var rawToken UserToken
	var body []byte
	//Проверка метода, должен быть POST
	if request.Method != http.MethodPost {
		return co.ParseError(co.ErrInvalidMethod)
	}
	// Чтение тела запроса
	_, err := request.Body.Read(body)
	if err != nil {
		return co.ParseError(co.ErrInternalServerError)
	}
	// Парсинг тела запроса
	err = json.Unmarshal(body, &rawToken)
	if err != nil {
		return co.ParseError(co.ErrJsonParsingError)
	}
	// Создание токена
	token := co.Token{
		Value: rawToken.Token,
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
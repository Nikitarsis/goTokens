package controller

import (
	"encoding/json"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

// TokenRefresher - это класс для обновления токенов
type TokensRefresher struct {
	getPairTokens func(co.UUID) (TokensPair, error)
	parseToken    func(co.Token) (co.TokenData, error)
	dropKey       func(co.UUID) bool
}

// NewTokensRefresher - это конструктор для TokensRefresher
func NewTokensRefresher(
	getPairTokens func(co.UUID) (TokensPair, error),
	parseToken func(co.Token) (co.TokenData, error),
	dropKey func(co.UUID) bool,
	) *TokensRefresher {
	return &TokensRefresher{
		getPairTokens: getPairTokens,
		parseToken:    parseToken,
		dropKey:       dropKey,
	}
}

// RefreshTokens - это метод для обновления токенов
func (tr TokensRefresher) RefreshTokens(request *http.Request) (co.Response) {
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
	// Получение UID
	uid, err := co.GetUUIDFromString(rawToken.UID)
	if err != nil {
		return co.ParseError(co.ErrInvalidUserId)
	}
	// Парсинг токена
	parsedToken, err := tr.parseToken(token)
	if err != nil {
		return co.ParseError(co.ErrInvalidToken)
	}
	// Проверка соответствия UID
	if (parsedToken.UserId != uid) {
		// Удаление ключа как скомпроментированного
		go tr.dropKey(parsedToken.KeyId)
		return co.ParseError(co.ErrStealedToken)
	}
	// Получение пары токенов
	pair, err := tr.getPairTokens(uid)
	if err != nil {
		return co.ParseError(co.ErrInternalServerError)
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
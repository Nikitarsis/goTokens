package controller

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strconv"

	co "github.com/Nikitarsis/goTokens/common"
)

// traceIp - это функция для запуска трассировки ip адреса
//
// возвращает ошибку, если remoteAddr некорректен и если не удаётся разобрать port
func traceIp(token co.TokenData, remoteAddr string, ipTracer co.IIpTracer) error {
	ip, portRaw, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		if remoteAddr != "" {
			return co.ErrInternalServerError
		}
	}
	port, err := strconv.ParseUint(portRaw, 10, 16)
	if err != nil {
		if portRaw != "" {
			return co.ErrInternalServerError
		}
		port = 0
	}
	go func() {
		ipTracer.TraceIp(co.DataIP{
			IP:     net.ParseIP(ip),
			Port:   uint16(port),
			UserId: token.UserId,
			KeyId:  token.KeyId,
		})
	}()
	return nil
}

// parseBodyWithId - парсит тело запроса и извлекает токен и UID
//
// Возвращает ошибку, если не удаётся пропарсить токен и тело
func parseBodyWithId(request *http.Request) (co.Token, co.UUID, error) {
	var rawToken UserToken
	//Чтение тела запроса
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return co.Token{}, co.UUID{}, co.ErrInternalServerError
	}
	// Парсинг тела запроса
	err = json.Unmarshal(body, &rawToken)
	if err != nil {
		return co.Token{}, co.UUID{}, co.ErrJsonParsingError
	}
	// Создание токена
	token := co.Token{
		Value: rawToken.Token,
	}
	// Получение UID
	uid, err := co.GetUUIDFromString(rawToken.UID)
	if err != nil {
		return co.Token{}, co.UUID{}, co.ErrInvalidUserId
	}
	return token, uid, nil
}

// parseBody - парсит тело запроса и получает токен
func parseBody(request *http.Request) (co.Token, error) {
	var rawToken UserToken
	//Чтение тела запроса
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return co.Token{}, co.ErrInternalServerError
	}
	// Парсинг тела запроса
	err = json.Unmarshal(body, &rawToken)
	if err != nil {
		return co.Token{}, co.ErrJsonParsingError
	}
	// Создание токена
	token := co.Token{
		Value: rawToken.Token,
	}
	return token, nil
}

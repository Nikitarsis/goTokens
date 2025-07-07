package controller

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strconv"

	co "github.com/Nikitarsis/goTokens/common"
)

func traceIp(token co.TokenData, remoteAddr string, ipTracer co.IIpRepository) error {
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

func parseBodyWithId(request *http.Request) (co.Token, co.UUID, error) {
	var rawToken UserToken
	//Чтение тела запроса
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	//Проверка метода, должен быть POST
	if request.Method != http.MethodPost {
		return co.Token{}, co.UUID{}, co.ErrInvalidMethod
	}
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

func parseBody(request *http.Request) (co.Token, error) {
	var rawToken UserToken
	//Чтение тела запроса
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	//Проверка метода, должен быть POST
	if request.Method != http.MethodPost {
		return co.Token{}, co.ErrInvalidMethod
	}
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

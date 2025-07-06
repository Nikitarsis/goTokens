package controller

import (
	"encoding/json"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

type TokensPairGetter struct {
	callback func(co.UUID) (co.TokensPair, error)
}

func NewTokensPairGetter(callback func(co.UUID) (co.TokensPair, error)) *TokensPairGetter {
	return &TokensPairGetter{
		callback: callback,
	}
}

func (tg TokensPairGetter) parseRequestGet(request *http.Request) (co.UUID, error) {
	userIdStr := request.URL.Query().Get("uid")
	if userIdStr == "" {
		return co.UUID{}, ErrNoUserId
	}
	userId, err := co.GetUUIDFromString(userIdStr)
	if err != nil {
		return co.UUID{}, ErrInvalidUserId
	}
	return userId, nil
}

func (tg TokensPairGetter) parseRequestPost(request *http.Request) (co.UUID, error) {
	var userId UserId
	var body []byte
	_, err := request.Body.Read(body)
	if err != nil {
		return co.UUID{}, ErrJsonParsingError
	}
	err = json.Unmarshal(body, &userId)
	if err != nil {
		return co.UUID{}, ErrInvalidUserId
	}
	ret, err := co.GetUUIDFromString(userId.ID)
	if err != nil {
		return co.UUID{}, ErrInvalidUserId
	}
	return ret, nil
}

func (tg TokensPairGetter) GetTokensPair(request *http.Request) Response {
	method := request.Method
	var userId co.UUID
	var err error
	switch method {
	case http.MethodGet:
		userId, err = tg.parseRequestGet(request)
		if err != nil {
			return ParseError(err)
		}
	case http.MethodPost:
		userId, err = tg.parseRequestPost(request)
		if err != nil {
			return ParseError(err)
		}
	default:
		return Response{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    []byte("Method not allowed"),
		}
	}
	result, err:= tg.callback(userId)
	if err != nil {
		return ParseError(ErrInternalServerError)
	}
	ret, err := json.Marshal(result)
	if err != nil {
		return ParseError(ErrInternalServerError)
	}
	return Response{
		StatusCode: http.StatusOK,
		Message:    ret,
	}
}

func (tg TokensPairGetter) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := tg.GetTokensPair(r)
		w.WriteHeader(response.StatusCode)
		w.Write(response.Message)
	})
}
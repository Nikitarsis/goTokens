package controller

import (
	"encoding/json"
	"io"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
)

type TokensPairGetter struct {
	getPairTokens       func(co.UUID) (map[string]co.TokenData, error)
	userAgentRepository co.IUserAgentRepository
	ipRepository        co.IIpRepository
}

func NewTokensPairGetter(
	getPairTokens func(co.UUID) (map[string]co.TokenData, error),
	userAgentRepository co.IUserAgentRepository,
	ipRepository co.IIpRepository,
) *TokensPairGetter {
	return &TokensPairGetter{
		getPairTokens:       getPairTokens,
		userAgentRepository: userAgentRepository,
		ipRepository:        ipRepository,
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
	body, err := io.ReadAll(request.Body)
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
		return ParseError(ErrInvalidMethod)
	}
	result, err := tg.getPairTokens(userId)
	if err != nil {
		return ParseError(ErrInternalServerError)
	}
	pair := TokensPair{
		Access:  result["access"].Token.ToString(),
		Refresh: result["refresh"].Token.ToString(),
	}
	ret, err := json.Marshal(pair)
	if err != nil {
		return ParseError(ErrInternalServerError)
	}
	go tg.userAgentRepository.SaveUserAgent(
		result["refresh"].KeyId,
		request.UserAgent(),
	)
	go tg.ipRepository.TraceIp(
		result["refresh"].KeyId,
		request.RemoteAddr,
	)
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

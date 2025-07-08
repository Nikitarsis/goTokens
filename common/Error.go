package common

import (
	"errors"
	"net/http"
)

var (
	// ErrInvalidToken - Ошибка, возникающая при недействительном токене (!Valid)
	ErrInvalidToken = errors.New("invalid token")
	// ErrNoFindKey - Ошибка, возникающая при отсутствии ключа
	ErrNoFindKey = errors.New("key not found")
	// ErrNoUserId - Ошибка, возникающая при отсутствии userId
	ErrNoUserId            = errors.New("no userId in request")
	// ErrInvalidUserId - Ошибка, возникающая при недействительном userId
	ErrInvalidUserId       = errors.New("invalid userId")
	// ErrJsonParsingError - Ошибка, возникающая при ошибке парсинга JSON
	ErrJsonParsingError    = errors.New("error parsing JSON")
	// ErrCannotParseUserId - Ошибка, возникающая при невозможности парсинга userId
	ErrCannotParseUserId   = errors.New("cannot parse userId")
	// ErrInternalServerError - Ошибка, возникающая при внутренней ошибке сервера
	ErrInternalServerError = errors.New("internal server error")
	// ErrStealedToken - Ошибка, возникающая при использовании украденного токена
	ErrStealedToken        = errors.New("stealed token")
	// ErrInvalidMethod - Ошибка, возникающая при недопустимом методе
	ErrInvalidMethod       = errors.New("invalid method")
	// ErrWrongToken - Ошибка, возникающая при неверном токене
	ErrWrongToken          = errors.New("wrong token")
	// ErrInvalidUserAgent - Ошибка, возникающая при недопустимом User-Agent
	ErrInvalidUserAgent    = errors.New("invalid user agent")
)

// ParseError - функция для обработки ошибок
func ParseError(err error) Response {
	switch err {
	case ErrNoUserId:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("User ID is wrong or missing"),
		}
	case ErrInvalidUserId:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("Invalid User ID"),
		}
	case ErrJsonParsingError:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("Error parsing JSON"),
		}
	case ErrCannotParseUserId:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("Cannot parse User ID"),
		}
	case ErrInternalServerError:
		return Response{
			StatusCode: http.StatusInternalServerError,
			Message:    []byte("Internal Server Error"),
		}
	case ErrInvalidToken:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("Invalid Token"),
		}
	case ErrStealedToken:
		return Response{
			StatusCode: http.StatusForbidden,
			Message:    []byte("Token is outdated"),
		}
	case ErrInvalidMethod:
		return Response{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    []byte("Invalid Method"),
		}
	case ErrWrongToken:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("Wrong Token"),
		}
	case ErrInvalidUserAgent:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("Invalid User Agent"),
		}
	default:
		return Response{
			StatusCode: http.StatusInternalServerError,
			Message:    []byte("Internal Server Error"),
		}
	}
}

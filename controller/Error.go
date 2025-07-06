package controller

import (
	"errors"
	"net/http"
)

var (
	// ErrNoUserId - Ошибка, возникающая при отсутствии userId
	ErrNoUserId = errors.New("no userId in request")
	ErrInvalidUserId = errors.New("invalid userId")
	ErrJsonParsingError = errors.New("error parsing JSON")
	ErrCannotParseUserId = errors.New("cannot parse userId")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidToken = errors.New("invalid token")
	ErrStealedToken = errors.New("stealed token")
	ErrInvalidMethod = errors.New("invalid method")
)

func ParseError(err error) Response {
	switch err {
	case ErrNoUserId:
		return Response{
			StatusCode: http.StatusBadRequest,
			Message:    []byte("User ID is required"),
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
	default:
		return Response{
			StatusCode: http.StatusInternalServerError,
			Message:    []byte("Internal Server Error"),
		}
	}
}
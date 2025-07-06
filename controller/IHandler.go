package controller

import "net/http"

type IHandler interface {
	GetHandler() http.Handler
}

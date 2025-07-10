package controller

import "net/http"

// IHandler - интерфейс для обработчиков HTTP-запросов
type IHandler interface {
	GetHandler() http.Handler
}

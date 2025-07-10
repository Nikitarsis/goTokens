package controller

import "net/http"

// HttpServerBuilder - структура для построения HTTP-сервера
type HttpServerBuilder struct {
	handlerMap  map[string]IHandler
	serverAddr  string
}

// InitHttpServerBuilder - инициализирует HttpServerBuilder
func InitHttpServerBuilder() *HttpServerBuilder {
	return &HttpServerBuilder{
		handlerMap: make(map[string]IHandler),
		serverAddr: ":8080",
	}
}

// SetServerAddr - устанавливает адрес сервера
func (hs *HttpServerBuilder) SetServerAddr(path string) *HttpServerBuilder {
	hs.serverAddr = path
	return hs
}

// AddHandler - добавляет обработчик для указанного пути
func (hs *HttpServerBuilder) AddHandler(path string, handler IHandler) *HttpServerBuilder {
	hs.handlerMap[path] = handler
	return hs
}

// Build - строит HTTP-сервер
func (hs *HttpServerBuilder) Build() http.Server {
	mux := http.NewServeMux()
	for path, handler := range hs.handlerMap {
		mux.Handle(path, handler.GetHandler())
	}
	return http.Server{
		Addr:    hs.serverAddr,
		Handler: mux,
	}
}

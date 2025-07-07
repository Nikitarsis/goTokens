package controller

import "net/http"

type HttpServerBuilder struct {
	handlerMap  map[string]IHandler
	serverPort  string
}

func InitHttpServerBuilder() *HttpServerBuilder {
	return &HttpServerBuilder{
		handlerMap: make(map[string]IHandler),
		serverPort: ":8080",
	}
}

func (hs *HttpServerBuilder) GetHandler(path string) *HttpServerBuilder {
	hs.serverPort = path
	return hs
}

func (hs *HttpServerBuilder) AddHandler(path string, handler IHandler) *HttpServerBuilder {
	hs.handlerMap[path] = handler
	return hs
}

func (hs *HttpServerBuilder) Build() http.Server {
	mux := http.NewServeMux()
	for path, handler := range hs.handlerMap {
		mux.Handle(path, handler.GetHandler())
	}
	return http.Server{
		Addr:    hs.serverPort,
		Handler: mux,
	}
}

package main

import (
	"fmt"
	"net/http"

	co "github.com/Nikitarsis/goTokens/common"
	con "github.com/Nikitarsis/goTokens/controller"
)

func GetTestHandler(isDebug bool) con.IHandler {
	if isDebug {
		return &DebugHandler{}
	}
	return &StubDebugHandler{}
}

type DebugHandler struct{}

func (dh *DebugHandler) greeting(r *http.Request) string {
	ret := fmt.Sprintf("Hiiii, %s from %s (*ﾟｰﾟ)ゞ\n", r.UserAgent(), r.RemoteAddr)
	ret += "Debug mode is active\n"
	ret += "U can fear nothin ( ◡‿◡ *)\n"
	return ret
}

func (dh *DebugHandler) getInfo(r *http.Request, isWrongPath bool) string {
	ret := dh.greeting(r)
	if isWrongPath {
		ret += "But you might want to check your request... (￣▽￣)\n"
		ret += fmt.Sprintf("%s is not valid\n", r.URL.Path)
		ret += "And I DO HATE when I'm being worried unnecessarily. (◉Θ◉)\n"
	}
	ret += "Available methods:\n"
	ret += " GET /test/check - checking debug\n"
	ret += " GET /test/id - getting random id\n"
	ret += " GET /test/key - getting random user key\n"
	return ret
}

func (dh *DebugHandler) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/test/info":
			w.Write([]byte(dh.getInfo(r, false)))
		case "/test/check":
			w.Write([]byte(dh.greeting(r)))
		case "/test/id":
			w.Write([]byte(co.GetTestUUID().ToString()))
		case "/test/key":
			w.Write([]byte(co.CreateTestKey().ToString()))
		default:
			w.Write([]byte(dh.getInfo(r, true)))
		}
	})
}

type StubDebugHandler struct{}

func (h *StubDebugHandler) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ret := fmt.Sprintf("Hiiii, %s from %s (*ﾟｰﾟ)ゞ\n", r.UserAgent(), r.RemoteAddr)
		ret += "Lets give u all sentitive data! (ノ°ο°)ノ\n"
		ret += "... (´･_･｀)\n"
		ret += "Waiiiiit! (」°ロ°)」\n"
		ret += "Debug mode is turned Off. (*μ_μ)\n"
		ret += "What are u lookin for here, buddy? <(￣ ﹌ ￣)>\n"
		ret += "Probably you want to know more about me... (o^ ^o)\n"
		ret += "I know where u live... (◉Θ◉)\n"
		ret += "Just kidding! (￣▽￣)\n"
		ret += "Bye! (◉Θ◉)"
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprint(w, ret)
	})
}

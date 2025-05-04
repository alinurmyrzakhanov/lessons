package router

import (
	"net/http"

	pingHandler "github.com/alinurmyrzakhanov/lessons/lesson4/example1/internal/api/handler/ping"
)

func ping(mux *http.ServeMux) {
	h := pingHandler.New()

	mux.Handle("GET /ping", http.HandlerFunc(h.Ping))
}

package router

import (
	"net/http"

	aiHandler "github.com/alinurmyrzakhanov/lessons/lesson4/example1/internal/api/handler/ai"
)

func ai(mux *http.ServeMux, apiKey string) {
	h := aiHandler.New(apiKey)

	mux.Handle("POST /prompt", http.HandlerFunc(h.Prompt))
}

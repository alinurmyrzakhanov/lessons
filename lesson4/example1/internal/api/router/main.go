package router

import (
	"net/http"

	apiT "github.com/alinurmyrzakhanov/lessons/lesson4/example1/internal/api/types"
)

// SetupRoutes setup router api
func SetupRoutes(mux *http.ServeMux, api apiT.Api) {
	ping(mux)
	ai(mux, api.Config().AIApiKey)
}

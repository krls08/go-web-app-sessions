package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/krls08/go-web-app-sessions/internal/config"
	"github.com/krls08/go-web-app-sessions/internal/handlers"
)

func NewRouter(app *config.AppConfig, hr *handlers.HandlerRepo) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(hr.Home))
	mux.Get("/about", http.HandlerFunc(hr.About))

	return mux
}

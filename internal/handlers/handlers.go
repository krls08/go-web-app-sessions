package handlers

import (
	"net/http"

	"github.com/krls08/go-web-app-sessions/internal/config"
	"github.com/krls08/go-web-app-sessions/internal/render/domain"
	render_service "github.com/krls08/go-web-app-sessions/internal/render/service"
)

type HandlerRepo struct {
	App *config.AppConfig
	rs  render_service.RenderService
}

func NewHanldersRepo(a *config.AppConfig, render render_service.RenderService) *HandlerRepo {
	return &HandlerRepo{
		App: a,
		rs:  render,
	}
}

// Home is the home page handler
func (h *HandlerRepo) Home(w http.ResponseWriter, r *http.Request) {
	h.rs.RenderTemplate(w, "home.page.tmpl", &domain.TemplateData{})
}

// About is the about page handler
func (h *HandlerRepo) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello again..."
	h.rs.RenderTemplate(w, "about.page.tmpl", &domain.TemplateData{
		StringMap: stringMap,
	})
}

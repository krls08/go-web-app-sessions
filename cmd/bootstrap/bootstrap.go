package bootstrap

import (
	"context"
	"log"

	"github.com/krls08/go-web-app-sessions/cmd/routes"
	"github.com/krls08/go-web-app-sessions/internal/config"
	"github.com/krls08/go-web-app-sessions/internal/handlers"
	render_service "github.com/krls08/go-web-app-sessions/internal/render/service"
	"github.com/krls08/go-web-app-sessions/internal/server"
)

func Run() error {
	var app config.AppConfig
	app.UseCache = false

	rs := render_service.NewTemplates(&app)

	tc, err := rs.CreateTemplateCache()
	if err != nil {
		log.Fatal("[ERROR] Cannot create template cache:", err)
	}

	app.TemplateCache = tc

	hr := handlers.NewHanldersRepo(&app, rs)
	mux := routes.NewRouter(&app, hr)

	ctx := context.TODO()
	s := server.New(ctx, "localhost", 60002, mux)

	return s.Run(ctx)
}

package bootstrap

import (
	"context"
	"log"

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

	ctx := context.TODO()
	s := server.NewServer(ctx, "localhost", 60002, hr)

	return s.Run(ctx)
}

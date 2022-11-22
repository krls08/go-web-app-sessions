package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/krls08/go-web-app-sessions/internal/server/handlers"
)

type Server struct {
	httpAddr string
	mux      *http.ServeMux
	// h handlers.HomeHandlers
}

func New(ctx context.Context, host string, port uint) Server { //(context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		mux:      http.NewServeMux(),

		// handlers
		//hh: homeHandlers,
	}

	srv.registerRoutes()
	//return serverContext(ctx), srv
	return srv

}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Listening on %s\n", s.httpAddr)

	return http.ListenAndServe(s.httpAddr, s.mux)
}

func (s *Server) registerRoutes() {
	fmt.Println("registring routes...")
	s.mux.HandleFunc("/", handlers.Home)
	s.mux.HandleFunc("/about", handlers.About)
}

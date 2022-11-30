package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/krls08/go-web-app-sessions/internal/handlers"
)

type Server struct {
	httpAddr string
	mux      *pat.PatternServeMux
	h        *handlers.HandlerRepo

	//mux      http.Handler
	//mux      *http.ServeMux
}

func NewServer(ctx context.Context, host string, port uint, hr *handlers.HandlerRepo) Server { //(context.Context, Server) {
	//func New(ctx context.Context, host string, port uint, router http.Handler, hr *handlers.HandlerRepo) Server { //(context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		//mux:      http.NewServeMux(),
		//mux: pat.New(),

		// Handlers
		h: hr,
	}

	//return serverContext(ctx), srv
	return srv

}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Listening on %s\n", s.httpAddr)

	s.mux = s.registerRoutes()

	return http.ListenAndServe(s.httpAddr, s.mux)
}

func (s *Server) registerRoutes() *pat.PatternServeMux {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(s.h.Home))
	mux.Get("/about", http.HandlerFunc(s.h.About))
	return mux
}

//func (s *Server) registerRoutes() {
//	fmt.Println("registring routes...")
//	s.mux.HandleFunc("/", s.h.Home)
//	s.mux.HandleFunc("/about", s.h.About)
//}

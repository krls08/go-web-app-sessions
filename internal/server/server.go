package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

type Server struct {
	httpAddr string
	//mux      *http.ServeMux
	mux http.Handler
	//h   *handlers.HandlerRepo
}

func New(ctx context.Context, host string, port uint, router http.Handler) Server { //(context.Context, Server) {
	//func New(ctx context.Context, host string, port uint, router http.Handler, hr *handlers.HandlerRepo) Server { //(context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		//mux:      http.NewServeMux(),
		mux: router,

		// Handlers
		//h: hr,
	}

	//return serverContext(ctx), srv
	return srv

}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Listening on %s\n", s.httpAddr)

	return http.ListenAndServe(s.httpAddr, s.mux)
}

func (s *Server) registerRoutes(mux *pat.PatternServeMux) {
	s.mux = mux
}

//func (s *Server) registerRoutes() {
//	fmt.Println("registring routes...")
//	s.mux.HandleFunc("/", s.h.Home)
//	s.mux.HandleFunc("/about", s.h.About)
//}

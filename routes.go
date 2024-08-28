package imgdemo

import (
	"net/http"
	// _ "net/http/pprof"
)

func (s *Server) routes() {

	r := s.router

	//withLogging := s.WithLogging
	//withMetrics := s.WithMetrics(withLogging)

	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	r.HandleFunc("/resize", s.handleImage())

}

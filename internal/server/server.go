package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.temporal.io/sdk/client"

	"github.com/jacobjlee/temporal-detection-workflow/internal/detection"
)

// healthCheckHandler is a simple handler to check the health of the server
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "OK")
}

// Server struct
type Server struct {
	Router           *chi.Mux
	temporalClient   client.Client
	detectionService detection.Service
}

// NewServer creates a new Server struct
func NewServer(temporalClient client.Client, detectionService detection.Service) *Server {
	return &Server{
		Router:           chi.NewRouter(),
		temporalClient:   temporalClient,
		detectionService: detectionService,
	}
}

// MountHandlers mounts the handlers to the router
func (s *Server) MountHandlers() {
	// Mount health check handler
	s.Router.Get("/health", healthCheckHandler)

	// Mount v1 api router
	s.Router.Mount("/v1", apiRouter(s.temporalClient, s.detectionService))
}

// Run starts the server
func (s *Server) Run(host string, port int) error {
	addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	return http.ListenAndServe(addr, s.Router)
}

package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.temporal.io/sdk/client"

	"github.com/jacobjlee/temporal-detection-workflow/internal/detection"
)

// apiRouter mounts the api handlers to the router
func apiRouter(temporalClient client.Client, detectionService detection.Service) chi.Router {
	r := chi.NewRouter()

	r.Mount("/detection", detectionRouter(temporalClient, detectionService))

	return r
}

// detectionRouter creates a router for the detection api
func detectionRouter(temporalClient client.Client, detectionService detection.Service) chi.Router {
	r := chi.NewRouter()

	r.Post("/start", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := detectionService.Start(ctx, temporalClient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r
}

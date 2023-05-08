package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.temporal.io/sdk/client"

	"github.com/jacobjlee/temporal-detection-workflow/internal/detection"
)

// DetectionEndRequest is the request body for the detection end endpoint
type DetectionEndRequest struct {
	WorkflowID string `json:"workflow_id"`
}

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

	r.Post("/end", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req DetectionEndRequest
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := detectionService.End(ctx, temporalClient, req.WorkflowID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r
}

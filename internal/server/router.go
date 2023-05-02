package server

import (
	"github.com/go-chi/chi/v5"
	"go.temporal.io/sdk/client"
)

// apiRouter mounts the api handlers to the router
func apiRouter(temporalClient client.Client) chi.Router {
	r := chi.NewRouter()

	return r
}

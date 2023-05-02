package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func TestHealthRoute(t *testing.T) {

	// Create a New Server Struct
	s := NewServer(nil)
	// Mount Handlers
	s.MountHandlers()

	// Create a New Request
	req, _ := http.NewRequest("GET", "/health", nil)

	// Execute Request
	response := executeRequest(req, s)

	assert.Equal(t, http.StatusOK, response.Code)  // Check the response code
	require.Equal(t, "OK", response.Body.String()) // Check the response body
}

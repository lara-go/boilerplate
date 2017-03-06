package middleware

import (
	"github.com/lara-go/larago/http"
	"github.com/lara-go/larago/http/responses"
)

// Simple test.
type Simple struct{}

// Handle request.
func (m *Simple) Handle(request *http.Request, next http.Handler) responses.Response {
	return next(request)
}

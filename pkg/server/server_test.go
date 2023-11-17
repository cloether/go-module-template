package server

import (
	"context"
	"net/http/httptest"
	"testing"
)

func TestHandleAbout(t *testing.T) {
	srv := New(context.Background())
	srv.routes()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	srv.router.ServeHTTP(w, req)
}

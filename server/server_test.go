package main

import (
	"context"
	"net/http/httptest"
	"testing"
)

func TestHandleAbout(t *testing.T) {
	srv := NewServer(context.Background())
	srv.routes()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	srv.router.ServeHTTP(w, req)
}
